package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IncidentLog struct {
	Description string        `json:"description"`
	CreatedBy   string        `json:"createdBy"`
	CreatedOn   bson.DateTime `json:"createdOn"`
	Questions   []Question    `json:"questions"`
}
