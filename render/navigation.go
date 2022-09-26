package render

type navlink struct {
	Text     string
	URL      string
	Active   bool
	Children []navlink
}

type navbar struct {
	Logo  string
	Items []navlink
}
