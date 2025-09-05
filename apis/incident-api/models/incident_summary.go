package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IncidentSummary struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Summary   string        `json:"summary"`
	Status    string        `json:"status"`
	CreatedBy string        `json:"createdBy"`
	CreatedOn time.Time     `json:"createdOn"`
}
