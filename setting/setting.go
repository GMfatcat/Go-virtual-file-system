/* System Setting*/
package setting

import "time"

// App for real application,otherwise testing
const (
	UserInfoPath    = "./user/user_info.json"
	AppUserInfoPath = "./user/app_user_info.json"
)

var SupportWhitespace = false

/* Help Command */
var HelpCommand = `
Type "clear" to clear screen (only for windows system)
Type "status" to check whitespace support status
Type "whitespace" to enable/disable support for whitespace in input
--> ex. "user name"

===========  System Commands =============

[] means that the input is required
[]? means that the input is optional

Usage: Register one user
Command: register [username]

Usage: Create one folder for one user, description is optional
Command: create-folder [username] [foldername] [description]?

Usage: Delete one folder for one user
Command: delete-folder [username] [foldername]

Usage: List all folders of one user with conditions
Command: list-folders [username] [--sort-name | --sort-created] [asc|desc]

Usage: Rename one folder for one user
Command: rename-folder [username] [foldername] [new-folder-name]

Usage: Create one file in one folder for one user, description is optional
Command: create-file [username] [foldername] [filename] [description]?

Usage: Delete one file in one folder for one user
Command: delete-file [username] [foldername] [filename]

Usage: List all files in one folder of one user with conditions
Command: list-files [username] [foldername] [--sort-name | --sort-created] [asc|desc]
`

/* Basic Data Structures */
type UserInfo struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Files       []File `json:"files"`
	CreatedAt   time.Time
}

type File struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time
}
