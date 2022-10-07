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
							Text: "Bilsport Rallycup 2022",
							URL:  "/results/bilsport2022",
						},
						{
							Text: "Mini-WRC 2022",
							URL:  "/results/miniwrc2022",
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
	if err := render(homepageFile, "templates/home.tpl.html", &site); err != nil {
		return err
	}

	// TODO: render the results page
	allSeries := []series{
		{
			Name: "Bilsport Rallycup 2022",
			URL:  "bilsport2022",
			Events: []event{
				{
					Name: "Rally Anderslöv",
					ID:   47024,
					Results: []result{
						{
							Driver:    "Tommi Mäkinen",
							Total:     "24:50.083",
							DiffPrev:  "--:--.---",
							DiffFirst: "--:--.---",
						},
						{
							Driver:    "Ari Vatanen",
							Total:     "24:59.683",
							DiffPrev:  "00:09.600",
							DiffFirst: "00:09.600",
						},
						{
							Driver:    "Hanu Mikkola",
							Total:     "26:24.486",
							DiffPrev:  "01:24.803",
							DiffFirst: "01:34.402",
						},
					},
				},
				{
					Name: "Zabra-rallyt",
					ID:   47699,
				},
			},
		},
	}

	for _, series := range allSeries {
		site.Navigation.activate(navResults)
		site.Content = series
		path := filepath.Join(outdir, "results", series.URL, "index.html")
		if err := render(path, "templates/results.tpl.html", &site); err != nil {
			return err
		}
	}

	// TODO: render the standings page
	//site.Navigation.activate(navStandings)
	//site.Content = []standings{
	//	{
	//		Name: "Bilsport Rallycup 2022",
	//	},
	//	{
	//		Name: "Mini-WRC 2022",
	//	},
	//}

	//standingsFile := filepath.Join(outdir, "standings", "index.html")
	//if err := render(standingsFile, "templates/standings.tpl.html", &site); err != nil {
	//	return err
	//}

	return nil
}
