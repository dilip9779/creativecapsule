package web

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func InternalServerError(msg string) error {
	return echo.NewHTTPError(http.StatusInternalServerError, msg)
}

func BadRequest(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"msg": msg})
}

func UnAuthorized() error {
	return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized!")
}

func MakeResponse(data interface{}, msg string) Response {
	return Response{
		Data: data,
		Msg:  msg,
	}
}

func Error(err error) error {
	if errors.Cause(err) == sql.ErrNoRows {
		return NotFound("not found")
	}

	return InternalServerError(err.Error())
}

func NotFound(msg string) error {
	return echo.NewHTTPError(http.StatusNotFound, map[string]string{"msg": msg})
}

func OK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
