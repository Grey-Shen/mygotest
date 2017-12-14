package main

import (
	"net/http"
	"time"

	"github.com/fvbock/endless"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// type AppParams struct {
// 	AppName  string `json:"app_name" binding:"required,min=4,max=6"`
// 	FullName string `json:"full_name"`
// }

type AppParams struct {
	AppName  string `json:"app_name,omitempty"`
	FullName string `json:"full_name,omitempty"`
}

func main() {
	router := gin.Default()
	router.POST("/createapp", Hello)

	endless.ListenAndServe(":8080", router)
}

func Hello(context *gin.Context) {
	var appinfo AppParams
	if err := context.BindJSON(&appinfo); err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	time.Sleep(time.Second * 1)
	context.JSON(http.StatusOK, appinfo)
}

// type ErrorHandler interface {
// 	HandleError(err *validator.FieldError) string
// }

// func RenderValidatorError(context *gin.Context, object ErrorHandler, err error) {
// 	validatorErr, ok := err.(validator.ValidationErrors)
// 	if ok {
// 		messages := make([]string, 0, len(validatorErr))
// 		for _, fieldErr := range validatorErr {
// 			errmsg := object.HandleError(fieldErr)
// 			if !arrayutils.Contains(messages, errmsg) {
// 				messages = append(messages, errmsg)
// 			}
// 		}
// 		context.JSON(http.StatusBadRequest, NewMessages(messages...))
// 	} else {
// 		context.JSON(http.StatusBadRequest, NewMessages("Bad Request Error"))
// 	}
// }

// func (_ *AppParams) HandleError(err *validator.FieldError) string {
// 	switch err.Field {
// 	case "AppName":
// 		switch err.Tag {
// 		case "required":
// 			return "app_name is missing"
// 		default:
// 			return "unknown error from app_name"
// 		}
// 	default:
// 		return fmt.Sprintf("unknown error from %s", strings.ToLower(err.Field))
// 	}
// }
