package site

import (
	"github.com/caarlos0/log"
)

type navitem int64

const (
	navHome navitem = iota
	navResults
	navStandings
)

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

func (n *navbar) activate(ni navitem) {
	log.WithField("navitem", ni).Debug("Activating navigation item")

	for i := range n.Items {
		activate := (ni == navitem(i))
		n.Items[i].Active = activate
	}
}
