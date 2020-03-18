package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func CheckLogin(c echo.Context) bool {
	//loginSession := service.GetLoginInfo(c)
	//if loginSession.UserID <= 0 {
	FailedMsgResponse(c, 401, "请重新登录!")
	//return false
	//}
	return true
}

func SuccessMsgResponse(c echo.Context, msg string) error {
	res := &Response{Code: 200, Message: msg}
	return c.JSON(http.StatusOK, res)
}

func FailedMsgResponse(c echo.Context, code int, msg string) error {
	res := &Response{Code: code, Message: msg}
	return c.JSON(http.StatusOK, res)
}

func SuccessResponse(c echo.Context, msg string, data interface{}) error {
	res := &Response{Code: 200, Message: msg, Data: data}
	return c.JSON(http.StatusOK, res)
}

func FailedResponse(c echo.Context, code int, msg string, data interface{}) error {
	res := &Response{Code: code, Message: msg, Data: data}
	return c.JSON(http.StatusOK, res)
}
