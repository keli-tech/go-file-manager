package handler

import (
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) (err error) {
	//nowatime := time.Now().Nanosecond()
	//return c.Render(http.StatusOK, "index", nowatime)
	return SuccessMsgResponse(c, "asdf")
}
