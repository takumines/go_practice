package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "takumi kurogi",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "Have a great Day",
				Author:  "admin",
			},
			Comment{
				Id:      4,
				Content: "How are you Today",
				Author:  "jhon",
			},
		},
	}

	jsonFile, err := os.Create("postA.json")
	if err != nil {
		fmt.Println("Error creating Json file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding Json to file", err)
		return
	}
}
