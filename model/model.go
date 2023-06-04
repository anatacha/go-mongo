package models
import (
	"go.mongodb.org/mongo-driver/bson/primitive"

)
type Music struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type Movie struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type Game struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type Book struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string `bson:"title"`
}
