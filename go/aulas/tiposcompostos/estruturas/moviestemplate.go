// Fonte: https://gowebexamples.com/templates/
package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"html/template"
	"os"
)

type MoviePage struct {
	PageTitle string
	Movies []Movie
}

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

	data := MoviePage {
		PageTitle: "Template Example",
		Movies: titles,
	}

	// Processa o arquivo "movies.tmpl" e prepara um objeto a ser preenchido pelos campos da estrutura.
	t := template.Must(template.ParseFiles("movies.tmpl"))	

	// Joga na saída padrão o html preenchido
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}