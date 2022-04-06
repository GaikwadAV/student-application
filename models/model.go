package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Student ...
// swagger:model
type Student struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name            string             `bson:"name,omitempty" json:"name,omitempty"`
	City            string             `bson:"city,omitempty" json:"city,omitempty"`
	Country         string             `bson:"country,omitempty" json:"country,omitempty"`
	Course          string             `bson:"course,omitempty" json:"course,omitempty"`
	YearOfAdmission int                `bson:"YearOfAdmission,omitempty" json:"Year_Of_Admission,omitempty"`
}
