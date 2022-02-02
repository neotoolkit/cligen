package cligen

import (
	"embed"
	"log"
	"text/template"
)

//go:embed tmpl/*.tmpl
var fs embed.FS

func Template(cli string) *template.Template {
	t, err := template.ParseFS(fs, "tmpl/"+cli+".tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	return t
}
