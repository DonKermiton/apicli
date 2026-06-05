package collections

import (
	"fmt"
	"os"
)

type Collection struct {
	Name     string
	Location string
}

func getAppConfigDirectory() {

}

func SaveCollection() error {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return err
	}

	fmt.Println(configDir)
	return nil
}
