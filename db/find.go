package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(collection string, documents any) error {
	client, ctx, _, _ := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	cursor, err := c.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	defer cursor.Close(ctx)

	return cursor.All(ctx, documents)
}

func FindById(collection string, id string, document any) error {
	client, ctx, _, _ := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}

	return c.FindOne(ctx, filter).Decode(document)
}
