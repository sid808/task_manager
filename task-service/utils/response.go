package utils

import (
	"net/http"
	"task-service/models"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, err string) {
	c.JSON(code, models.GenericResponse{
		Success: false,
		Error:   err,
	})
}

func RespondWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.GenericResponse{
		Success: true,
		Data:    data,
	})
}
