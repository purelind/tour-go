package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"unsafe"
)

type Post struct {
	Id       int       "json1:\"id\""
	Content  string    //`json:"content"`
	Author   Author    //`json:"author"`
	Comments []Comment //`json:"comments"`
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
			Name: "kevin",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	a := 'A'
	b := `sdfsd` + string(a) + `fs`
	fmt.Println(unsafe.Sizeof(a), a, b)
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
