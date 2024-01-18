/* System Setting*/
package setting

const (
	UserInfoPath = "./user/user_info.json"
)

/* Help Command */
var HelpCommand = `
Valid Commands:
[] means that the input is required
[]? means that the input is optional

Usage: Register one user
Command: register [username]

TBD
TBD
TBD
`

/* Basic Data Structures */
type UserInfo struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Files       []File `json:"files"`
}

type File struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
