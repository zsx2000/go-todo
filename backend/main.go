package main

import (
	controller "go-todo/backend/controller"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:8080"}, //frontent address
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// e.File("/", "../frontend/dist/index.html")
	// e.Static("/css", "../frontend/dist/css/")
	// e.Static("/js", "../frontend/dist/js/")

	e.GET("/api/task/:id", controller.GetTask)
	e.GET("/api/tasks", controller.GetTasks)
	e.POST("/api/task/add", controller.AddTask)
	e.DELETE("/api/task/delete/:id", controller.DeleteTask)
	e.PUT("/api/task/markdone/:id", controller.MarkTaskDone)

	e.Logger.Fatal(e.Start(":1323"))
}
