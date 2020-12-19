package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/o-z/featuretoggle/base/base_model"
	"github.com/o-z/featuretoggle/db"
	"github.com/o-z/featuretoggle/model/request"
	"github.com/o-z/featuretoggle/service"
)

type FeatureController struct {
	featureService service.Service
}

func NewController(db *db.Database) FeatureController {
	return FeatureController{
		featureService: service.NewService(db),
	}

}
func (controller FeatureController) GetById(c echo.Context) base_model.CustomResponse {

	id := c.Param(ID)

	featureDocument, contextError := controller.featureService.GetById(id)
	return base_model.CustomResponse{
		ResponseTime:  0,
		StatusCode:    0,
		Data:          featureDocument,
		ContextErrors: []base_model.ContextError{contextError},
	}
}

func (controller FeatureController) Save(c echo.Context) base_model.CustomResponse {
	featureRequest := new(request.FeatureRequest)
	err := c.Bind(featureRequest)
	if err != nil {
		fmt.Println(err)

	}
	featureDocument, contextError := controller.featureService.Save(*featureRequest)

	return base_model.CustomResponse{
		ResponseTime:  0,
		StatusCode:    0,
		Data:          featureDocument,
		ContextErrors: []base_model.ContextError{contextError},
	}

}

func (controller FeatureController) Update(c echo.Context) base_model.CustomResponse {
	featureRequest := new(request.FeatureRequest)
	_ = c.Bind(featureRequest)
	featureId := c.Param(NAME)

	featureDocument, _ := controller.featureService.Update(featureId, *featureRequest)

	return base_model.CustomResponse{
		ResponseTime: 0,
		StatusCode:   0,
		Data:         featureDocument,
		ContextErrors: []base_model.ContextError{{
			ErrorCode: "",
			ErrorDesc: "",
		}},
	}
}

func (controller FeatureController) DeleteWithID(c echo.Context) base_model.CustomResponse {

	featureId := c.Param(ID)

	_ = controller.featureService.DeleteWithID(featureId)
	return base_model.CustomResponse{
		ResponseTime: 0,
		StatusCode:   0,
		Data:         nil,
		ContextErrors: []base_model.ContextError{{
			ErrorCode: "",
			ErrorDesc: "",
		}},
	}
}

func (controller FeatureController) GetAll(c echo.Context) base_model.CustomResponse {
	featureName := c.QueryParam(NAME)
	featureDocuments, contextError := controller.featureService.GetAll(featureName)

	return base_model.CustomResponse{
		ResponseTime:  0,
		StatusCode:    0,
		Data:          featureDocuments,
		ContextErrors: []base_model.ContextError{contextError},
	}
}
