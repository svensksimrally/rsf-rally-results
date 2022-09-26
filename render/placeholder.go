package render

type Placeholder struct {
	Title      string
	Navigation navbar
}

func (p *Placeholder) Render(outfile string) error {

	// todo: read the actual model here
	navigation := []navlink{
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
			Active: true,
		},
	}

	p.Title = "Rally Results"

	p.Navigation = navbar{
		Logo:  "/assets/images/logo-fallback.jpg",
		Items: navigation,
	}

	return render(outfile, p)
}
