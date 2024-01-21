package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	s "virtual-file-system/setting"
)

/* Os Operation */

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

func (jsonObj *JSONData) OsDeleteFile(inputParts []string, userInfoPath string) error {
	var username, foldername, filename string
	username = inputParts[1]
	foldername = inputParts[2]
	filename = inputParts[3]

	//Update Json
	folderIndex := jsonObj.findFolderIndex(username, foldername)
	fileIndex := jsonObj.findFileIndex(username, foldername, filename)
	userInfo, _ := jsonObj.Data[username]
	userInfo.Folders[folderIndex].Files = append(userInfo.Folders[folderIndex].Files[:fileIndex],
		userInfo.Folders[folderIndex].Files[fileIndex+1:]...)
	jsonObj.Data[username] = userInfo

	// Save Json
	if err := jsonObj.saveUserInfoToFile(userInfoPath); err != nil {
		return fmt.Errorf("Error saving JSON data: %v", err)
	}

	//Os delete file
	var rootPath string = "./app"
	filePath := filepath.Join(rootPath, username, foldername, filename)
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("Error deleting file: %v", err)
	}

	fmt.Fprintf(os.Stdout, "Delete %s in %s/%s successfully.\n",
		filename, username, foldername)

	return nil
}

/* End of Os Operation */
