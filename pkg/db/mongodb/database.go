package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoClient struct {
	Client *mongo.Client
}

func MongoClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	/**
	* Verificar conexión
	 */
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}
