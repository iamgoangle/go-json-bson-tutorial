package main

import (
	"encoding/json"
	"fmt"
)

// User represent user type
type User struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname,omitempty"`
	Favorites []string `json:"favorites,omitempty"`
	IsActive  bool     `json:"IsActive,omitempty"`
}

func main() {
	test := &User{
		Firstname: "",
		Lastname:  "My Lastname",
		Favorites: nil,
	}

	b, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", string(b))
}
