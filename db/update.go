package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateById(collection string, id string, data, result any) error {
	client, ctx, _, _ := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	opts := options.FindOneAndUpdate().SetUpsert(false)
	err = c.FindOneAndUpdate(ctx, filter, bson.M{"$set": data}, opts).Decode(result)

	if err != nil {
		return err
	}

	return c.FindOne(ctx, filter).Decode(result)

}
