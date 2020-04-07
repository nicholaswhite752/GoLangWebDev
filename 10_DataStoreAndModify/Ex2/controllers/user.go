package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/08_hands-on/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"io/ioutil"
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

	// Retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	temp, _ := uuid.NewV4()
	u.Id = temp.String()

	// store the user
	uc.session[u.Id] = u

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

	//f, err := os.Open("data")
	//if err != nil {
	//	fmt.Println(err)
	//	f, _ = os.Create("data")
	//	return
	//}
	//defer f.Close()

	dataWrite, err := json.Marshal(uc.session)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("data", []byte(dataWrite), 0644)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")

	//f, err := os.Open("data")
	//if err != nil {
	//	fmt.Println(err)
	//	f, _ = os.Create("data")
	//	return
	//}
	//defer f.Close()

	dataWrite, err := json.Marshal(uc.session)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("data", []byte(dataWrite), 0644)

}
