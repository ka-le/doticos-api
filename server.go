package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/ka-le/doticos-api/controllers"
	"gopkg.in/mgo.v2"
	"os"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPlayerController()

	r.GET("/player/:accountId", pc.GetPlayerInfo)
	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	// Create a new user
	r.POST("/user", uc.CreateUser)

	// Remove an existing user
	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}