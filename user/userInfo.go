package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* UserInfo Functions */

func ReadUserInfo(jsonPath string) (JSONData, error) {
	var jsonObj JSONData

	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return jsonObj, fmt.Errorf("Error reading JSON file:%v", err)
	}

	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		return jsonObj, fmt.Errorf("Error unmarshalling JSON:%v", err)
	}

	return jsonObj, nil
}

func CheckUserInfoExists(jsonPath string) error {

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

/* End of UserInfo Functions */
