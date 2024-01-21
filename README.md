# Go Command Line Virtual File System (VFS)
## _Ease to Use, Cross Platform Tool_ :ghost:

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

VFS is a simple file management system written in [Go](https://go.dev/), which can be easily build in different OS. 

The Markdown format follows the **Github Markdown Rule**.

- Simple Command :sunglasses:
- Clear Screen Function :smile:
- Support Whitespace Input :muscle:

## Functions :fire:
- Register user
- Create Folder / Delete Folder / Rename Folder / List Folders
- Create File / Delete File / List Files
- Help Command for command information
- Whitespace Input Support Switch

## How to use it? :hammer:

1. Make sure to get a Go environment, version 1.20+ **(My go version: go1.21.6 windows/amd64)**
2. Simple ```go run``` or ```go build```, more details bellow.

## Installation :computer:
**No** third party package needed, VFS only use standard package.

## Build & Run :running:

Go run -- No Executable file generated
```
go run main.go
```
Go build -- Executable file generated
```
// Windows, -o ANY_NAME_YOU_WANT.exe
> go build -o vfs.exe main.go
> vfs.exe

// Linux -o ANY_NAME_YOU_WANT
> go build -o vfs main.go
> ./vfs
```

## System Architecture :clipboard:
All user's folder and file will be store under ```app``` folder.

```app_user_info.json``` will record all user information, this file will be created automatically, the file can be found in the ```user``` folder.
| Folder | Usage |
| ------ | ------ |
| app | All user's folders and files |
| process | process package |
| setting | setting package |
| user | user package |

## Quick Start :tea:

There are 3 basic command you should know first - ```help```, ```clear``` and ```exit```
```
// Usage : List command information
> help
// Usage : Clear screen
> clear
// Usage : Exit System
> exit
```

Also there are limitations for the input (Username, Foldername, Fileame and Description):
1. Upper case characters accepted (A-Z)
2. Lower case characters accepted (a-z)
3. Digital accepted (0-9)
4. Length range from 3 to 10 characters
5. Whitespace not accepted by default, see more information bellow.


## Basic User & Folder Usage :file_folder:

Let's register a user, type ```register [username]```
```
// Example
register fatcat

// Invalid Example
register fat@cat
```

We would want to create a folder for the user, use ```create-folder``` command.
```
Usage: Create one folder for one user, description is optional
Command: create-folder [username] [foldername] [description]

// Example without description
create-folder fatcat folder1

// Example with description
create-folder fatcat folder1 ddd

// Invalid Foldername
create-folder fatcat fold@er

// Invalid Description
create-folder fatcat folder dd@dd
```

We can also rename the folder using ```rename-folder``` command.
```
Usage: Rename one folder for one user
Command: rename-folder [username] [foldername] [new-folder-name]

// Example 
rename-folder fatcat folder1 folder2
```

Then, we delete the folder using ```delete-folder``` command.
```
Usage: Delete one folder for one user
Command: delete-folder [username] [foldername]

// Example 
delete-folder fatcat folder1
```

We can list all the user's folder by ```list-folders``` command
```
Usage: List all folders of one user with conditions
Command: list-folders [username] [--sort-name | --sort-created] [asc|desc]

// Example sorting with name ascending
list-folders fatcat --sort-name asc

// Example sorting with time ascending
list-folders fatcat --sort-created asc
```

## Basic User & File Usage :blue_book:

We can create a file in one of the user's folder by ```create-file``` command
```
Usage: Create one file in one folder for one user, description is optional
Command: create-file [username] [foldername] [filename] [description]

// Example without description
create-file fatcat folder1 file1

// Example with description
create-file fatcat folder1 file1 sdsd
```

Delete a file in one of the user's folder by ```delete-file``` command
```
Usage: Delete one file in one folder for one user
Command: delete-file [username] [foldername] [filename]

// Example
delete-file fatcat folder1 file1
```

We can also list all files under one folder by ```list-files``` command
```
Usage: List all files in one folder of one user with conditions
Command: list-files [username] [foldername] [--sort-name | --sort-created] [asc|desc]

// Example sorting with name ascending
list-files fatcat folder1 --sort-name asc

// Example sorting with time ascending
list-files fatcat folder1 --sort-created asc
```

## Whitespace Input Support :gem:

VFS supports whitespace input, but **not in default mode**.

Type ```status``` command to check whether the whitespace input is enabled.
```
status

// Output
Whitespace Support: true
Whitespace Support: false
```

To enable / disable whitespace input, type ```whitespace```
```
whitespace

// Output
Set Whitespace Support: true
Set Whitespace Support: false
```

:red_circle: **Important: use double quotes to close the input with whitespace.**

:red_circle: **Important: use double quotes to close the input with whitespace.**

:red_circle: **Important: use double quotes to close the input with whitespace.**

Let's see what's the difference between whitespace enabled and whitespace disabled
```
// whitespace disabled
register fatcat -- PASS
register "fat cat" -- ERROR

// whitespace enabled
register fatcat -- PASS
register "fat cat" -- PASS
```

# Unit Test :cop:

Unit Test file in both ```process``` and ```user``` folder. Test command as bellow.

The ```app``` folder and ```user_info.json``` in ```user``` folder is for testing, not related with the main program, so you can just leave them there.
```
// Unit test is not avaliable in the code for now.
// Need to remove t.Skip() in the TestFunction to execute the test
go test -v ./user
go test -v ./process
```

You can test the unit test by removing ```t.Skip()``` in the TestFunction.
Quick Example:
```
func TestExtractWhiteSpaceInput(t *testing.T) {

	t.Skip() // remove this line to test

	var testString = []string{
		"register fatcat",
		`register "fat cat"`,
		"some one select",
		`"register" the "fat cat"`}

	var expected = []int{2, 2, 3, 3}

	// test string extraction
	for i, str := range testString {
		parts, _ := extractWhiteSpaceInput(str)
		if len(parts) != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], len(parts))
			fmt.Println("Failed Parts:", parts)
			continue
		}

		for j, word := range parts {
			fmt.Printf("%d %s\n", j, word)
		}
		fmt.Printf("========\n")

	}
}
```

## Contact :smirk_cat:

Author : Shang-Fong Yang (GMfatcat)

Email : a60102244@gmail.com

Github : [GMfatcat](https://github.com/GMfatcat)

Project Done at 2024/1/21


