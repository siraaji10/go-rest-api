package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/siraaji10/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connection
type App struct{}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	handler := transportHTTP.NewHandler()
	handler.SetUpRoutes()

	if err := http.ListenAndServe(":8081", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

// Our main entrypoint for the application
func main() {
	fmt.Println("Go Rest API")

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our Rest API")
		fmt.Println(err)
	}
}
