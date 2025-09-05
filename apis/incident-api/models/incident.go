package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Incident struct {
	ID            bson.ObjectID `bson:"_id" json:"id"`
	Summary       string        `json:"summary"`
	Status        string        `json:"status"`
	CreatedBy     string        `json:"createdBy"`
	CreatedOn     bson.DateTime `json:"createdOn"`
	LastUpdatedBy string        `json:"lastUpdatedBy"`
	LastUpdatedOn bson.DateTime `json:"lastUpdatedOn"`
	Log           []IncidentLog `json:"log"`
	Resolution    Resolution    `json:"resolution"`
}
