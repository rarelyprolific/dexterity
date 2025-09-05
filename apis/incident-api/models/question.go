package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Question struct {
	Question   string        `json:"question"`
	AskedBy    string        `json:"askedBy"`
	AskedOn    bson.DateTime `json:"askedOn"`
	Answer     string        `json:"answer"`
	AnsweredBy string        `json:"answeredBy"`
	AnsweredOn bson.DateTime `json:"answeredOn"`
}
