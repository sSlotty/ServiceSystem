package service

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//check connection mongoDB
func CheckConnection(uri string) bool {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)

	}
	return true
}

func MongoDBConnection() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	return client
}

//save data to mongoDB
func SaveData(collection *mongo.Collection, data interface{}) bool {
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

//get data from mongoDB
func GetData(collection *mongo.Collection, data interface{}) bool {
	err := collection.FindOne(context.TODO(), data).Decode(&data)
	if err != nil {
		panic(err)
	}
	return true
}

//update data to mongoDB
func UpdateData(collection *mongo.Collection, data interface{}) bool {
	updateResult, err := collection.UpdateOne(context.TODO(), data, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}

//delete data to mongoDB
func DeleteData(collection *mongo.Collection, data interface{}) bool {
	deleteResult, err := collection.DeleteOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted a single document: ", deleteResult.DeletedCount)
	return true
}

//get all data from mongoDB
func GetAllData(collection *mongo.Collection) []interface{} {
	var data []interface{}
	cur, err := collection.Find(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			panic(err)
		}
		data = append(data, result)
	}
	return data
}
