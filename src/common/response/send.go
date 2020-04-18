package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataList struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

func SendSuccess(Context *gin.Context, data interface{}) {
	response := gin.H{
		"success": true,
		"data":    data,
	}

	Context.JSON(http.StatusOK, response)
}

func SendError(Context *gin.Context, error error) {
	response := gin.H{
		"success": false,
		"message": error.Error(),
	}

	Context.JSON(http.StatusOK, response)
}
