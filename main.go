package main

import (
	"encoding/json"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	base64Encode()
	p1 := person{
		First: "Jenny",
	}

	p2 := person {
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

func base64Encode() {
	fmt.Println("encode string to 64 demo")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}