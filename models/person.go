package models

import (
	"fmt"
	"github.com/daniel/rest-api-persons/db"
	"github.com/satori/go.uuid"
	"time"
)

type Person struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastName, omitempty"`
}

func CreatePerson(person Person) {
	sqlAddPerson := `INSERT OR REPLACE INTO person (
		Id,
		Name,
		LastName,
		created_at
	) values(?,?,?, CURRENT_TIMESTAMP)`

	statement, err := db.DBCon.Prepare(sqlAddPerson)
	if err != nil {
		panic(err)
	}
	defer statement.Close()
	_, err = statement.Exec(person.Id, person.Name, person.LastName)
}

func FindPerson(id string) Person {
	rows, err := db.DBCon.Query("SELECT * FROM person WHERE ID = ?", id)

	if err != nil {
		panic(err)
	}
	var person Person

	var created_at time.Time
	for rows.Next() {
		err = rows.Scan(&person.Id, &person.Name, &person.LastName, &created_at)

		if err != nil {
			panic(err)
		}
	}

	return person
}

func UpdatePerson(person Person) {
	statement, err := db.DBCon.Prepare("UPDATE person set Name=?, LastName=? WHERE ID = ?")

	if err != nil {
		panic(err)
	}

	response, err := statement.Exec(person.Name, person.LastName, person.Id)

	if err != nil {
		panic(err)
	}

	affect, err := response.RowsAffected()

	fmt.Println(affect)
	if err != nil {
		panic(err)
	}

}

func FindAll() []Person {
	rows, err := db.DBCon.Query("SELECT * FROM person")

	if err != nil {
		panic(err)
	}

	var person Person
	var people []Person
	var created_at time.Time

	for rows.Next() {
		err = rows.Scan(&person.Id, &person.Name, &person.LastName, &created_at)

		if err != nil {
			panic(err)
		}
		people = append(people, person)
	}

	return people
}
