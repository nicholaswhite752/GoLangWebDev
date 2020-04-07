package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oid := bson.ObjectIdHex(id)

	// composite literal
	u, ok :=  uc.session[oid.String()]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User At ID Does Not Exist", u.Id, "\n")
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = bson.NewObjectId()

	if _, ok := uc.session[u.Id.String()]; ok { // will be false if person is not in the map
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User At ID Already Exist", u.Id, "\n")
		return
	}

	uc.session[u.Id.String()] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if _, ok := uc.session[oid.String()]; !ok { // will be false if person is not in the map
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User At ID Does not Exist", oid, "\n")
		return
	}

	delete(uc.session, oid.String())
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
