package controller

import (
	"context"
	"fmt"
	"log"

	model "github.com/deep0072/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connnectionString = "mongodb://localhost:27017"

const databaseName = "Netflix"

const collectionName = "watchlist"

var collection *mongo.Collection

// connect with mongo db

func init() {

	// Connect to MongoDB
	clientOption := options.Client().ApplyURI(connnectionString) // here client refer to DB and this is used to tell that we are
	//using this connection string

	//connect to db

	client, err := mongo.Connect(context.TODO(), clientOption)
	// here context.TODO keyword is use to send cancellation signal if any error cutt off
	// context.background() is use to send cancellatio it doesnot cut the connection though

	if err != nil {
		log.Fatal(err) // it  will print the error in terminal with more info about the error as compare to panic

	}

	fmt.Println(" connection success")

	collection = client.Database(databaseName).Collection(collectionName) // db and collection are created

	fmt.Println(" instance create in mongo db")

}

//here model.Netflix is the type of the struct that is import from models.go
// we did not use here models. because in models.go file use package name "model"
func insertData(movie model.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted a single document: ", insertResult.InsertedID)
}

//update record in mongodb
func updateOneMovie(movieID string) {
	id, _:= primitive.ObjectIDFromHex(movieID) // it will convert the string to mongo object id

	filter := bson.M("_id": id) // here get the id we want to update. "_id" it is the typical id of mongodb. it is kind of getter
	//bson. M for the filter, it usually results in shorter and clearer filter declaration, the order of fields doesn't matter,

	update := bson.M("$set": bson.M("watched": true) // here we are setting the watched field to true that will be updated

	result, err := collection.UpdateOne(context.Background(), filter, update) // it will update the watched field to true

	if err ! = {
		log.Fatal(err)
	}

	fmt.Println("modified ", result.ModifiedCount, "record")

	






}
