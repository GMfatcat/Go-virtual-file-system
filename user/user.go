package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
	s "virtual-file-system/setting"
)

type JSONData struct {
	Data map[string]s.UserInfo `json:"data"`
}

/* User Command Functions */
/*
User name length 3~10,A-Za-z0-9
Folder name length 3~10,A-Za-z0-9
File name length 3~10,A-Za-z0-9
*/

/* Command: register [username] */
func (jsonObj *JSONData) RegisterName(username, userInfoPath string) error {

	// Check if username contains invalid characters using regex
	if regexCheckErr := RegexCheck(username); regexCheckErr != nil {
		return regexCheckErr
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

/* Command: create-folder [username] [foldername] [description]? */
func (jsonObj *JSONData) CreateFolder(inputParts []string, userInfoPath string) error {

	var username, foldername string
	var description bool = false

	commandLength := len(inputParts)
	if !(commandLength == 3 || commandLength == 4) {
		return fmt.Errorf("create-folder requires 3 or 4 arguments.\n")
	}

	username = inputParts[1]
	foldername = inputParts[2]
	// Check folder name format
	if regexCheckErr := RegexCheck(foldername); regexCheckErr != nil {
		return regexCheckErr
	}

	// Check if description contains invalid characters using regex
	if commandLength == 4 {
		if regexCheckErr := RegexCheck(inputParts[3]); regexCheckErr != nil {
			return regexCheckErr
		}
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

/* Command: rename-folder [username] [foldername] [new-folder-name] */
func (jsonObj *JSONData) RenameFolder(inputParts []string, userInfoPath string) error {
	var username, foldername, newFoldername string
	// Input check
	commandLength := len(inputParts)
	if commandLength != 4 {
		return fmt.Errorf("rename-folder requires 4 arguments.\n")
	}
	username = inputParts[1]
	foldername = inputParts[2]
	newFoldername = inputParts[3]
	// username check
	if usernameErr := jsonObj.UsernameCheck(username); usernameErr != nil {
		return usernameErr
	}
	// foldername check, need to found exist folder
	if foldernameErr := jsonObj.FoldernameCheck(username, foldername); foldernameErr == nil {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}
	// newFolder name check
	if regexCheckErr := RegexCheck(newFoldername); regexCheckErr != nil {
		return regexCheckErr
	}
	// Rename Folder
	if FolderErr := jsonObj.OsRenameFolder(inputParts, userInfoPath); FolderErr != nil {
		return FolderErr
	}
	return nil
}

/* Command: delete-folder [username] [foldername] */
func (jsonObj *JSONData) DeleteFolder(inputParts []string, userInfoPath string) error {
	var username, foldername string
	// Input check
	commandLength := len(inputParts)
	if commandLength != 3 {
		return fmt.Errorf("delete-folder requires 3 arguments.\n")
	}
	username = inputParts[1]
	foldername = inputParts[2]
	// username check
	if usernameErr := jsonObj.UsernameCheck(username); usernameErr != nil {
		return usernameErr
	}
	// foldername check, need to found exist folder
	if foldernameErr := jsonObj.FoldernameCheck(username, foldername); foldernameErr == nil {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}
	// Delete Folder
	if deleteFolderErr := jsonObj.OsDeleteFolder(inputParts, userInfoPath); deleteFolderErr != nil {
		return deleteFolderErr
	}

	return nil
}

/* Command: list-folders [username] [--sort-name | --sort-created] [asc|desc] */
func (jsonObj *JSONData) ListFolders(inputParts []string) error {
	var username, sortType, sortRule string
	// Input check
	commandLength := len(inputParts)
	if commandLength != 4 {
		return fmt.Errorf("rename-folder requires 4 arguments.\n")
	}

	username = inputParts[1]
	sortType = inputParts[2]
	sortRule = inputParts[3]

	if !(sortType == "--sort-name" || sortType == "--sort-created") {
		return fmt.Errorf("Invalid Sort Type, use --sort-name or --sort-created.\n")
	}

	if !(sortRule == "asc" || sortRule == "desc") {
		return fmt.Errorf("Invalid Sort Rule, use asc or desc.\n")
	}
	// User check + User folder check
	if usernameErr := jsonObj.UsernameCheck(username); usernameErr != nil {
		return usernameErr
	}
	userFolderNum := jsonObj.FolderNum(username)
	if userFolderNum == 0 {
		return fmt.Errorf("Warning: The %s doesn't have any folders.\n", username)
	}
	// Sort Folder
	jsonObj.SortFolder(inputParts)

	return nil
}

/* Command: create-file [username] [foldername] [filename] [description]? */
func (jsonObj *JSONData) CreateFile(inputParts []string, userInfoPath string) error {
	var username, foldername, filename string
	var description bool = false

	commandLength := len(inputParts)
	if !(commandLength == 4 || commandLength == 5) {
		return fmt.Errorf("create-file requires 4 or 5 arguments.\n")
	}

	username = inputParts[1]
	foldername = inputParts[2]
	filename = inputParts[3]
	// Check file name format
	if regexCheckErr := RegexCheck(filename); regexCheckErr != nil {
		return regexCheckErr
	}
	// Check if description contains invalid characters using regex
	if commandLength == 5 {
		if regexCheckErr := RegexCheck(inputParts[4]); regexCheckErr != nil {
			return regexCheckErr
		}
		description = true
	}
	// username check
	if usernameErr := jsonObj.UsernameCheck(username); usernameErr != nil {
		return usernameErr
	}
	// foldername check: need folder to be exist
	if foldernameErr := jsonObj.FoldernameCheck(username, foldername); foldernameErr == nil {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}
	// Create File
	if FileErr := jsonObj.OsCreateFile(inputParts, description, userInfoPath); FileErr != nil {
		return FileErr
	}

	return nil
}

/* Command: delete-file [username] [foldername] [filename] */
func (jsonObj *JSONData) DeleteFile(inputParts []string, userInfoPath string) error {
	fmt.Printf("Not implemented yet.\n")
	return nil
}

/* Command: list-files [username] [foldername] [--sort-name | --sort-created] [asc|desc] */
func (jsonObj *JSONData) ListFiles(inputParts []string) error {
	fmt.Printf("Not implemented yet.\n")
	return nil
}

/* End of User Command Functions */

/* Username & Foldername & Filename Check & Edit */
func (jsonObj *JSONData) FolderNum(username string) int {
	userInfo, _ := jsonObj.Data[username]
	numFolders := len(userInfo.Folders)
	return numFolders
}

func (jsonObj *JSONData) SortFolder(inputParts []string) {
	var username, sortType, sortRule string
	username = inputParts[1]
	sortType = inputParts[2]
	sortRule = inputParts[3]

	// Select Sort Conditions
	userInfo := jsonObj.Data[username]

	switch sortType {
	case "--sort-name":
		if sortRule == "asc" {
			sort.Sort(byName(userInfo.Folders))
		} else {
			sort.Sort(sort.Reverse(byName(userInfo.Folders)))
		}
	case "--sort-created":
		if sortRule == "asc" {
			sort.Sort(byTime(userInfo.Folders))
		} else {
			sort.Sort(sort.Reverse(byTime(userInfo.Folders)))
		}
	}

	// Show Sort Result
	fmt.Fprintf(os.Stdout, "Sort Type: %s, Sort Rule: %s\n", sortType, sortRule)
	for _, folder := range userInfo.Folders {
		fmt.Fprintf(os.Stdout, "Name:%s Time:%s\n",
			folder.Name, folder.CreatedAt.Format(time.RFC822))
	}
}

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

// return the index of folder, otherwise return -1
func (jsonObj *JSONData) findFolderIndex(username, foldername string) int {
	for i, folder := range jsonObj.Data[username].Folders {
		if folder.Name == foldername {
			return i
		}
	}
	return -1
}

func (jsonObj *JSONData) OsRenameFolder(inputParts []string, userInfoPath string) error {
	var username, foldername, newFoldername string
	username = inputParts[1]
	foldername = inputParts[2]
	newFoldername = inputParts[3]

	//Update Json
	folderIndex := jsonObj.findFolderIndex(username, foldername)
	if folderIndex == -1 {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}

	jsonObj.Data[username].Folders[folderIndex].Name = newFoldername

	if err := jsonObj.saveUserInfoToFile(userInfoPath); err != nil {
		return fmt.Errorf("Error saving JSON data: %v", err)
	}

	//OS rename folder
	var rootPath string = "./app"

	oldFolderPath := filepath.Join(rootPath, username, foldername)
	newFolderPath := filepath.Join(rootPath, username, newFoldername)

	if err := os.Rename(oldFolderPath, newFolderPath); err != nil {
		return fmt.Errorf("Error renaming folder: %v", err)
	}

	fmt.Fprintf(os.Stdout, "Rename %s to %s successfully.\n", foldername, newFoldername)
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

func (jsonObj *JSONData) OsDeleteFolder(inputParts []string, userInfoPath string) error {
	var username, foldername string
	username = inputParts[1]
	foldername = inputParts[2]
	//Update Json
	folderIndex := jsonObj.findFolderIndex(username, foldername)
	if folderIndex == -1 {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}

	userInfo, _ := jsonObj.Data[username]
	userInfo.Folders = append(userInfo.Folders[:folderIndex], userInfo.Folders[folderIndex+1:]...)
	jsonObj.Data[username] = userInfo
	// Save Json
	if err := jsonObj.saveUserInfoToFile(userInfoPath); err != nil {
		return fmt.Errorf("Error saving JSON data: %v", err)
	}
	//Os delete folder
	var rootPath string = "./app"
	folderPath := filepath.Join(rootPath, username, foldername)
	if err := os.RemoveAll(folderPath); err != nil {
		return fmt.Errorf("Error deleting folder: %v", err)
	}

	fmt.Fprintf(os.Stdout, "Delete %s successfully.\n", foldername)
	return nil
}

func (jsonObj *JSONData) OsCreateFile(inputParts []string, description bool, userInfoPath string) error {
	var username, foldername, filename string
	var fileDescription string = ""

	username = inputParts[1]
	foldername = inputParts[2]
	filename = inputParts[3]

	if description {
		fileDescription = inputParts[4]
	}

	folderIndex := jsonObj.findFolderIndex(username, foldername)
	if folderIndex == -1 {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}
	// Record to Userinfo
	newFile := s.File{
		Name:        filename,
		Description: fileDescription,
		CreatedAt:   time.Now(),
	}
	userInfo, _ := jsonObj.Data[username]
	userInfo.Folders[folderIndex].Files = append(userInfo.Folders[folderIndex].Files, newFile)
	jsonObj.Data[username] = userInfo
	// Save Json
	if err := jsonObj.saveUserInfoToFile(userInfoPath); err != nil {
		return fmt.Errorf("Error saving JSON data: %v", err)
	}
	// Os Create File
	var rootPath string = "./app"
	filePath := filepath.Join(rootPath, username, foldername, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(os.Stdout, "Create %s in %s/%s successfully.\n", filename, username, foldername)

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

func RegexCheck(input string) error {
	inputRegex, err := regexp.Compile(`^[A-Za-z0-9]{3,10}$`)
	if err != nil {
		return fmt.Errorf("Regex Compile Error: %v\n", err)
	}

	if !inputRegex.MatchString(input) {
		return fmt.Errorf("The %s contain invalid chars.\n", input)
	}
	return nil
}

/* End of Username & Foldername & Filename Check & Edit */

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

/* Sort Interface */
type byName []s.Folder

func (a byName) Len() int      { return len(a) }
func (a byName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// define in asc
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type byTime []s.Folder

func (a byTime) Len() int      { return len(a) }
func (a byTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// define in asc
func (a byTime) Less(i, j int) bool { return a[i].CreatedAt.Before(a[j].CreatedAt) }
