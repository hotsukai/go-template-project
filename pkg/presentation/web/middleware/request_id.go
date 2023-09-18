package middleware

import "github.com/labstack/echo/v4"

func SetRequestIDContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストIDを取得
		requestID := c.Response().Header().Get(echo.HeaderXRequestID)
		// リクエストIDをContextにセット
		c.Set("requestId", requestID)
		// 次のミドルウェアまたはハンドラに移動
		return next(c)
	}
}

func GetRequestId(c echo.Context) string {
	return c.Get("requestId").(string)
}
