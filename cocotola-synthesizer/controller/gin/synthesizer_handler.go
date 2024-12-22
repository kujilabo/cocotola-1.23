package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"
	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"
)

type AudioResponse struct {
	ID      int    `json:"id"`
	Lang5   string `json:"lang5"`
	Text    string `json:"text"`
	Content string `json:"content"`
}

type SynthesizerUsecase interface {
	Synthesize(ctx context.Context, lang5 *libdomain.Lang5, void, text string) (*domain.AudioModel, error)
}

type SynthesizerHandler struct {
	synthesizerUsecase SynthesizerUsecase
	logger             *slog.Logger
}

func NewSynthesizerHandler(synthesizerUsecase SynthesizerUsecase) *SynthesizerHandler {
	return &SynthesizerHandler{
		synthesizerUsecase: synthesizerUsecase,
		logger:             slog.Default().With(slog.String(rsliblog.LoggerNameKey, "SynthesizerHandler")),
	}
}

func (h *SynthesizerHandler) Synthesize(c *gin.Context) {
	ctx := c.Request.Context()

	synthesizeParameter := libapi.SynthesizeParameter{}
	if err := c.ShouldBindJSON(&synthesizeParameter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	lang5, err := libdomain.NewLang5(synthesizeParameter.Lang5)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	audioModel, err := h.synthesizerUsecase.Synthesize(ctx, lang5, synthesizeParameter.Voice, synthesizeParameter.Text)
	if err != nil {
		h.logger.ErrorContext(ctx, "synthesizerUsecase.Synthesize", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, &libapi.SynthesizeResponse{
		AudioContent:           audioModel.Content,
		AudioLengthMillisecond: int(audioModel.Length.Milliseconds()),
	})
}

func (h *SynthesizerHandler) FindAudioByID(c *gin.Context) {

}

// func (h *SynthesizerHandler) errorHandle(ctx context.Context, logger *slog.Logger, c *gin.Context, err error) bool {
// 	// if errors.Is(err, service.ErrAudioNotFound) {
// 	// 	logger.Warnf("PrivateSynthesizerHandler err: %+v", err)
// 	// 	c.JSON(http.StatusNotFound, gin.H{"message": "Audio not found"})
// 	// 	return true
// 	// }
// 	logger.ErrorContext(ctx, fmt.Sprintf("SynthesizerHandler. error: %+v", err))
// 	return false
// }

func NewInitSynthesizerRouterFunc(synthesizerUsecase SynthesizerUsecase) InitRouterGroupFunc {
	return func(parentRouterGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) error {
		workbook := parentRouterGroup.Group("synthesize")
		SynthesizerHandler := NewSynthesizerHandler(synthesizerUsecase)
		for _, m := range middleware {
			workbook.Use(m)
		}
		workbook.POST("synthesize", SynthesizerHandler.Synthesize)
		workbook.GET("audio/:audioID", SynthesizerHandler.FindAudioByID)
		return nil
	}
}
