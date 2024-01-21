package user

import (
	"fmt"
	"os"
	"sort"
	"time"
	s "virtual-file-system/setting"
)

/* Sort Functions */

func (jsonObj *JSONData) SortFolder(inputParts []string) {
	var username, sortType, sortRule string
	username = inputParts[1]
	sortType = inputParts[2]
	sortRule = inputParts[3]

	// Select Sort Conditions
	userInfo := jsonObj.Data[username]

	switch sortType {
	case "--sort-name":
		if sortRule == "asc" {
			sort.Sort(byFolderName(userInfo.Folders))
		} else {
			sort.Sort(sort.Reverse(byFolderName(userInfo.Folders)))
		}
	case "--sort-created":
		if sortRule == "asc" {
			sort.Sort(byFolderTime(userInfo.Folders))
		} else {
			sort.Sort(sort.Reverse(byFolderTime(userInfo.Folders)))
		}
	}

	// Show Sort Result
	fmt.Fprintf(os.Stdout, "Sort Type: %s, Sort Rule: %s\n", sortType, sortRule)
	for _, folder := range userInfo.Folders {
		fmt.Fprintf(os.Stdout, "Name:%s Time:%s\n",
			folder.Name, folder.CreatedAt.Format(time.RFC822))
	}
}

func (jsonObj *JSONData) SortFile(inputParts []string) {
	var username, foldername, sortType, sortRule string
	username = inputParts[1]
	foldername = inputParts[2]
	sortType = inputParts[3]
	sortRule = inputParts[4]
	// Select Sort Conditions
	folderIndex := jsonObj.findFolderIndex(username, foldername)
	userInfo := jsonObj.Data[username].Folders[folderIndex]

	switch sortType {
	case "--sort-name":
		if sortRule == "asc" {
			sort.Sort(byFileName(userInfo.Files))
		} else {
			sort.Sort(sort.Reverse(byFileName(userInfo.Files)))
		}
	case "--sort-created":
		if sortRule == "asc" {
			sort.Sort(byFileTime(userInfo.Files))
		} else {
			sort.Sort(sort.Reverse(byFileTime(userInfo.Files)))
		}
	}

	// Show Sort Result
	fmt.Fprintf(os.Stdout, "Folder: %s, Sort Type: %s, Sort Rule: %s\n",
		foldername, sortType, sortRule)
	for _, file := range userInfo.Files {
		fmt.Fprintf(os.Stdout, "Name:%s Time:%s\n",
			file.Name, file.CreatedAt.Format(time.RFC822))
	}

}

/* End of Sort Functions */

/* Sort Interface */

/* Folder Sort */
type byFolderName []s.Folder

func (a byFolderName) Len() int      { return len(a) }
func (a byFolderName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// define in asc
func (a byFolderName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type byFolderTime []s.Folder

func (a byFolderTime) Len() int      { return len(a) }
func (a byFolderTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// define in asc
func (a byFolderTime) Less(i, j int) bool { return a[i].CreatedAt.Before(a[j].CreatedAt) }

/* File Sort */
type byFileName []s.File

func (a byFileName) Len() int      { return len(a) }
func (a byFileName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// define in asc
func (a byFileName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type byFileTime []s.File

func (a byFileTime) Len() int      { return len(a) }
func (a byFileTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// define in asc
func (a byFileTime) Less(i, j int) bool { return a[i].CreatedAt.Before(a[j].CreatedAt) }

/* End of Sort Interface */
