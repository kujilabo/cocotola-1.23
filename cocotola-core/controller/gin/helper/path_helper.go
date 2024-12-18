package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/domain"
)

func GetIntFromPath(c *gin.Context, param string) (int, error) {
	idS := c.Param(param)
	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetStringFromPath(c *gin.Context, param string) string {
	return c.Param(param)
}

func GetWorkbookIDFromPath(c *gin.Context, param string) (*domain.WorkbookID, error) {
	value, err := GetIntFromPath(c, param)
	if err != nil {
		return nil, err
	}

	workbookID, err := domain.NewWorkbookID(value)
	if err != nil {
		return nil, err
	}

	return workbookID, nil
}
