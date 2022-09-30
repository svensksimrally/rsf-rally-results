package site

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/log"
)

func render(outfileName string, data interface{}) (err error) {
	log.WithField("file", outfileName).Debug("Rendering template")

	parent := filepath.Dir(outfileName)
	err = os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		return
	}

	outfile, err := os.Create(outfileName)
	if err != nil {
		return
	}

	functions := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	t, err := template.New("base").Funcs(functions).ParseFiles("templates/nav.tpl.html", "templates/base.tpl.html")
	if err != nil {
		return
	}

	return t.ExecuteTemplate(outfile, "base", data)
}
