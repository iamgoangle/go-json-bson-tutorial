package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Address struct {
	Street string
	City   string
	State  string
}

type Student struct {
	FirstName string   `bson:"first_name,omitempty"`
	LastName  string   `bson:"last_name,omitempty"`
	IsActive  bool     `bson:"is_active,omitempty"`
	Address   *Address `bson:"address"`
	Age       int
}

func main() {
	b, err := bson.Marshal(&Student{Address: nil})
	if err != nil {
		panic(err)
	}

	var result Student
	err = bson.Unmarshal(b, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", string(b))
	fmt.Printf("%+v", result)
}
