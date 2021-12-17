package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Score struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Point float32            `bson:"point" json:"point"`
	Name  string             `bson:"name" json:"name"`
}

func InitScore(ID primitive.ObjectID, Point float32, Name string) *Score {
	return &Score{
		ID:    ID,
		Point: Point,
		Name:  Name,
	}
}
