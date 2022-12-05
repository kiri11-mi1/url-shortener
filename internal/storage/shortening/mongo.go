package shortening

import "go.mongodb.org/mongo-driver/mongo"

type mgo struct {
	db *mongo.Database
}

func NewMongoDB(client *mongo.Client) *mgo {
	return &mgo{db: client.Database("url-shortener")}
}

func (m *mgo) col() *mongo.Collection {
	return m.db.Collection("shortenings")
}
