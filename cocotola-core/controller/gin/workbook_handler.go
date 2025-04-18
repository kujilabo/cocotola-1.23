package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"
	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/controller/gin/helper"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
)

const defaultPageSize = 10

type WorkbookQueryUsecase interface {
	FindWorkbooks(ctx context.Context, operator service.OperatorInterface, param *libapi.WorkbookFindParameter) (*libapi.WorkbookFindResult, error)

	RetrieveWorkbookByID(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID) (*libapi.WorkbookRetrieveResult, error)
}

type WorkbookCommandUsecase interface {
	AddWorkbook(ctx context.Context, operator service.OperatorInterface, param *service.WorkbookAddParameter) (*domain.WorkbookID, error)
	UpdateWorkbook(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID, version int, param *service.WorkbookUpdateParameter) error
}

type WorkbookHandler struct {
	workbookQueryUsecase   WorkbookQueryUsecase
	workbookCommandUsecase WorkbookCommandUsecase
	logger                 *slog.Logger
}

func NewWorkbookHandler(workbookQueryUsecase WorkbookQueryUsecase, workbookCommandUsecase WorkbookCommandUsecase) *WorkbookHandler {
	return &WorkbookHandler{
		workbookQueryUsecase:   workbookQueryUsecase,
		workbookCommandUsecase: workbookCommandUsecase,
		logger:                 slog.Default().With(slog.String(rsliblog.LoggerNameKey, "WorkbookHandler")),
	}
}

func (h *WorkbookHandler) FindWorkbooks(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		param := libapi.WorkbookFindParameter{
			PageNo:   1,
			PageSize: defaultPageSize,
		}
		result, err := h.workbookQueryUsecase.FindWorkbooks(ctx, operator, &param)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, result)
		return nil
	}, h.errorHandle)
}

// func (h *WorkbookHandler) toWorkbookFindResultEntity(model *studentusecase.WorkbookFindResult) *WorkbookFindResult {
// 	results := make([]*WorkbookFindModel, len(model.Results))
// 	for i, r := range model.Results {
// 		results[i] = &WorkbookFindModel{ID: r.ID, Name: r.Name}
// 	}

// 	return &WorkbookFindResult{
// 		TotalCount: model.TotalCount,
// 		Results:    results,
// 	}
// }

func (h *WorkbookHandler) RetrieveWorkbookByID(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		workbookIDInt, err := helper.GetIntFromPath(c, "workbookID")
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("GetIntFromPath. err: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}

		workbookID, err := domain.NewWorkbookID(workbookIDInt)
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("NewWorkbookID. err: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}

		result, err := h.workbookQueryUsecase.RetrieveWorkbookByID(ctx, operator, workbookID)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, result)
		return nil
	}, h.errorHandle)
}

func (h *WorkbookHandler) AddWorkbook(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		apiParam := libapi.WorkbookAddParameter{}
		if err := c.ShouldBindJSON(&apiParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return nil
		}

		param := service.WorkbookAddParameter{
			Name:        apiParam.Name,
			Lang2:       apiParam.Lang2,
			ProblemType: apiParam.ProblemType,
			Description: apiParam.Description,
			Content:     apiParam.Content,
		}
		workbookID, err := h.workbookCommandUsecase.AddWorkbook(ctx, operator, &param)
		if err != nil {
			h.logger.ErrorContext(ctx, fmt.Sprintf("workbookCommandUsecase.AddWorkbook. err: %+v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return nil
		}

		c.JSON(http.StatusOK, gin.H{"id": workbookID.Int()})
		return nil
	}, h.errorHandle)
}

func (h *WorkbookHandler) UpdateWorkbook(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		version, err := helper.GetIntFromQuery(c, "version")
		if err != nil {
			return rslibdomain.ErrInvalidArgument
		}

		workbookID, err := helper.GetWorkbookIDFromPath(c, "workbookID")
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("GetIntFromPath. err: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}

		apiParam := libapi.WorkbookUpdateParameter{}
		if err := c.ShouldBindJSON(&apiParam); err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("ShouldBindJSON. err: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}

		param := service.WorkbookUpdateParameter{
			Name:        apiParam.Name,
			Description: apiParam.Description,
			Content:     apiParam.Content,
		}

		if err := h.workbookCommandUsecase.UpdateWorkbook(ctx, operator, workbookID, version, &param); err != nil {
			return rsliberrors.Errorf("workbookCommandUsecase.UpdateWorkbook. err: %w", err)
		}

		c.Status(http.StatusOK)
		return nil
	}, h.errorHandle)
}

// func (h *WorkbookHandler) toWorkbookRetrieveResultEntity(model *workbookretrievedomain.WorkbookModel) *WorkbookWithProblem {
// 	problems := make([]*Problem, len(model.Problems))
// 	for i, r := range model.Problems {
// 		problems[i] = &Problem{
// 			Type:       r.Type,
// 			Properties: r.Properties,
// 		}
// 	}

// 	return &WorkbookWithProblem{
// 		ID:       model.ID,
// 		Problems: problems,
// 	}
// }

func (h *WorkbookHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	if errors.Is(err, rslibdomain.ErrInvalidArgument) {
		h.logger.WarnContext(ctx, fmt.Sprintf("PrivateWorkbookHandler err: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return true
	}
	if errors.Is(err, service.ErrWorkbookNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
		return true
	}
	h.logger.ErrorContext(ctx, fmt.Sprintf("WorkbookHandler. error: %+v", err))
	return false
}

func NewInitWorkbookRouterFunc(workbookQueryUsecase WorkbookQueryUsecase, workbookCommandUsecase WorkbookCommandUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		workbook := parentRouterGroup.Group("workbook")
		workbookHandler := NewWorkbookHandler(workbookQueryUsecase, workbookCommandUsecase)
		for _, m := range middleware {
			workbook.Use(m)
		}
		workbook.GET("", workbookHandler.FindWorkbooks)
		workbook.GET(":workbookID", workbookHandler.RetrieveWorkbookByID)
		// workbook.POST(":workbookID", privateWorkbookHandler.FindWorkbooks)
		// workbook.GET(":workbookID", privateWorkbookHandler.FindWorkbookByID)
		workbook.PUT(":workbookID", workbookHandler.UpdateWorkbook)
		// workbook.DELETE(":workbookID", privateWorkbookHandler.RemoveWorkbook)
		workbook.POST("", workbookHandler.AddWorkbook)
	}
}
