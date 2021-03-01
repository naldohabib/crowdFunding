package utils

import (
	"crowFunding/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleSuccess(c *gin.Context, message string, data interface{}) {
	meta := model.Meta{
		Message: message,
		Code:    http.StatusOK,
		Status:  true,
	}

	jsonResponse := model.Respons{
		Meta: meta,
		Data: data,
	}

	c.JSON(http.StatusOK, jsonResponse)

}

func HandleError(c *gin.Context, code int, message string, ) {
	meta := model.Meta{
		Code:    code,
		Status:  false,
		Message: message,
	}

	jsonResponse := model.Respons{
		Meta: meta,
	}

	c.JSON(code, jsonResponse)
}

//func HandleSucces(c *gin.Context, data interface{}) {
//	var returnData = model.Respons{
//		Meta: model.Meta{},
//		//Success: true,
//		//Message: "Success",
//		Data: data,
//	}
//	c.JSON(http.StatusOK, returnData)
//}

//func HandleError(c *gin.Context, status int, message string) {
//	var returnData = model.Respons{
//		Success: false,
//		Message: message,
//	}
//	c.JSON(status, returnData)
//}