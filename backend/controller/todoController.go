package controller

import (
	"context"
	db "go-todo/backend/db"
	model "go-todo/backend/model"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/http"

	"github.com/labstack/echo"
)

func GetTask(c echo.Context) (err error) {

	task := new(model.Task)
	taskID := c.Param("id")

	collection, err := db.Db()
	if err != nil {
		log.Fatal(err)
	}

	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		log.Warn(err)
	}
	filter := bson.M{"_id": objectID}
	err = collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		log.Warn(err)
	}

	return c.JSON(http.StatusOK, task)
}

func GetTasks(c echo.Context) (err error) {

	// var tasks = []model.Task{}
	tasks := new([]model.Task)

	collection, err := db.Db()
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{}
	result, err := collection.Find(context.TODO(), filter)
	_ = result
	if err != nil {
		log.Fatal(err)
	}

	if err = result.All(context.TODO(), tasks); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, tasks)
}

func AddTask(c echo.Context) error {

	task := new(model.Task)

	collection, err := db.Db()
	if err != nil {
		log.Fatal(err)
	}

	if err = c.Bind(task); err != nil {
		log.Warn(err)
	}

	// task := model.Task{"1", "test1", false}
	// task.IDTask = MakeRandomString()
	task.ID = primitive.NewObjectID()
	task.Done = false

	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusCreated, result) //result.InsertedID
}

func DeleteTask(c echo.Context) error {

	taskID := c.Param("id")

	collection, err := db.Db()
	if err != nil {
		log.Fatal(err)
	}

	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		log.Warn(err)
	}

	filter := bson.M{"_id": objectID}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, result) //result.DeletedCount
}

func MarkTaskDone(c echo.Context) error {

	taskID := c.Param("id")

	collection, err := db.Db()
	if err != nil {
		log.Fatal(err)
	}

	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		log.Warn(err)
	}

	filter := bson.M{"_id": objectID}
	result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.D{{"done", true}}}})
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, result)
}
