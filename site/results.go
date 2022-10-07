package site

type result struct {
	Driver    string
	Total     string
	DiffPrev  string
	DiffFirst string
}

type event struct {
	Name    string
	ID      int64
	Results []result
}

type series struct {
	Name   string
	URL    string
	Events []event
}
