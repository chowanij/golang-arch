package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("PRINT JSON", string(bs))

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("back into a Go data structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.HandleFunc("/encode/persons", encodePerson)
	http.HandleFunc("/decode/persons", decodePersons)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}

}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decode bad data", err)
	}
	log.Println("Person", p1)
}

func encodePerson(w http.ResponseWriter, r *http.Request) {
	persons := []person{
		{
			First: "Jenny",
		},
		{
			First: "Jennys",
		},
		{
			First: "Johnes",
		},
	}

	err := json.NewEncoder(w).Encode(persons)
	if err != nil {
		log.Println("Encoded bad data", err)
	}

}

func decodePersons(w http.ResponseWriter, r *http.Request) {
	persons := []person{}
	err := json.NewDecoder(r.Body).Decode(&persons)
	if err != nil {
		log.Println("Decode bad data", err)
	}
	log.Println("Person", persons)
}
