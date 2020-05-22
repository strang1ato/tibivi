package config

import (
	"os"

	"github.com/oltarzewskik/tibivi/pkg/common"
)

// createDotTibivi creates tibivi's dotfolder with its files if they don't exist
func CreateDotTibivi() error {
	// Create dotfolder
	if _, err := os.Stat(common.DotTibivi); os.IsNotExist(err) {
		if err := os.Mkdir(common.DotTibivi, 0755); err != nil {
			return err
		}
	}

	// Create tibivi's datafiles
	for _, day := range common.Days {
		if _, err := os.Stat(common.DotTibivi + day + ".txt"); os.IsNotExist(err) {
			if _, err := os.Create(common.DotTibivi + day + ".txt"); err != nil {
				return err
			}
		}
	}
	return nil
}
