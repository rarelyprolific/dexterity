package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Page struct {
	ID            bson.ObjectID `bson:"_id" json:"id"`
	Title         string        `json:"title"`
	ShortLink     string        `json:"shortLink"`
	CreatedBy     string        `json:"createdBy"`
	CreatedOn     bson.DateTime `json:"createdOn"`
	LastUpdatedBy string        `json:"lastUpdatedBy"`
	LastUpdatedOn bson.DateTime `json:"lastUpdatedOn"`
	Text          string        `json:"text"`
	Tags          []string      `json:"tags"`
}
