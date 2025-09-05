package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Resolution struct {
	Description string        `json:"description"`
	ResolvedBy  string        `json:"resolvedBy"`
	ResolvedOn  bson.DateTime `json:"resolvedOn"`
}
