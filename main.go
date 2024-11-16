package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/renderer"
	"fitness-api/cmd/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Add a handler function that returns all the data from the table

func main() {
	e := echo.New()
	e.Renderer = renderer.NewRenderer("./*.html", true)
	type M map[string]interface{}
	e.GET("/index", func(c echo.Context) error { // did this to render a static website
		data := M{"message": "Hello World !"}
		return c.Render(http.StatusOK, "index.html", data)
	})
	storage.InitDB()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "print_users.html", M{})
	})
	e.POST("/users", handlers.CreateUser)
	e.POST("/measurements", handlers.CreateMeasurement)
	e.PUT("/users/:id", handlers.HandleUpdateUser)
	e.GET("/get_users/:id", handlers.Return_selected_rows)
	e.Use(handlers.LogRequest)
	//----------------
	e.Logger.Fatal(e.Start(":8080"))
}

// These are the following requests you can make

// curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d "{\"id\": 1, \"name\": \"John Doe\", \"email\": \"john@mail.com\", \"password\": \"really?\", \"created_at\": \"2022-01-01T09:00:35Z\", \"updated_at\": \"2022-01-01T09:00:35Z\"}"
// curl -X POST http://localhost:8080/measurements -H "Content-Type: application/json" -d "{\"user_id\": 1, \"weight\" : 80, \"height\" : 180, \"body_fat\" : 20}"
// curl -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d "{\"name\" : \"Jane Wanjiru\",\"email\" : \"jane@mail.com\", \"password\" : \"Nah Bruh\"}"
// curl -X PUT http://localhost:8080/get_users/1

// inspiration : https://www.kelche.co/blog/go/golang-echo-tutorial/
