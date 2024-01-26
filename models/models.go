package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// creating  to do list model
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}
