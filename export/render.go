package export

import (
	"html/template"
	"os"

	"github.com/caarlos0/log"
)

// TODO: this is a placeholder for now
func render(filename string) (err error) {
	log.WithField("file", filename).Debug("Rendering template")

	outfile, err := os.Create(filename)
	if err != nil {
		return
	}

	t, err := template.ParseFiles("templates/results.tpl.html", "templates/base.tpl.html", "templates/nav.tpl.html")
	if err != nil {
		return
	}

	return t.ExecuteTemplate(outfile, "base", nil) // TODO: no data (yet) ...
}
