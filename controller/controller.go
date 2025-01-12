package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title  string             `bson:"title" json:"title"`
	Status string             `bson:"status" json:"status"`
}

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connection := os.Getenv("DB_URL")
	dbname := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connection)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to MongoDB")

	collection = client.Database(dbname).Collection(collName)
	fmt.Println("Collection Instance Created")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload, err := getAllTasks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching todos: %v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payload)
}

func getAllTasks() ([]Todo, error) {
	var todos []Todo

	ctx := context.TODO()

	cursor, err := collection.Find(ctx, bson.M{}) // Fetch all documents
	if err != nil {
		return nil, fmt.Errorf("error fetching todos: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Printf("error decoding todo: %v", err)
			continue
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return todos, nil
}

// func CreateTask() {

// }

// func UpdateTask() {

// }

// func UndoTask() {

// }

// func deleteTask() {

// }

// func deleteTasks() {

// }
