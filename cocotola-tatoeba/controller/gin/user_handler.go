package handler

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	handlerhelper "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/controller/gin/helper"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

type UserUsecase interface {
	FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchCondition) (service.TatoebaSentencePairSearchResult, error)

	FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (service.TatoebaSentence, error)
}

type UserHandler struct {
	userUsecase UserUsecase
}

func NewUserHandler(userUsecase UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) logger() *slog.Logger {
	return slog.Default().With(slog.String(rsliblog.LoggerNameKey, "tatoeba.UserHandler"))
}

// FindSentencePairs godoc
// @Summary     find pair of sentences
// @Description find pair of sentences
// @Tags        tatoeba
// @Accept      json
// @Produce     json
// @Param       param body entity.TatoebaSentenceFindParameter true "parameter to find sentences"
// @Success     200 {object} entity.TatoebaSentencePairFindResponse
// @Failure     400
// @Failure     401
// @Router      /v1/user/sentence_pair/find [post]
// @Security    BasicAuth
func (h *UserHandler) FindSentencePairs(c *gin.Context) {
	handlerhelper.HandleFunction(c, func(ctx context.Context) error {
		// param := entity.TatoebaSentenceFindParameter{}
		// if err := c.ShouldBindJSON(&param); err != nil {
		// 	c.Status(http.StatusBadRequest)
		// 	return nil
		// }
		// logger.Debugf("FindSentencePairs. param: %+v", param)
		// parameter, err := converter.ToTatoebaSentenceSearchCondition(ctx, &param)
		// if err != nil {
		// 	return liberrors.Errorf("convert parameter to TatoebaSentenceSearchCondition. err: %w", err)
		// }
		// result, err := h.userUsecase.FindSentencePairs(ctx, parameter)
		// if err != nil {
		// 	return liberrors.Errorf("execute FindSentencePairs. err: %w", err)
		// }
		// response, err := converter.ToTatoebaSentenceFindResponse(ctx, result)
		// if err != nil {
		// 	return liberrors.Errorf("convert result to TatoebaSentenceFindResponse. err: %w", err)
		// }

		// c.JSON(http.StatusOK, response)
		return nil
	}, h.errorHandle)
}

// FindSentenceBySentenceNumber godoc
// @Summary     import links
// @Description import links
// @Tags        tatoeba
// @Accept      json
// @Produce     json
// @Param       sentenceNumber path int true "Sentence number"
// @Success     200 {object} entity.TatoebaSentenceResponse
// @Failure     400
// @Failure     401
// @Router      /v1/user/sentence/{sentenceNumber} [get]
// @Security    BasicAuth
func (h *UserHandler) FindSentenceBySentenceNumber(c *gin.Context) {
	handlerhelper.HandleFunction(c, func(ctx context.Context) error {
		// sentenceNumber, err := helper.GetIntFromPath(c, "sentenceNumber")
		// if err != nil {
		// 	return rslibdomain.ErrInvalidArgument
		// }

		// result, err := h.userUsecase.FindSentenceBySentenceNumber(ctx, sentenceNumber)
		// if err != nil {
		// 	return liberrors.Errorf("execute FindSentenceBySentenceNumber. err: %w", err)
		// }
		// response, err := converter.ToTatoebaSentenceResponse(ctx, result)
		// if err != nil {
		// 	return liberrors.Errorf("convert result to TatoebaSentenceResponse. err: %w", err)
		// }

		// c.JSON(http.StatusOK, response)
		return nil
	}, h.errorHandle)
}

func (h *UserHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	h.logger().ErrorContext(ctx, fmt.Sprintf("userHandler. err: %+v", err))
	return false
}
