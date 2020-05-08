package tibivi

// setCurrentViewOnTop executes SetCurrentView and SetViewOnTop
func (tbv *Tibivi) setCurrentViewOnTop(name string) error {
	if _, err := tbv.g.SetCurrentView(name); err != nil {
		return err
	}
	if _, err := tbv.g.SetViewOnTop(name); err != nil {
		return err
	}
	return nil
}
