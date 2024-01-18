/* System Setting*/
package setting

var (
	UserInfoPath = "./user/user_info.json"
)

/* Basic Data Structures */

type UserInfo struct {
	Folders []Folder
}

type Folder struct {
	Name        string
	Description string
	Files       []File
}

type File struct {
	Name        string
	Description string
}
