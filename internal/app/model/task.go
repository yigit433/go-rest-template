package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string              `json:"title" bson:"title"`
	Description string              `json:"description" bson:"description"`
	IsCompleted bool                `json:"isCompleted" bson:"isCompleted"`
	CreatedAt   time.Time           `json:"createdAt" bson:"createdAt"`
}
