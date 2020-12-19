package repository

import (
	"context"
	"github.com/o-z/featuretoggle/base/base_model"
	"github.com/o-z/featuretoggle/base/context_error"
	"github.com/o-z/featuretoggle/db"
	"github.com/o-z/featuretoggle/model/document"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FeatureRepository struct {
	Collection *mongo.Collection
}

func NewRepository(mongoDB *db.Database) FeatureRepository {
	return FeatureRepository{Collection: NewCollection(mongoDB)}
}

func NewCollection(database *db.Database) *mongo.Collection {
	return database.MongoDB.Collection("FeatureDocument")

}

func (repository *FeatureRepository) GetByFilter(id string) (document.FeatureDocument, base_model.ContextError) {
	var feature document.FeatureDocument
	filter := bson.M{"_id": id}
	err := repository.Collection.FindOne(context.TODO(), filter).Decode(&feature)
	var contextError base_model.ContextError
	if err != nil {
		contextError = context_error.GetContextError("feature_not_found")
	}
	return feature, contextError
}

func (repository *FeatureRepository) Save(feature document.FeatureDocument) (document.FeatureDocument, base_model.ContextError) {

	feature.InsertDate = time.Now()
	result, err := repository.Collection.InsertOne(context.TODO(), feature)
	var contextError base_model.ContextError
	if err != nil {
		contextError = context_error.GetContextError("feature_not_saved")
	}

	feature.ID = result.InsertedID.(primitive.ObjectID)
	return feature, contextError
}

func (repository *FeatureRepository) Update(featureId string, feature document.FeatureDocument) (document.FeatureDocument, error) {
	objectID, _ := primitive.ObjectIDFromHex(featureId)
	feature.ID = objectID
	filter := bson.M{"_id": objectID}
	updateFeature := bson.M{
		"$set": feature,
	}
	_, err := repository.Collection.UpdateOne(context.TODO(), filter, updateFeature)
	return feature, err
}

func (repository *FeatureRepository) DeleteWithID(featureId string) error {
	var feature document.FeatureDocument
	objectID, _ := primitive.ObjectIDFromHex(featureId)
	feature.ID = objectID
	filter := bson.M{"_id": objectID}
	_, err := repository.Collection.DeleteOne(context.TODO(), filter)
	return err
}
func (repository *FeatureRepository) GetAll(featureName string) ([]document.FeatureDocument, base_model.ContextError) {
	var feature []document.FeatureDocument
	filter := bson.D{{}}

	if featureName != "" {
		filter = bson.D{{Key: "name", Value: featureName}}
	}
	cursor, err := repository.Collection.Find(context.TODO(), filter)
	var contextError base_model.ContextError
	if err != nil {
		contextError = context_error.GetContextError("feature_not_found")
	}
	err = cursor.All(context.TODO(), &feature)
	if err != nil {
		contextError = context_error.GetContextError("feature_not_found")
	}

	if feature == nil {
		contextError = context_error.GetContextError("feature_not_found")
	}
	return feature, contextError
}
