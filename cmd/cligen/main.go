package main

import (
	"log"
	"os"
	"text/template"

	"github.com/go-dummy/cligen"
)

func main() {
	t, err := template.ParseFiles("./tmpl/acmd.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile("run.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	api, err := cligen.Parse("./cligen.yml")
	if err != nil {
		log.Fatalln(err)
	}

	if err := t.Execute(f, api); err != nil {
		log.Fatalln(err)
	}
}
