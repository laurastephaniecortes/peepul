package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"github.com/rs/cors"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Person struct {
	Name           string
	BodyType       string
	Age            string
	Sex            string
	Attractiveness string
}

type Hookup struct {
	Top    Person
	Bottom Person
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	db, err := scribble.New("peeps.db", nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	records, err := db.ReadAll("people")
	if err != nil {
		fmt.Println("Error", err)
	}
	json.NewEncoder(w).Encode(records)

	// for _, record := range records {
	// 	json.NewEncoder(w).Encode(w)
	// 	w.Write([]byte(string(record)))

	// }

	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func addPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	person := Person{}
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))

	json.Unmarshal(b, &person) 


	db, err := scribble.New("peeps.db", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	if err := db.Write("people", person.Name, person); err != nil {
		fmt.Println("Error", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", getPeople)
	r.HandleFunc("/addperson", addPerson)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET", "POST"}, // Allowing only get, just an example
	})
  
	if err := http.ListenAndServe(":"+port,  c.Handler(r)); err != nil {
		panic(err)
	}
}
