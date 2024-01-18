/* Process User Input Commands */
package process

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"virtual-file-system/setting"
	"virtual-file-system/user"
)

var (
	invalidCommandError = fmt.Errorf("Invalid command,type help for more information\n")
	noKeywordError      = fmt.Errorf("No keyword found in the input\n")
	noArgumentError     = fmt.Errorf("No keyword argument found in the input\n")
)

/*Example: register username*/
func ProcessInput(input string) error {

	// Extract command into parts, Keyword will be the first part
	parts, extractErr := extractInput(input)
	if extractErr != nil {
		return extractErr
	}

	// Read UserInfo File
	jsonObj, readUserInfoErr := user.ReadUserInfo(setting.UserInfoPath)
	if readUserInfoErr != nil {
		return readUserInfoErr
	}
	// Init UserInfo Data if necessary
	if jsonObj.Data == nil {
		jsonObj.Data = make(map[string]setting.UserInfo)
	}

	// Keyword decide the next move
	keyword := parts[0]

	switch keyword {

	case "register":
		registerNameErr := jsonObj.RegisterName(parts[1], setting.UserInfoPath)
		if registerNameErr != nil {
			return registerNameErr
		}

	case "create-folder":
		return nil

	case "delete-folder":
		return nil

	case "list-folders":
		return nil

	case "rename-folder":
		return nil

	case "create-file":
		return nil

	case "delete-file":
		return nil

	case "list-files":
		return nil

	default:
		return invalidCommandError
	}

	return nil
}

// 1 keyword + at least 1 args
func extractInput(input string) ([]string, error) {

	emptyStringList := []string{}

	parts := strings.Fields(input)

	if len(parts) == 0 {
		return emptyStringList, noKeywordError
	}

	if len(parts) == 1 {
		return emptyStringList, noArgumentError
	}

	return parts, nil
}

// Clear the screen (Only for windows system)
func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
