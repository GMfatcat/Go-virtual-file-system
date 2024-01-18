package user

import (
	"encoding/json"
	"fmt"
	"os"
	"virtual-file-system/setting"
)

/* User Functions */
func SetName(name string) error {

	// Check if userinfo json exists, create if not exist
	if err := checkUserInfoExists(setting.UserInfoPath); err != nil {
		return err
	}

	// Check input name with userinfo json
	return nil
}

func checkUserInfoExists(jsonPath string) error {

	_, err := os.Stat(jsonPath)
	if os.IsNotExist(err) {

		emptyJSON := make(map[string]interface{})

		file, err := os.Create(jsonPath)
		if err != nil {
			return err
		}

		defer file.Close()

		encoder := json.NewEncoder(file)
		err = encoder.Encode(emptyJSON)
		if err != nil {
			return err
		}

		fmt.Printf("Create Empty Json：%s\n", jsonPath)
	} else if err != nil {
		return err
	}

	return nil
}
