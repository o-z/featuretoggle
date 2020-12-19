package route

import "C"
import (
	"github.com/labstack/echo/v4"
	"github.com/o-z/featuretoggle/base/base_handler"
	"github.com/o-z/featuretoggle/controller"
	"github.com/o-z/featuretoggle/db"
)

type FeatureRouterType struct {
	Echo              *echo.Echo
	FeatureController controller.FeatureController
	DBConnection      *db.Database
}

func InitFeatureRouter(route *FeatureRouterType) {
	route.FeatureController = controller.NewController(route.DBConnection)
	route.Echo.GET(controller.FEATURE_BY_ID, base_handler.ServerHandler(route.FeatureController.GetById))
	route.Echo.POST(controller.FEATURES, base_handler.ServerHandler(route.FeatureController.Save))
	route.Echo.PUT(controller.FEATURE_BY_ID, base_handler.ServerHandler(route.FeatureController.Update))
	route.Echo.DELETE(controller.FEATURE_BY_ID, base_handler.ServerHandler(route.FeatureController.DeleteWithID))
	route.Echo.GET(controller.FEATURES, base_handler.ServerHandler(route.FeatureController.GetAll))

}
