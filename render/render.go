package render

import (
	"html/template"
	"os"

	"github.com/caarlos0/log"
)

func render(outfileName string, data interface{}) (err error) {
	log.WithField("file", outfileName).Debug("Rendering template")

	outfile, err := os.Create(outfileName)
	if err != nil {
		return
	}

	t, err := template.ParseFiles("templates/nav.tpl.html", "templates/base.tpl.html")
	if err != nil {
		return
	}

	return t.ExecuteTemplate(outfile, "base", data)
}
