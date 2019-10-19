package main

import (
	"os"
	"text/template"
)

func main() {
	tpt := `{{- . -}}`
	tmpl, err := template.New("test").Parse(tpt)
	if err != nil {
		panic(err)
	}

	content := " jack "
	if err = tmpl.Execute(os.Stdout, content); err != nil {
		panic(err)
	}
}
