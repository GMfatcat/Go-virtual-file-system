package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	s "virtual-file-system/setting"
)

type JSONData struct {
	Data map[string]s.UserInfo `json:"data"`
}

/* User Functions */
/*User name length 3~10,A-Za-z0-9*/
func (jsonObj *JSONData) RegisterName(userName, userInfoPath string) error {

	// Check if userName contains invalid characters using regex
	userNameRegex, err := regexp.Compile(`^[A-Za-z0-9]{3,10}$`)
	if err != nil {
		return fmt.Errorf("Regex Compile Error: %v\n", err)
	}

	if !userNameRegex.MatchString(userName) {
		return fmt.Errorf("The %s contain invalid chars.\n", userName)
	}

	// Check input name exist in userinfo json
	if _, ok := jsonObj.Data[userName]; !ok {
		jsonObj.Data[userName] = s.UserInfo{Folders: []s.Folder{}}

		// Save JSON
		if err := jsonObj.saveUserInfoToFile(userInfoPath); err != nil {
			return fmt.Errorf("Error saving JSON data: %v", err)
		}
	} else {
		return fmt.Errorf("The %s has already existed.\n", userName)
	}

	return nil
}

func (jsonObj *JSONData) saveUserInfoToFile(userInfoPath string) error {
	jsonData, err := json.Marshal(jsonObj)
	if err != nil {
		return err
	}

	// Write File
	if err = ioutil.WriteFile(userInfoPath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

/* UserInfo Functions*/

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
