package db

import "go.mongodb.org/mongo-driver/bson/primitive"

func Insert(collection string, data any) (primitive.ObjectID, error) {
	client, ctx, _, err := getConnection()
	defer client.Disconnect(ctx)

	if err != nil {
		return primitive.NilObjectID, err
	}

	coll := client.Database(dbname).Collection(collection)
	res, err := coll.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
