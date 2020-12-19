package service

import (
	"github.com/o-z/featuretoggle/base/base_model"
	"github.com/o-z/featuretoggle/db"
	"github.com/o-z/featuretoggle/model/document"
	"github.com/o-z/featuretoggle/model/request"
	"github.com/o-z/featuretoggle/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FeatureService struct {
	featureRepository repository.FeatureRepository
}

type Service interface {
	GetById(name string) (document.FeatureDocument, base_model.ContextError)
	Save(featureRequest request.FeatureRequest) (document.FeatureDocument, base_model.ContextError)
	Update(featureId string, featureRequest request.FeatureRequest) (document.FeatureDocument, error)
	DeleteWithID(id string) error
	GetAll(featureName string) ([]document.FeatureDocument, base_model.ContextError)
}

func NewService(db *db.Database) FeatureService {
	return FeatureService{
		featureRepository: repository.NewRepository(db),
	}
}

func (featureService FeatureService) GetById(id string) (document.FeatureDocument, base_model.ContextError) {
	return featureService.featureRepository.GetByFilter(id)

}

func (featureService FeatureService) Save(featureRequest request.FeatureRequest) (document.FeatureDocument, base_model.ContextError) {
	featureDocument := new(document.FeatureDocument)
	featureDocument.ID = primitive.NewObjectID()
	featureDocument.Name = featureRequest.Name
	featureDocument.InsertDate = time.Now()
	featureDocument.Status = featureRequest.Status
	featureDocument.StartDate = featureRequest.StartDate
	featureDocument.EndDate = featureRequest.EndDate
	return featureService.featureRepository.Save(*featureDocument)

}

func (featureService FeatureService) Update(featureId string, featureRequest request.FeatureRequest) (document.FeatureDocument, error) {
	featureDocument := new(document.FeatureDocument)
	featureDocument.Name = featureRequest.Name
	featureDocument.UpdateDate = time.Now()
	featureDocument.Status = featureRequest.Status
	featureDocument.StartDate = featureRequest.StartDate
	featureDocument.EndDate = featureRequest.EndDate
	return featureService.featureRepository.Update(featureId, *featureDocument)

}

func (featureService FeatureService) DeleteWithID(id string) error {
	return featureService.featureRepository.DeleteWithID(id)

}

func (featureService FeatureService) GetAll(featureName string) ([]document.FeatureDocument, base_model.ContextError) {
	return featureService.featureRepository.GetAll(featureName)

}
