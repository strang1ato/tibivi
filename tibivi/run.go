package tibivi

import (
	"github.com/oltarzewskik/gocui"
)

// Run runs Tibivi
func Run() error {
	tbv := newTibivi()

	if err := tbv.createDotTibivi(); err != nil {
		return err
	}

	if err := tbv.setSchedule(); err != nil {
		return err
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	tbv.g = g
	defer tbv.g.Close()

	tbv.g.Highlight = true
	tbv.g.SelFgColor = gocui.ColorGreen

	tbv.g.SetManagerFunc(tbv.layout)

	if err := tbv.keybindings(); err != nil {
		return err
	}

	if err := tbv.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}
