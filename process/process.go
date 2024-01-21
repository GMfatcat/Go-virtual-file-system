/* Process User Input Commands */
package process

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"virtual-file-system/setting"
	"virtual-file-system/user"
)

var (
	invalidCommandError = fmt.Errorf("Invalid command,type help for more information\n")
	noKeywordError      = fmt.Errorf("No keyword found in the input\n")
	noArgumentError     = fmt.Errorf("No keyword argument found in the input\n")
	unClosedQuoteError  = fmt.Errorf("Unclosed double quote\n")
	registerArgsError   = fmt.Errorf("Register requires 2 arguments\n")
)

func ProcessInput(input string) error {

	var parts []string
	var extractErr error

	// Check if whitespace support is enabled
	// Extract command into parts, Keyword will be the first part
	if setting.SupportWhitespace {
		parts, extractErr = extractWhiteSpaceInput(input)
		if extractErr != nil {
			return extractErr
		}
	} else {
		parts, extractErr = extractInput(input)
		if extractErr != nil {
			return extractErr
		}
	}

	// Read UserInfo File
	jsonObj, readUserInfoErr := user.ReadUserInfo(setting.AppUserInfoPath)
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

		if len(parts) != 2 {
			return registerArgsError
		}

		username := parts[1]

		registerNameErr := jsonObj.RegisterName(username, setting.AppUserInfoPath)
		if registerNameErr != nil {
			return registerNameErr
		}

	case "create-folder":
		createFolderErr := jsonObj.CreateFolder(parts, setting.AppUserInfoPath)
		if createFolderErr != nil {
			return createFolderErr
		}

	case "delete-folder":
		deleteFolderErr := jsonObj.DeleteFolder(parts, setting.AppUserInfoPath)
		if deleteFolderErr != nil {
			return deleteFolderErr
		}

	case "list-folders":
		listFoldersErr := jsonObj.ListFolders(parts)
		if listFoldersErr != nil {
			return listFoldersErr
		}

	case "rename-folder":
		renameFolderErr := jsonObj.RenameFolder(parts, setting.AppUserInfoPath)
		if renameFolderErr != nil {
			return renameFolderErr
		}

	case "create-file":
		createFileErr := jsonObj.CreateFile(parts, setting.AppUserInfoPath)
		if createFileErr != nil {
			return createFileErr
		}

	case "delete-file":
		deleteFileErr := jsonObj.DeleteFile(parts, setting.AppUserInfoPath)
		if deleteFileErr != nil {
			return deleteFileErr
		}

	case "list-files":
		listFileErr := jsonObj.ListFiles(parts)
		if listFileErr != nil {
			return listFileErr
		}

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

// 1 keyword + at least 1 args, support arguments in ""
func extractWhiteSpaceInput(input string) ([]string, error) {

	emptyStringList := []string{}

	var parts []string
	var currentPart string
	var insideQuotes bool = false

	for _, char := range input {
		if char == ' ' && !insideQuotes {
			if currentPart != "" {
				parts = append(parts, currentPart)
				currentPart = ""
			}
		} else if char == '"' {
			insideQuotes = !insideQuotes
		} else {
			currentPart += string(char)
		}
	}

	if currentPart != "" {
		parts = append(parts, currentPart)
	}

	if insideQuotes {
		return emptyStringList, unClosedQuoteError
	}

	if len(parts) == 0 {
		return emptyStringList, noKeywordError
	}

	if len(parts) == 1 {
		return emptyStringList, noArgumentError
	}

	return parts, nil
}

// Clear the screen
func ClearConsole() {
	// detect os system
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {

		cmd = exec.Command("cmd", "/c", "cls")

	} else {

		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	cmd.Run()
}
