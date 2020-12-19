package main

import (
	"github.com/labstack/echo/v4"
	"github.com/o-z/featuretoggle/base/context_error"
	"github.com/o-z/featuretoggle/config"
	"github.com/o-z/featuretoggle/db"
	"github.com/o-z/featuretoggle/route"
)

func init() {
	context_error.SetAllProperties("./error/*.properties")
}
func main() {
	mongoDB, _ := db.Get(config.Get())
	e := echo.New()
	router := &route.FeatureRouterType{
		Echo:         e,
		DBConnection: mongoDB,
	}
	route.InitFeatureRouter(router)

	e.Logger.Info(e.Start(":8089"))
}
