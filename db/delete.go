package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteById(collection string, id string) error {
	client, ctx, _, _ := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = c.DeleteOne(ctx, filter)

	return err
}
