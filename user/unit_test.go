/* Unit Test for user package */
package user

import (
	"testing"
	"time"
	"virtual-file-system/setting"
	s "virtual-file-system/setting"
)

// userInfoPath : test empty file
// testUserInfoPath : test not-empty file
var (
	userInfoPath     = "test_user_info.json"
	testUserInfoPath = "user_info.json"
)

func TestCheckUserInfoExists(t *testing.T) {
	t.Skip()
	// Check if userinfo json exists, create if not exist
	if err := CheckUserInfoExists(userInfoPath); err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	} else {
		t.Log("Expected no error, and got none")
	}

}

func TestReadUserInfo(t *testing.T) {

	t.Skip()

	_, err := ReadUserInfo(userInfoPath)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	} else {
		t.Log("Expected no error, and got none")
	}
}

func TestRegisterName(t *testing.T) {

	t.Skip()

	jsonObj, err := ReadUserInfo(userInfoPath)

	// Init Data if necessary
	if jsonObj.Data == nil {
		jsonObj.Data = make(map[string]s.UserInfo)
	}

	// Test Valid Name
	err = jsonObj.RegisterName("validName", userInfoPath)
	if err == nil {
		t.Log("Expected no error, and got none")
	} else {
		t.Error("Expected no error, but got one")
	}

	// Test Invalid Name
	err = jsonObj.RegisterName("invalid@Name", userInfoPath)
	if err != nil {
		t.Log("Expected an error, and got one")
	} else {
		t.Error("Expected an error, but got none")
	}

	// Test Existing Name
	err = jsonObj.RegisterName("validName", userInfoPath)
	if err != nil {
		t.Log("Expected an error, and got one")
	} else {
		t.Error("Expected an error, but got none")
	}
}

func TestUsernameCheck(t *testing.T) {

	t.Skip()

	var userNameList = []string{"AAAaa123", "fatcat", "jjjadi12", "vasdga321"}
	var existName int = 0
	var nonExistName int = 0

	jsonObj, err := ReadUserInfo(testUserInfoPath)
	if err != nil {
		t.Error("Expected no error,but got one")
	}
	for _, name := range userNameList {
		if err = jsonObj.UsernameCheck(name); err != nil {
			nonExistName++
		} else {
			existName++
		}
	}

	if nonExistName == 1 && existName == 3 {
		t.Log("Run as Expected")
	} else {
		t.Errorf("Expected 3 exist 1 non exist,got %d exist %d non exist", existName, nonExistName)
	}
}

func TestOsCreateFolder(t *testing.T) {

	t.Skip()

	jsonObj, err := ReadUserInfo(testUserInfoPath)
	if err != nil {
		t.Error("Expected no error,but got one")
	}
	// set input
	inputParts := []string{
		"create-folder",
		"fatcat",
		"folder1",
		"",
	}
	var description bool = false
	// Create Folder
	if FolderErr := jsonObj.OsCreateFolder(inputParts, description, testUserInfoPath); FolderErr != nil {
		t.Errorf("Expected no error,but got one:%v", FolderErr)
	}
	// set input 2
	inputParts2 := []string{
		"create-folder",
		"fatcat",
		"folder2",
		"DDDD",
	}
	var description2 bool = true
	// Create Folder
	if FolderErr := jsonObj.OsCreateFolder(inputParts2, description2, testUserInfoPath); FolderErr != nil {
		t.Errorf("Expected no error,but got one:%v", FolderErr)
	}
	// set input 3
	inputParts3 := []string{
		"create-folder",
		"AAAaa123",
		"folder2",
		"BBBB",
	}
	var description3 bool = true
	// Create Folder
	if FolderErr := jsonObj.OsCreateFolder(inputParts3, description3, testUserInfoPath); FolderErr != nil {
		t.Errorf("Expected no error,but got one:%v", FolderErr)
	}
}

func TestFoldernameCheck(t *testing.T) {

	t.Skip()

	jsonObj, err := ReadUserInfo(testUserInfoPath)
	if err != nil {
		t.Error("Expected no error,but got one")
	}
	// Exist Folder
	var username string = "fatcat"
	var foldername string = "folder2"
	// foldername check
	if foldernameErr := jsonObj.FoldernameCheck(username, foldername); foldernameErr != nil {
		t.Log("Expected one error,and got one")
	} else {
		t.Error("Expected one error,but got none")
	}
	// Non-Exist Folder
	var username2 string = "fatcat"
	var foldername2 string = "folder3"
	// foldername check
	if foldernameErr := jsonObj.FoldernameCheck(username2, foldername2); foldernameErr == nil {
		t.Log("Expected no error,and got none")
	} else {
		t.Error("Expected no error,but got one")
	}
}

func TestOsRenameFolder(t *testing.T) {

	t.Skip()

	jsonObj, err := ReadUserInfo(testUserInfoPath)
	if err != nil {
		t.Error("Expected no error,but got one")
	}
	// Change Foldername
	var inputParts = []string{"rename-folder", "fatcat", "folder2", "folderNew"}
	if changeFoldernameErr := jsonObj.OsRenameFolder(inputParts, testUserInfoPath); changeFoldernameErr != nil {
		t.Error("Expected no error, but got one")
	} else {
		t.Log("Expected no error, and got none")
	}
	// Change it back
	var inputParts2 = []string{"rename-folder", "fatcat", "folderNew", "folder2"}
	if changeFoldernameErr2 := jsonObj.OsRenameFolder(inputParts2, testUserInfoPath); changeFoldernameErr2 != nil {
		t.Error("Expected no error, but got one")
	} else {
		t.Log("Expected no error, and got none")
	}
	// Non-Exist Folder
	var inputParts3 = []string{"rename-folder", "fatcat", "folderNew", "folder2"}
	if changeFoldernameErr3 := jsonObj.OsRenameFolder(inputParts3, testUserInfoPath); changeFoldernameErr3 != nil {
		t.Log("Expected one error, and got one")
	} else {
		t.Error("Expected no error, but got one")
	}
}

func TestOsDeleteFolder(t *testing.T) {

	t.Skip()

	jsonObj, err := ReadUserInfo(testUserInfoPath)
	if err != nil {
		t.Error("Expected no error,but got one")
	}
	// Delete non-Exist folder
	var inputParts = []string{"delete-folder", "fatcat", "NoFolder"}
	if deleteFolderErr := jsonObj.OsDeleteFolder(inputParts, testUserInfoPath); deleteFolderErr != nil {
		t.Log("Expected one error, and got one")
	} else {
		t.Error("Expected one error, and got none")
	}
	// Delet Exist folder
	var inputParts2 = []string{"delete-folder", "fatcat", "folder2"}
	if deleteFolderErr2 := jsonObj.OsDeleteFolder(inputParts2, testUserInfoPath); deleteFolderErr2 != nil {
		t.Error("Expected no error, but got one")
	} else {
		t.Log("Expected no error, and got none")
	}

	// Add the folder back
	inputParts3 := []string{
		"create-folder",
		"fatcat",
		"folder2",
		"",
	}
	var description bool = false
	// Create Folder
	if FolderErr := jsonObj.OsCreateFolder(inputParts3, description, testUserInfoPath); FolderErr != nil {
		t.Errorf("Expected no error,but got one:%v", FolderErr)
	}

}

func TestFolderNum(t *testing.T) {

	t.Skip()

	jsonObj, err := ReadUserInfo(testUserInfoPath)
	if err != nil {
		t.Error("Expected no error,but got one")
	}
	// user fatcat got 2 folders
	folderNum := jsonObj.FolderNum("fatcat")
	if folderNum != 2 {
		t.Errorf("Expected 2 folder,but got %d", folderNum)
	} else {
		t.Log("Expected 2 folder,and got 2")
	}
}

func TestSortFolder(t *testing.T) {

	t.Skip()

	// Test with smaller case, not json file
	jsonObj := &JSONData{
		Data: map[string]setting.UserInfo{
			"user1": {
				Folders: []setting.Folder{
					{Name: "A", CreatedAt: time.Now()},
					{Name: "B", CreatedAt: time.Now().Add(1 * time.Hour)},
					{Name: "C", CreatedAt: time.Now().Add(2 * time.Hour)},
				},
			},
		},
	}

	// define a test func: 4 combinations need to be tested
	runTest := func(t *testing.T, sortType, sortRule string, expected []string) {
		t.Helper()

		jsonObj.SortFolder([]string{"list-folders", "user1", sortType, sortRule})

		// Validation
		for i, folder := range jsonObj.Data["user1"].Folders {
			if folder.Name != expected[i] {
				t.Errorf("Expected: %s, Got: %s", expected[i], folder.Name)
			}
		}
	}

	// By name + asc
	t.Run("Sort by Name Ascending", func(t *testing.T) {
		expected := []string{"A", "B", "C"}
		runTest(t, "--sort-name", "asc", expected)
	})

	// By name + desc
	t.Run("Sort by Name Descending", func(t *testing.T) {
		expected := []string{"C", "B", "A"}
		runTest(t, "--sort-name", "desc", expected)
	})

	// By time + asc
	t.Run("Sort by Created Time Ascending", func(t *testing.T) {
		expected := []string{"A", "B", "C"}
		runTest(t, "--sort-created", "asc", expected)
	})

	// By time + desc
	t.Run("Sort by Created Time Descending", func(t *testing.T) {
		expected := []string{"C", "B", "A"}
		runTest(t, "--sort-created", "desc", expected)
	})
}

func TestOsCreateFile(t *testing.T) {

	t.Skip()

	jsonObj, err := ReadUserInfo(testUserInfoPath)
	if err != nil {
		t.Error("Expected no error,but got one")
	}
	// Create File in Exist Folder without description
	var inputParts = []string{"create-file", "fatcat", "folder2", "file1"}
	var description bool = false
	if createFileErr := jsonObj.OsCreateFile(inputParts, description, testUserInfoPath); createFileErr == nil {
		t.Log("Expected no error, and got none")
	} else {
		t.Error("Expected no error, but got one")
	}
	// Create File in Exist Folder with description
	var inputParts2 = []string{"create-file", "fatcat", "folder2", "file2", "dddd"}
	var description2 bool = true
	if createFileErr := jsonObj.OsCreateFile(inputParts2, description2, testUserInfoPath); createFileErr == nil {
		t.Log("Expected no error, and got none")
	} else {
		t.Error("Expected no error, but got one")
	}
	// Create File in Non-Exist Folder
	var inputParts3 = []string{"create-file", "fatcat", "folder87", "file3"}
	var description3 bool = false
	if createFileErr := jsonObj.OsCreateFile(inputParts3, description3, testUserInfoPath); createFileErr != nil {
		t.Log("Expected one error, and got one")
	} else {
		t.Error("Expected one error, but got none")
	}
}
