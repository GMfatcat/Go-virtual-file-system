package user

import (
	"fmt"
	"os"
	s "virtual-file-system/setting"
)

type JSONData struct {
	Data map[string]s.UserInfo `json:"data"`
}

/* User Command Functions */

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
	if createFileErr := jsonObj.OsCreateFile(inputParts, description, userInfoPath); createFileErr != nil {
		return createFileErr
	}

	return nil
}

/* Command: delete-file [username] [foldername] [filename] */
func (jsonObj *JSONData) DeleteFile(inputParts []string, userInfoPath string) error {
	var username, foldername, filename string
	commandLength := len(inputParts)
	if commandLength != 4 {
		return fmt.Errorf("delete-file requires 4 arguments.\n")
	}

	username = inputParts[1]
	foldername = inputParts[2]
	filename = inputParts[3]

	// username check
	if usernameErr := jsonObj.UsernameCheck(username); usernameErr != nil {
		return usernameErr
	}
	// foldername check: need folder to be exist
	if foldernameErr := jsonObj.FoldernameCheck(username, foldername); foldernameErr == nil {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}
	// filename check: need file to exist
	if filenameErr := jsonObj.FilenameCheck(username, foldername, filename); filenameErr == nil {
		return fmt.Errorf("The %s doesn't exist.\n", filename)
	}
	// Delete File
	if deleteFileErr := jsonObj.OsDeleteFile(inputParts, userInfoPath); deleteFileErr != nil {
		return deleteFileErr
	}

	return nil
}

/* Command: list-files [username] [foldername] [--sort-name | --sort-created] [asc|desc] */
func (jsonObj *JSONData) ListFiles(inputParts []string) error {
	var username, foldername, sortType, sortRule string
	// Input check
	commandLength := len(inputParts)
	if commandLength != 5 {
		return fmt.Errorf("list-files requires 5 arguments.\n")
	}
	username = inputParts[1]
	foldername = inputParts[2]
	sortType = inputParts[3]
	sortRule = inputParts[4]

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
	// foldername check, need to found exist folder
	if foldernameErr := jsonObj.FoldernameCheck(username, foldername); foldernameErr == nil {
		return fmt.Errorf("The %s doesn't exist.\n", foldername)
	}
	// folder file number check
	userFolderFileNum := jsonObj.FolderFileNum(username, foldername)
	if userFolderFileNum == 0 {
		return fmt.Errorf("Warning: The folder is empty.\n")
	}

	// Sort folder file
	jsonObj.SortFile(inputParts)

	return nil
}

/* End of User Command Functions */
