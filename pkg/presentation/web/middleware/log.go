package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var responseLogFormat = `{` +
	`"time":"${time_custom}",` +
	`"id":"${id}",` +
	`"remote_ip":"${remote_ip}",` +
	`"host":"${host}",` +
	`"method":"${method}",` +
	`"uri":"${uri}",` +
	`"user_agent":"${user_agent}",` +
	`"status":${status},` +
	`"error":"${error}",` +
	`"latency":${latency},` +
	`"latency_human":"${latency_human}",` +
	`"bytes_in":${bytes_in},` +
	`"bytes_out":${bytes_out},` +
	`"forwarded-for":"${header:x-forwarded-for}",` +
	`"same-as-id":${header:X-Request-Id},` +
	`"query":${query:lang}` +
	`}`

func CustomLogger() middleware.LoggerConfig {
	cl := middleware.DefaultLoggerConfig
	cl.Skipper = customSkipper
	cl.Format = responseLogFormat
	cl.CustomTimeFormat = "2006/01/02 15:04:05.00000"

	return cl
}

func customSkipper(c echo.Context) bool {
	if c.Path() == "/healthcheck" {
		return true
	}
	return false
}
