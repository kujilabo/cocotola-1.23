//go:build small

package controller_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	controller "github.com/kujilabo/cocotola-1.23/cocotola-auth/controller/gin"
	controllermock "github.com/kujilabo/cocotola-1.23/cocotola-auth/controller/gin/mocks"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/domain"
	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
)

func initGoogleRouter(t *testing.T, ctx context.Context, googleUser controller.GoogleUserUsecase) *gin.Engine {
	t.Helper()
	fn := controller.NewInitGoogleRouterFunc(googleUser)

	initPublicRouterFuncs := []libcontroller.InitRouterGroupFunc{fn}
	// initPrivateRouterFuncs := []libcontroller.InitRouterGroupFunc{}

	router := libcontroller.InitRootRouterGroup(ctx, corsConfig, debugConfig)
	api := router.Group("api")
	v1 := api.Group("v1")

	libcontroller.InitPublicAPIRouterGroup(ctx, v1, initPublicRouterFuncs)

	return router
}

func TestGoogleAuthHandler_Authorize_shouldReturn400_whenRequestBodyIsEmpty(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	googleUserUsecase := new(controllermock.GoogleUserUsecase)
	r := initGoogleRouter(t, ctx, googleUserUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/google/authorize", bytes.NewBuffer([]byte("")))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusBadRequest, w.Code, "status code should be 400")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Equal(t, "Bad Request", message[0], "message should be 'Bad Request'")
}

func TestGoogleAuthHandler_Authorize_shouldReturn400_whenRequestBodyIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	googleUserUsecase := new(controllermock.GoogleUserUsecase)
	r := initGoogleRouter(t, ctx, googleUserUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/google/authorize", bytes.NewBuffer([]byte("[]")))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusBadRequest, w.Code, "status code should be 400")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Equal(t, "Bad Request", message[0], "message should be 'Bad Request'")
}

func TestGoogleAuthHandler_Authorize_shouldReturn500_whenErrorOccur(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	googleUserUsecase := new(controllermock.GoogleUserUsecase)
	googleUserUsecase.On("Authorize", anyOfCtx, "VALID_STATE", "ERROR_CODE", "ORG_NAME").Return(nil, errors.New("ERROR"))
	r := initGoogleRouter(t, ctx, googleUserUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/google/authorize", bytes.NewBuffer([]byte(`{"organizationName": "ORG_NAME", "sessionState": "VALID_STATE", "paramState": "VALID_STATE", "code": "ERROR_CODE"}`)))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Equal(t, "Internal Server Error", message[0], "message should be 'Internal Server Error'")
}

func TestGoogleAuthHandler_Authorize_shouldReturn401_whenCodeIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	googleUserUsecase := new(controllermock.GoogleUserUsecase)
	googleUserUsecase.On("Authorize", anyOfCtx, "VALID_STATE", "INVALID_CODE", "ORG_NAME").Return(nil, domain.ErrUnauthenticated)
	r := initGoogleRouter(t, ctx, googleUserUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/google/authorize", bytes.NewBuffer([]byte(`{"organizationName": "ORG_NAME", "sessionState": "VALID_STATE", "paramState": "VALID_STATE", "code": "INVALID_CODE"}`)))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code, "status code should be 500")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Equal(t, "Unauthorized", message[0], "message should be 'Unauthorized'")
}

func TestGoogleAuthHandler_Authorize_shouldReturn401_whenCodeIsValid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authToksenSet := &domain.AuthTokenSet{
		AccessToken:  "ACCESS_TOKEN",
		RefreshToken: "REFRESH_TOKEN",
	}
	googleUserUsecase := new(controllermock.GoogleUserUsecase)
	googleUserUsecase.On("Authorize", anyOfCtx, "VALID_STATE", "VALID_CODE", "ORG_NAME").Return(authToksenSet, nil)
	r := initGoogleRouter(t, ctx, googleUserUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/google/authorize", bytes.NewBuffer([]byte(`{"organizationName": "ORG_NAME", "sessionState": "VALID_STATE", "paramState": "VALID_STATE", "code": "VALID_CODE"}`)))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Len(t, message, 0, "message should be null")

	accessTokenExpr := parseExpr(t, "$.accessToken")
	accessToken := accessTokenExpr.Get(jsonObj)
	assert.Equal(t, "ACCESS_TOKEN", accessToken[0], "accessToken should be 'ACCESS_TOKEN'")

	refreshTokenExpr := parseExpr(t, "$.refreshToken")
	refreshToken := refreshTokenExpr.Get(jsonObj)
	assert.Equal(t, "REFRESH_TOKEN", refreshToken[0], "refreshToken should be 'REFRESH_TOKEN'")
}
