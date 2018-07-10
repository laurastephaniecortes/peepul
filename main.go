package main
import (
  "net/http"
  "encoding/json"
  "fmt"

  "github.com/nanobox-io/golang-scribble"

)

type Person struct {
    Name string
    BodyType string
    Age int64
    Sex bool
    Attractiveness float64 `validate:"max=1"`
}

type Hookup struct {
  Top Person
  Bottom Person
}

func getPeople(w http.ResponseWriter, r *http.Request) {
  db, err := scribble.New("peeps.db", nil)
  if err != nil {
    fmt.Println("Error", err)
  }
  records, err := db.ReadAll("people")
  if err != nil {
    fmt.Println("Error", err)
  }
  for _, record := range records {


  w.Write([]byte(string(record)))
}
}

func addPerson(w http.ResponseWriter, r *http.Request) {
  person := Person{}
  decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&person)
    if err != nil {
        panic(err)
    }

  db, err := scribble.New("peeps.db", nil)
  if err != nil {
    fmt.Println("Error", err)
  }

if err := db.Write("people", person.Name, person); err != nil {
  fmt.Println("Error", err)
}


}


func main() {
  http.HandleFunc("/", getPeople)
  http.HandleFunc("/addperson", addPerson)
  port := os.Getenv("PORT")
if port == "" {
  return "", fmt.Errorf("$PORT not set")
}
return ":" + port, nil
}
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }
}
