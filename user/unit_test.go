/* Unit Test for user package */
package user

import (
	"testing"
	s "virtual-file-system/setting"
)

var (
	userInfoPath = "test_user_info.json"
)

func TestCheckUserInfoExists(t *testing.T) {
	// Check if userinfo json exists, create if not exist
	if err := CheckUserInfoExists(userInfoPath); err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	} else {
		t.Log("Expected no error, and got none")
	}

}

func TestReadUserInfo(t *testing.T) {

	_, err := ReadUserInfo(userInfoPath)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	} else {
		t.Log("Expected no error, and got none")
	}
}

func TestRegisterName(t *testing.T) {

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
