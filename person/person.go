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
/*type Person struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastName, omitempty"`
}*/

var person DBPerson.Person

func CreatePerson(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		panic(err.Error())
	}

	person.Id = uuid.NewV4()
	fmt.Println(person.Name)
	personJson, err := json.Marshal(person)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(personJson)
}
