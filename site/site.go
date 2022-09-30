package site

import (
	"path/filepath"

	"github.com/tobier/rsf-rally-results/model"
)

type site struct {
	Title      string
	Navigation navbar
	Content    interface{}
}

// Render the entire website to the given output directory.
func Render(m *model.Model, outdir string) error {
	// Render the front page
	site := site{
		Title: m.Name,
		Navigation: navbar{
			Logo: "/assets/images/logo-fallback.jpg",
			Items: []navlink{
				{
					Text:   "Home",
					URL:    "/",
					Active: false,
				},
				{
					Text:   "Results",
					URL:    "#",
					Active: false,
					Children: []navlink{
						{
							Text: "Series A",
							URL:  "/results/series-a",
						},
						{
							Text: "Series B",
							URL:  "/results/series-b",
						},
					},
				},
				{
					Text:   "Standings",
					URL:    "/standings",
					Active: false,
				},
			},
		},
	}

	site.Navigation.activate(navHome)

	homepageFile := filepath.Join(outdir, "index.html")
	if err := render(homepageFile, &site); err != nil {
		return err
	}

	// TODO: render the results page

	// TODO: render the standings page
	site.Navigation.activate(navStandings)

	standingsFile := filepath.Join(outdir, "standings", "index.html")
	if err := render(standingsFile, &site); err != nil {
		return err
	}

	return nil
}
