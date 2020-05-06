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

// setCurrentViewOnTop executes SetCurrentView and SetViewOnTop
func (tbv *Tibivi) setCurrentViewOnTop(name string) (*gocui.View, error) {
	if _, err := tbv.g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return tbv.g.SetViewOnTop(name)
}
