package main

import (
	"flag"
	"github.com/neotoolkit/cligen"
	"log"
	"os"
)

const GENERATE_FILE_NAME = "run.go"

func main() {
	var cli string
	flag.StringVar(&cli, "cli", "acmd", "")

	flag.Parse()

	api, err := cligen.Parse("./.cligen.yml")
	if err != nil {
		log.Fatalln(err)
	}

	if err := cligen.Validate(api.Commands); err != nil {
		log.Fatalln(err)
	}

	t := cligen.Template(cli)

	f, err := os.OpenFile(GENERATE_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	executeTemplateErr := t.Execute(f, api)
	fileCloseErr := f.Close()

	if fileCloseErr != nil {
		log.Fatalln(err)
	}

	if executeTemplateErr != nil {
		deleteFileErr := os.Remove(GENERATE_FILE_NAME)

		if deleteFileErr != nil {
			log.Println(deleteFileErr)
		}

		log.Fatalln(executeTemplateErr)
	}
}
