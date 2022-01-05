package main

import (
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
)

type Movie struct {
	Title string 
	Year int        `json:"released"`
	Color bool      `json:"color,omitempty"`
	Actors []string
}

func main() {
	var titles []Movie

	file, err := ioutil.ReadFile("movies.json")
	if err != nil {
		log.Fatal("JSON read from file failed: %s", err)
	}

	if err = json.Unmarshal([]byte(file), &titles); err != nil {
		log.Fatal("JSON unmarshaling failed: %s", err)
	}

	for _, movie := range titles {
		fmt.Println(movie.Title)
	}
}