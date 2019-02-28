package main

import( 
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`	
}

func main() {
	bytes, err := json.Marshal(Person{
		FirstName : "John",
		LastName  : "Dow",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}