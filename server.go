package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"os"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	mongoSession := getSession()

	r.GET("/player/:accountId", getPlayerInfoHandler)
	// Get a user resource
	r.GET("/user/:id", newGetUserHandler(mongoSession))

	// Create a new user
	r.POST("/user", newCreateUserHandler(mongoSession))

	// Remove an existing user
	r.DELETE("/user/:id", newRemoveUserHandler(mongoSession))

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
