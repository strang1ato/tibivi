package tibivi

import (
	"github.com/oltarzewskik/gocui"
)

// Views is struct of all views in tibivi
type Views struct {
	days, extra map[string]*gocui.View
}

// setCurrentViewOnTop executes SetCurrentView and SetViewOnTop
func (tbv *Tibivi) setCurrentViewOnTop(name string) (*gocui.View, error) {
	if _, err := tbv.g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return tbv.g.SetViewOnTop(name)
}
