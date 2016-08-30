package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

func newGetUserHandler(mongoSession *mgo.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Grab id
		id := p.ByName("id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			return
		}

		// Grab id
		oid := bson.ObjectIdHex(id)

		// Stub user
		u := &user{}

		// Fetch user
		if err := mongoSession.DB("go_rest_tutorial").C("users").FindId(oid).One(u); err != nil {
			w.WriteHeader(404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&u)
	}
}

// CreateUser creates a new user resource
func newCreateUserHandler(mongoSession *mgo.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		u := &user{}

		// Populate the user data
		json.NewDecoder(r.Body).Decode(u)

		// Add an Id
		u.ID = bson.NewObjectId()

		// Write the user to mongo
		mongoSession.DB("go_rest_tutorial").C("users").Insert(u)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
		w.WriteHeader(201)
	}
}

// RemoveUser removes an existing user resource
func newRemoveUserHandler(mongoSession *mgo.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Grab id
		id := p.ByName("id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			return
		}

		// Grab id
		oid := bson.ObjectIdHex(id)

		// Remove user
		if err := mongoSession.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
			w.WriteHeader(404)
			return
		}

		// Write status
		w.WriteHeader(204)
	}
}
