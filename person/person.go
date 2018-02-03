package person

import (
	"encoding/json"
	"fmt"
	DBPerson "github.com/daniel/rest-api-persons/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"net/http"
)

// Person is the structure definitionfor a Person
// name: .. Name of a user
type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastName, omitempty"`
}

func CreatePerson(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var person Person
	fmt.Println("Create Person")
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(person.Name)
	personJson, err := json.Marshal(person)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(personJson)
}
