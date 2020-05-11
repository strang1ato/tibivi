package tibivi

import (
	"github.com/oltarzewskik/tibivi-gocui"
)

// Views is struct of tibivi views and its etceteras
type Views struct {
	bar, days, menu  map[string]*gocui.View
	addBlockFields   []string
	currentViewOnTop string
}
