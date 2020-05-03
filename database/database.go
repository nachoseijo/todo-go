package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Todo structure
type Todo struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	Completed bool               `json:"completed"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

var collection *mongo.Collection = new(mongo.Collection)

//Connect implements database connection
func Connect(collectionName, dbName, dbURI string) {
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Printf("Error creating client: %v", err)
		os.Exit(-1)
	}

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Printf("Error connecting database: %v", err)
		os.Exit(-1)
	}
	//To close the connection at the end
	defer cancel()

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	db := client.Database(dbName)

	collection = db.Collection(collectionName)
	return
}

//Delete deletes a todo from database
func Delete(todoID string) error {
	objectID, err := primitive.ObjectIDFromHex(todoID)

	if err != nil {
		log.Printf("Error while getting objectID from hex, Reason: %v\n", err)
		return nil
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})

	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		return err
	}
	return nil
}

//FindAll return all todos in database
func FindAll() []Todo {
	todos := []Todo{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return nil
	}
	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	return todos
}

//FindOne return a todo in database looking for todo ID
func FindOne(todoID string) *Todo {
	todo := &Todo{}
	objectID, err := primitive.ObjectIDFromHex(todoID)

	if err != nil {
		log.Printf("Error while getting objectID from hex, Reason: %v\n", err)
		return nil
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&todo)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil
	}

	return todo
}

//InsertOne inserts one todo into the database
func InsertOne(todo Todo) (*Todo, error) {

	newTodo := Todo{
		ID:        primitive.NewObjectID(),
		Title:     todo.Title,
		Body:      todo.Body,
		Completed: todo.Completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	log.Println(newTodo)
	_, err := collection.InsertOne(context.TODO(), newTodo)

	if err != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
		return nil, err
	}
	return &newTodo, nil
}

//UpdateOne updates todo in database and returns and error wether there is a problem updating
func UpdateOne(todoID string, todo Todo) error {

	newData := bson.M{
		"$set": bson.M{
			"title":      todo.Title,
			"body":       todo.Body,
			"completed":  todo.Completed,
			"updated_at": time.Now(),
		},
	}

	objectID, err := primitive.ObjectIDFromHex(todoID)

	if err != nil {
		log.Printf("Error while getting objectID from hex, Reason: %v\n", err)
		return nil
	}

	response, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return err
	}

	log.Println(response)
	log.Println(err)

	return nil
}
