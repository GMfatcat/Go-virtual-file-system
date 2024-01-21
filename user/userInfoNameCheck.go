package user

import (
	"fmt"
	"regexp"
	s "virtual-file-system/setting"
)

/*
User name length 3~10,A-Za-z0-9
Folder name length 3~10,A-Za-z0-9
File name length 3~10,A-Za-z0-9
*/
func RegexCheck(input string) error {

	var validInputRegexString string = `^[A-Za-z0-9]{3,10}$`

	if s.SupportWhitespace {
		validInputRegexString = `^[A-Za-z0-9 ]{3,10}$`
	}

	inputRegex, err := regexp.Compile(validInputRegexString)
	if err != nil {
		return fmt.Errorf("Regex Compile Error: %v\n", err)
	}

	if !inputRegex.MatchString(input) {
		return fmt.Errorf("The %s contain invalid chars.\n", input)
	}

	return nil
}

/* Username & Foldername & Filename : Check */

func (jsonObj *JSONData) FolderNum(username string) int {
	userInfo, _ := jsonObj.Data[username]
	numFolders := len(userInfo.Folders)
	return numFolders
}

func (jsonObj *JSONData) FolderFileNum(username, foldername string) int {
	folderIndex := jsonObj.findFolderIndex(username, foldername)
	userInfo := jsonObj.Data[username].Folders[folderIndex]
	numFolderFiles := len(userInfo.Files)
	return numFolderFiles
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

func (jsonObj *JSONData) FilenameCheck(username, foldername, filename string) error {

	folderIndex := jsonObj.findFolderIndex(username, foldername)

	for _, file := range jsonObj.Data[username].Folders[folderIndex].Files {
		if file.Name == filename {
			return fmt.Errorf("File %s has already exist.\n", filename)
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

func (jsonObj *JSONData) findFileIndex(username, foldername, filename string) int {

	folderIndex := jsonObj.findFolderIndex(username, foldername)

	for i, file := range jsonObj.Data[username].Folders[folderIndex].Files {
		if file.Name == filename {
			return i
		}
	}

	return -1
}

/* End of Username & Foldername & Filename : Check */
