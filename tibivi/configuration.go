package tibivi

import (
	"os"
)

// createDotTibivi creates tibivi's dotfolder with its files if they don't exist
func (tbv *Tibivi) createDotTibivi() error {
	if _, err := os.Stat(tbv.dotTibivi); os.IsNotExist(err) {
		if err := os.Mkdir(tbv.dotTibivi, 0755); err != nil {
			return err
		}
	}

	tbv.createDatafiles()

	/* UPCOMING */
	// if _, err := os.Stat(tbv.dotTibivi + "config"); os.IsNotExist(err) {
	// 	if _, err := os.Create(tbv.dotTibivi + "config"); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
