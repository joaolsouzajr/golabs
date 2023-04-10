package main

import "encoding/json"

type User struct {
	Name string
	Age  int32
}

func JsonParser(jsonContent string) User {
	var user User
	json.Unmarshal([]byte(jsonContent), &user)
	return user
}

func ListJsonParser(data string) []User {
	var users []User
	json.Unmarshal([]byte(data), &users)
	return users
}
