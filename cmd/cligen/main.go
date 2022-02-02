package main

import (
	"flag"
	"log"
	"os"

	"github.com/neotoolkit/cligen"
)

func main() {
	var cli string
	flag.StringVar(&cli, "cli", "acmd", "")

	flag.Parse()

	t := cligen.Template(cli)

	f, err := os.OpenFile("run.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	api, err := cligen.Parse("./.cligen.yml")
	if err != nil {
		log.Fatalln(err)
	}

	if err := t.Execute(f, api); err != nil {
		log.Fatalln(err)
	}
}
