package main

import (
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
	movies := []Movie{
		{
			Title: "Casablanca",
			Year: 1942,
			Color: false,
			Actors: []string {
				"Humphrey Bogart",
				"Ingrid Bergman", 
			}, 
		},
		{
			Title: "Cool Hand Luke",
			Year: 1967,
			Color: true,
			Actors: []string {
				"Paul Newman", 
			}, 
		},
		{
			Title: "Bullit",
			Year: 1968,
			Color: true,
			Actors: []string {
				"Steve McQueen",
				"Jacqueline Bisset", 
			}, 
		},
	}

	// Conversão não formatada
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON marshaling failed: %s", err)
	}

	/* Conversão "embelazada"
	data, err = json.MarshalIndent(movies, ""," ")
	if err != nil {
		log.Fatal("JSON marshaling failed: %s", err)
	}
	*/

	// Gravar em um arquivo

	err = ioutil.WriteFile("movies.json", data, 0644)
	if err != nil {
		log.Fatal("JSON write to file failed: %s", err)
	}
}