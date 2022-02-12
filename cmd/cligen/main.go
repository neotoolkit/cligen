package main

import (
	"flag"
	"log"
	"os"

	"github.com/neotoolkit/cligen"
)

const _generateFileName = "run.go"

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

	f, err := os.OpenFile(_generateFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	executeTemplateErr := t.Execute(f, api)
	fileCloseErr := f.Close()

	if fileCloseErr != nil {
		log.Fatalln(err)
	}

	if executeTemplateErr != nil {
		deleteFileErr := os.Remove(_generateFileName)

		if deleteFileErr != nil {
			log.Println(deleteFileErr)
		}

		log.Fatalln(executeTemplateErr)
	}
}
