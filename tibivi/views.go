package tibivi

import (
	"github.com/oltarzewskik/tibivi-gocui"
)

// Views is struct of tibivi views
type Views struct {
	bar,
	days,
	menu,
	block map[string]*gocui.View
	currentViewOnTop string
}
