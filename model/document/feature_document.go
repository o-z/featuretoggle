package document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FeatureDocument struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	Name       string             `json:"name",bson:"name"`
	InsertDate time.Time          `json:"insertDate",bson:"insertDate,omitempty"`
	UpdateDate time.Time          `json:"updateDate",bson:"updateDate,omitempty"`
	Status     bool               `json:"status",bson:"status"`
	StartDate  time.Time          `json:"startDate",bson:"startDate,omitempty"`
	EndDate    time.Time          `json:"endDate",bson:"endDate,omitempty"`
}
