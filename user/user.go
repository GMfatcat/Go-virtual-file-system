package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
	s "virtual-file-system/setting"
)

type JSONData struct {
	Data map[string]s.UserInfo `json:"data"`
}

/* User Functions */
/*User name length 3~10,A-Za-z0-9*/
func (jsonObj *JSONData) RegisterName(username, userInfoPath string) error {

	// Check if username contains invalid characters using regex
	usernameRegex, err := regexp.Compile(`^[A-Za-z0-9]{3,10}$`)
	if err != nil {
		return fmt.Errorf("Regex Compile Error: %v\n", err)
	}

	if !usernameRegex.MatchString(username) {
		return fmt.Errorf("The %s contain invalid chars.\n", username)
	}

	// Check input name exist in userinfo json
	if _, ok := jsonObj.Data[username]; !ok {
		jsonObj.Data[username] = s.UserInfo{Folders: []s.Folder{}}

		// Save JSON
		if err := jsonObj.saveUserInfoToFile(userInfoPath); err != nil {
			return fmt.Errorf("Error saving JSON data: %v", err)
		}

		fmt.Fprintf(os.Stdout, "Add user %s successfully.\n", username)

	} else {
		return fmt.Errorf("The %s has already existed.\n", username)
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

/*
Command: create-folder [username] [foldername] [description]?
*/
func (jsonObj *JSONData) CreateFolder(inputParts []string, userInfoPath string) error {

	var username, foldername string
	var description bool = false

	commandLength := len(inputParts)
	username = inputParts[1]
	foldername = inputParts[2]

	if !(commandLength == 3 || commandLength == 4) {
		return fmt.Errorf("create-folder requires 3 or 4 arguments.\n")
	}

	if commandLength == 4 {
		description = true
	}
	// username check
	if usernameErr := jsonObj.UsernameCheck(username); usernameErr != nil {
		return usernameErr
	}
	// foldername check
	if foldernameErr := jsonObj.FoldernameCheck(username, foldername); foldernameErr != nil {
		return foldernameErr
	}
	// Create Folder
	if FolderErr := jsonObj.OsCreateFolder(inputParts, description, userInfoPath); FolderErr != nil {
		return FolderErr
	}

	return nil
}

/* Username & Foldername & Filename Check */
func (jsonObj *JSONData) UsernameCheck(username string) error {

	if _, ok := jsonObj.Data[username]; !ok {
		return fmt.Errorf("The %s doesn't exist.\n", username)
	}
	return nil
}

func (jsonObj *JSONData) FoldernameCheck(username, foldername string) error {

	for _, folder := range jsonObj.Data[username].Folders {
		if folder.Name == foldername {
			return fmt.Errorf("Folder %s has already exist.\n", foldername)
		}
	}

	return nil
}

func (jsonObj *JSONData) OsCreateFolder(inputParts []string, description bool, userInfoPath string) error {

	var folderDescription string = ""
	username := inputParts[1]
	foldername := inputParts[2]
	if description {
		folderDescription = inputParts[3]
	}
	// Record to Userinfo
	newFolder := s.Folder{
		Name:        foldername,
		Description: folderDescription,
		Files:       []s.File{},
		CreatedAt:   time.Now(),
	}
	userInfo, _ := jsonObj.Data[username]
	userInfo.Folders = append(userInfo.Folders, newFolder)
	jsonObj.Data[username] = userInfo

	// Save JSON
	if err := jsonObj.saveUserInfoToFile(userInfoPath); err != nil {
		return fmt.Errorf("Error saving JSON data: %v", err)
	}

	// Check if /app/username exists, if not, create it
	var rootPath string = "./app"
	userFolderPath := filepath.Join(rootPath, username)
	if err := os.MkdirAll(userFolderPath, os.ModePerm); err != nil {
		return fmt.Errorf("Error creating user folder: %v", err)
	}
	// Os create folder in /app/username folder
	if err := os.Mkdir(filepath.Join(userFolderPath, foldername), os.ModePerm); err != nil {
		return fmt.Errorf("Error creating folder: %v", err)
	}

	fmt.Fprintf(os.Stdout, "Create %s successfully.\n", foldername)
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
