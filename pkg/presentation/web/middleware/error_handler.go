package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"sample/pkg/util/app_error"
	"sample/pkg/util/log"
)

type errDto struct {
	StatusCode int    `json:"code"`
	VisibleMsg string `json:"message"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	requestId := GetRequestId(c)
	if me, ok := app_error.IsAppError(err); ok {
		log.AppLogger().Errorw(me.Error(), zap.String("requestId", requestId))
		c.JSON(me.StatusCode, errDto{StatusCode: me.StatusCode, VisibleMsg: me.VisibleMsg})
		return
	}
	if he, ok := err.(*echo.HTTPError); ok {
		log.AppLogger().Errorw("Unknown HTTP error: %v", he.Error(), zap.String("requestId", requestId))
		c.JSON(he.Code, errDto{StatusCode: he.Code, VisibleMsg: "something wrong"})
		return
	}

	log.AppLogger().Errorw("Unknown  error: %v", err.Error(), zap.String("requestId", requestId))
	c.JSON(http.StatusInternalServerError, errDto{StatusCode: http.StatusInternalServerError, VisibleMsg: "something wrong"})
}
