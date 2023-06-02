package models

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
