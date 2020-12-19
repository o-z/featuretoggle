package base_handler

import (
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties"
	"github.com/o-z/featuretoggle/base/base_model"
	"net/http"
	"time"
)

var property *properties.Properties

type CustomHandler func(c echo.Context) base_model.CustomResponse

func ServerHandler(customHandler CustomHandler) (err func(c echo.Context) (err error)) {

	return func(c echo.Context) (err error) {
		startTime := time.Now()

		response := customHandler(c)
		responseTime := time.Since(startTime).Milliseconds()
		if response.ContextErrors[0].ErrorCode != "" {
			return c.JSON(http.StatusNotFound, base_model.CustomResponse{
				StatusCode:    http.StatusNotFound,
				Data:          nil,
				ContextErrors: response.ContextErrors,
				ResponseTime:  responseTime,
			})
		} else {
			return c.JSON(http.StatusOK, base_model.CustomResponse{
				StatusCode:    http.StatusOK,
				Data:          response.Data,
				ContextErrors: nil,
				ResponseTime:  responseTime,
			})
		}

	}
}
