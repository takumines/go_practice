package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello Miyazaki",
		Author: Author{
			Id:   "2",
			Name: "hoge taro",
		},
	}
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = ioutil.WriteFile("postA.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing to XML:", err)
		return
	}
}
