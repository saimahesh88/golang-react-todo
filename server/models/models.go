package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"` //golang doesn't understand json, bson we need structs to define the type of the incoming object. these are the values that the database will have
	Task   string             `json:"task,omitempty"`
	Status bool               `json:status,omitempty"`
}
