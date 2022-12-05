package shortening

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"url-shortener/internal/model"
)

type mgo struct {
	db *mongo.Database
}

func NewMongoDB(client *mongo.Client) *mgo {
	return &mgo{db: client.Database("url-shortener")}
}

func (m *mgo) col() *mongo.Collection {
	return m.db.Collection("shortenings")
}

func (m *mgo) Put(ctx context.Context, shortenig model.Shortening) (*model.Shortening, error) {
	const op = "shortening.mgo.Put"
	shortenig.CreatedAt = time.Now().UTC()
	count, err := m.col().CountDocuments(ctx, bson.M{"_id": shortenig.Identifier})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if count > 0 {
		return nil, fmt.Errorf("%s: %w", op, model.ErrIdentifierIsExist)
	}
	_, err = m.col().InsertOne(ctx, mgoShorteningFromModel(shortenig))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &shortenig, nil
}

type mgoShortening struct {
	Identifier  string    `bson:"_id"`
	OriginalUrl string    `bson:"original_url"`
	Visits      int64     `bson:"visits"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func mgoShorteningFromModel(shortening model.Shortening) mgoShortening {
	return mgoShortening{
		Identifier:  shortening.Identifier,
		OriginalUrl: shortening.OriginalUrl,
		Visits:      shortening.Visits,
		CreatedAt:   shortening.CreatedAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}

func modelShorteningFromMgo(shortening mgoShortening) model.Shortening {
	return model.Shortening{
		Identifier:  shortening.Identifier,
		OriginalUrl: shortening.OriginalUrl,
		Visits:      shortening.Visits,
		CreatedAt:   shortening.CreatedAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}
