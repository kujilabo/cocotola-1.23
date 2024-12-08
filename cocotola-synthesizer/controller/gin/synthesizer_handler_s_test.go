//go:build small

package handler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	handler "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/controller/gin"
	handlermock "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/controller/gin/mocks"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"
)

var (
	anyOfCtx           = mock.MatchedBy(func(_ context.Context) bool { return true })
	corsConfig         cors.Config
	appConfig          *config.AppConfig
	debugConfig        *libconfig.DebugConfig
	internalAuthConfig config.InternalAuthConfig
	// authTokenManager  auth.AuthTokenManager
)

func init() {
	corsConfig = cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
	}
	appConfig = &config.AppConfig{
		Name:        "test",
		HTTPPort:    8080,
		MetricsPort: 8081,
	}
	debugConfig = &libconfig.DebugConfig{
		Gin:  false,
		Wait: false,
	}
	internalAuthConfig = config.InternalAuthConfig{
		Username: "username",
		Password: "password",
	}
}

func initSynthesizerRouter(t *testing.T, ctx context.Context, workbokQueryUsecase handler.SynthesizerUsecase) *gin.Engine {
	t.Helper()
	fn := handler.NewInitSynthesizerRouterFunc(workbokQueryUsecase)

	authMiddleware := gin.BasicAuth(gin.Accounts{
		internalAuthConfig.Username: internalAuthConfig.Password,
	})
	initPublicRouterFunc := []handler.InitRouterGroupFunc{}
	initPrivateRouterFunc := []handler.InitRouterGroupFunc{fn}

	router := gin.New()
	err := handler.InitRouter(ctx, router, authMiddleware, initPublicRouterFunc, initPrivateRouterFunc, corsConfig, debugConfig, appConfig.Name)
	require.NoError(t, err)

	return router
}

func TestSynthesizerHandler_Synthesize_shouldReturn200(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	synthesizerUsecase := new(handlermock.SynthesizerUsecase)
	synthesizerUsecase.On("Synthesize", anyOfCtx, mock.Anything, mock.Anything, mock.Anything).Return(&domain.AudioModel{
		AudioID: &domain.AudioID{Value: 1},
		Lang5:   libdomain.Lang5JAJP,
		Text:    "こんにちは",
		Content: "CONTENT",
		Length:  time.Duration(1234) * time.Millisecond,
	}, nil)

	// given
	r := initSynthesizerRouter(t, ctx, synthesizerUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/synthesize/synthesize", bytes.NewReader([]byte(`{"lang5":"ja-JP","voice":"ja-JP-Wavenet-A","text":"こんにちは"}`)))
	require.NoError(t, err)
	req.SetBasicAuth("username", "password")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

	jsonObj := parseJSON(t, respBytes)

	audioContentExpr := parseExpr(t, "$.audioContent")
	audioContent := audioContentExpr.Get(jsonObj)
	assert.Len(t, audioContent, 1, "response should has one audioContent")
	assert.Equal(t, "CONTENT", audioContent[0], "audioContent should be 'CONTENT'")

	audioLengthMillisecondExpr := parseExpr(t, "$.audioLengthMillisecond")
	audioLengthMillisecond := audioLengthMillisecondExpr.Get(jsonObj)
	assert.Equal(t, int64(1234), audioLengthMillisecond[0], "audioLengthMillisecond[0] should be 1234")
}

func TestSynthesizerHandler_Synthesize_shouldReturn401_whenAuthorizationHeaderIsEmpty(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	synthesizerUsecase := new(handlermock.SynthesizerUsecase)

	// given
	r := initSynthesizerRouter(t, ctx, synthesizerUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/synthesize/synthesize", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code, "status code should be 401")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Len(t, message, 0, "message should be empty")
}

func TestSynthesizerHandler_Synthesize_shouldReturn401_whenAuthorizationHeaderIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	synthesizerUsecase := new(handlermock.SynthesizerUsecase)

	// given
	r := initSynthesizerRouter(t, ctx, synthesizerUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/synthesize/synthesize", nil)
	require.NoError(t, err)
	req.SetBasicAuth("username", "invalid_password")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code, "status code should be 401")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Len(t, message, 0, "message should be empty")
}
