package controller

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	handlerhelper "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/controller/gin/helper"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

type TatoebaSentenceFindParameter struct {
	PageNo   int    `form:"pageNo" json:"pageNo" binding:"required,gte=1"`
	PageSize int    `form:"pageSize" json:"pageSize" binding:"required,gte=1"`
	Keyword  string `form:"keyword" json:"keyword"`
	Random   bool   `form:"random" json:"random"`
}

type TatoebaSentenceResponse struct {
	SentenceNumber int       `json:"sentenceNumber"`
	Lang2          string    `json:"lang2" binding:"len=2" validate:"oneof=ja en"`
	Text           string    `json:"text"`
	Author         string    `json:"author"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type TatoebaSentencePair struct {
	Src TatoebaSentenceResponse `json:"src"`
	Dst TatoebaSentenceResponse `json:"dst"`
}

type TatoebaSentencePairFindResponse struct {
	TotalCount int                   `json:"totalCount"`
	Results    []TatoebaSentencePair `json:"results"`
}

func ToTatoebaSentenceSearchCondition(ctx context.Context, param *TatoebaSentenceFindParameter) (*service.TatoebaSentenceSearchCondition, error) {
	return service.NewTatoebaSentenceSearchCondition(param.PageNo, param.PageSize, param.Keyword, param.Random)
}

func ToTatoebaSentenceFindResponse(ctx context.Context, result *service.TatoebaSentencePairSearchResult) (*TatoebaSentencePairFindResponse, error) {
	entities := make([]TatoebaSentencePair, len(result.Results))
	for i, m := range result.Results {
		src := TatoebaSentenceResponse{
			SentenceNumber: m.Src.SentenceNumber,
			Lang2:          m.Src.Lang3.ToLang2().String(),
			Text:           m.Src.Text,
			Author:         m.Src.Author,
			UpdatedAt:      m.Src.UpdatedAt,
		}
		if err := rslibdomain.Validator.Struct(src); err != nil {
			return nil, err
		}

		dst := TatoebaSentenceResponse{
			SentenceNumber: m.Dst.SentenceNumber,
			Lang2:          m.Dst.Lang3.ToLang2().String(),
			Text:           m.Dst.Text,
			Author:         m.Dst.Author,
			UpdatedAt:      m.Dst.UpdatedAt,
		}
		if err := rslibdomain.Validator.Struct(dst); err != nil {
			return nil, err
		}

		entities[i] = TatoebaSentencePair{
			Src: src,
			Dst: dst,
		}
	}

	return &TatoebaSentencePairFindResponse{
		TotalCount: result.TotalCount,
		Results:    entities,
	}, nil
}

type UserUsecase interface {
	FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) (*service.TatoebaSentencePairSearchResult, error)

	FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (*service.TatoebaSentence, error)
}

type UserHandler struct {
	userUsecase UserUsecase
	logger      *slog.Logger
}

func NewUserHandler(userUsecase UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		logger:      slog.Default().With(slog.String(rsliblog.LoggerNameKey, "tatoeba.UserHandler")),
	}
}

// func (h *UserHandler) logger() *slog.Logger {
// 	return slog.Default().With(slog.String(rsliblog.LoggerNameKey, "tatoeba.UserHandler"))
// }

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
		param := TatoebaSentenceFindParameter{}
		if err := c.ShouldBind(&param); err != nil {
			h.logger.InfoContext(ctx, fmt.Sprintf("FindSentencePairs. err: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}
		h.logger.DebugContext(ctx, fmt.Sprintf("FindSentencePairs. param: %+v", param))
		parameter, err := ToTatoebaSentenceSearchCondition(ctx, &param)
		if err != nil {
			return rsliberrors.Errorf("convert parameter to TatoebaSentenceSearchCondition. err: %w", err)
		}
		result, err := h.userUsecase.FindSentencePairs(ctx, parameter)
		if err != nil {
			return rsliberrors.Errorf("execute FindSentencePairs. err: %w", err)
		}
		response, err := ToTatoebaSentenceFindResponse(ctx, result)
		if err != nil {
			return rsliberrors.Errorf("convert result to TatoebaSentenceFindResponse. err: %w", err)
		}

		c.JSON(http.StatusOK, response)
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
	h.logger.ErrorContext(ctx, fmt.Sprintf("userHandler. err: %+v", err))
	return false
}

func NewInitUserRouterFunc(userUsecase UserUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		user := parentRouterGroup.Group("user")
		userHandler := NewUserHandler(userUsecase)
		for _, m := range middleware {
			user.Use(m)
		}
		user.GET("sentence_pair/find", userHandler.FindSentencePairs)
	}
}
