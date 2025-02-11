package controller

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	controllerhelper "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/controller/gin/helper"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

type AdminUsecase interface {
	ImportSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator) (*domain.ImportResult, error)

	ImportLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator) (*domain.ImportResult, error)
}

type AdminHandler struct {
	adminUsecase                         AdminUsecase
	newTatoebaSentenceAddParameterReader func(reader io.Reader) service.TatoebaSentenceAddParameterIterator
	newTatoebaLinkAddParameterReader     func(reader io.Reader) service.TatoebaLinkAddParameterIterator
	logger                               *slog.Logger
}

func NewAdminHandler(adminUsecase AdminUsecase, newTatoebaSentenceAddParameterReader func(reader io.Reader) service.TatoebaSentenceAddParameterIterator, newTatoebaLinkAddParameterReader func(reader io.Reader) service.TatoebaLinkAddParameterIterator) *AdminHandler {
	return &AdminHandler{
		adminUsecase:                         adminUsecase,
		newTatoebaSentenceAddParameterReader: newTatoebaSentenceAddParameterReader,
		newTatoebaLinkAddParameterReader:     newTatoebaLinkAddParameterReader,
		logger:                               slog.Default().With(slog.String(rsliblog.LoggerNameKey, "tatoeba.AdminHandler")),
	}
}

// func (h *AdminHandler) logger() *slog.Logger {
// 	return slog.Default().With(slog.String(rsliblog.LoggerNameKey, "tatoeba.AdminHandler"))
// }

// ImportSentences godoc
// @Summary     import sentences
// @Description import sentences
// @Tags        tatoeba
// @Param       file formData file true "***_sentences_detailed.tsv"
// @Success     200
// @Failure     400
// @Failure     401
// @Failure     500
// @Router      /v1/admin/sentence/import [post]
// @Security    BasicAuth
func (h *AdminHandler) ImportSentences(c *gin.Context) {
	controllerhelper.HandleFunction(c, func(ctx context.Context) error {
		h.logger.InfoContext(ctx, "ImportSentences")
		file, err := c.FormFile("file")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				h.logger.WarnContext(ctx, fmt.Sprintf("err: %+v", err))
				c.Status(http.StatusBadRequest)
				return nil
			}
			if errors.Is(err, http.ErrNotMultipart) {
				h.logger.WarnContext(ctx, fmt.Sprintf("err: %+v", err))
				c.Status(http.StatusBadRequest)
				return nil
			}
			return err
		}

		multipartFile, err := file.Open()
		if err != nil {
			return rsliberrors.Errorf("failed to file.Open. err: %w", err)
		}
		defer multipartFile.Close()

		iterator := h.newTatoebaSentenceAddParameterReader(multipartFile)

		importResult, err := h.adminUsecase.ImportSentences(ctx, iterator)
		if err != nil {
			return rsliberrors.Errorf("failed to ImportSentences. err: %w", err)
		}

		c.JSON(http.StatusOK, importResult)
		return nil
	}, h.errorHandle)
}

// ImportLinks godoc
// @Summary     import links
// @Description import links
// @Tags        tatoeba
// @Param       file formData file true "links.csv"
// @Success     200
// @Failure     400
// @Failure     401
// @Failure     500
// @Router      /v1/admin/link/import [post]
// @Security    BasicAuth
func (h *AdminHandler) ImportLinks(c *gin.Context) {
	controllerhelper.HandleFunction(c, func(ctx context.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				h.logger.WarnContext(ctx, fmt.Sprintf("err: %+v", err))
				c.Status(http.StatusBadRequest)
				return nil
			}
			if errors.Is(err, http.ErrNotMultipart) {
				h.logger.WarnContext(ctx, fmt.Sprintf("err: %+v", err))
				c.Status(http.StatusBadRequest)
				return nil
			}
			return err
		}

		multipartFile, err := file.Open()
		if err != nil {
			return rsliberrors.Errorf("failed to file.Open. err: %w", err)
		}
		defer multipartFile.Close()

		iterator := h.newTatoebaLinkAddParameterReader(multipartFile)

		importResult, err := h.adminUsecase.ImportLinks(ctx, iterator)
		if err != nil {
			return rsliberrors.Errorf("failed to ImportLinks. err: %w", err)
		}

		c.JSON(http.StatusOK, importResult)
		return nil
	}, h.errorHandle)
}

func (h *AdminHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	h.logger.ErrorContext(ctx, fmt.Sprintf("adminHandler. err: %+v", err))
	return false
}

func NewInitAdminRouterFunc(adminUsecase AdminUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		admin := parentRouterGroup.Group("admin")
		newSentenceReader := func(reader io.Reader) service.TatoebaSentenceAddParameterIterator {
			return gateway.NewTatoebaSentenceAddParameterReader(reader)
		}
		newLinkReader := func(reader io.Reader) service.TatoebaLinkAddParameterIterator {
			return gateway.NewTatoebaLinkAddParameterReader(reader)
		}
		adminHandler := NewAdminHandler(adminUsecase, newSentenceReader, newLinkReader)
		for _, m := range middleware {
			admin.Use(m)
		}
		admin.POST("sentence/import", adminHandler.ImportSentences)
		admin.POST("link/import", adminHandler.ImportLinks)
	}
}
