package tibivi

import (
	"os"
)

// createDotTibivi creates tibivi's dotfolder with its files if they don't exist
func (tbv *Tibivi) createDotTibivi() error {
	// Create dotfolder
	if _, err := os.Stat(tbv.dotTibivi); os.IsNotExist(err) {
		if err := os.Mkdir(tbv.dotTibivi, 0755); err != nil {
			return err
		}
	}

	// Create tibivi's datafiles
	for _, day := range tbv.days {
		if _, err := os.Stat(tbv.dotTibivi + day + ".txt"); os.IsNotExist(err) {
			if _, err := os.Create(tbv.dotTibivi + day + ".txt"); err != nil {
				return err
			}
		}
	}
	return nil
}
