package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"virtual-file-system/process"
	"virtual-file-system/setting"
	"virtual-file-system/user"
)

func main() {

	// Check if userinfo json exists, create if not exist
	if err := user.CheckUserInfoExists(setting.AppUserInfoPath); err != nil {
		fmt.Fprintf(os.Stderr, "Check userinfo Error：%v", err)
	}

	// Start System
	systemWelcome := "Enter a command ('help' for command information, 'exit' to quit)\n"
	fmt.Fprintf(os.Stdout, systemWelcome)

	// Create a Infinite Loop to simulate CLI Environment
	for {

		fmt.Print("# ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Input Error：%v", err)
			return
		}

		input = strings.TrimSpace(input)

		switch input {

		case "help":
			fmt.Fprintf(os.Stdout, "System Commands\n%s\n", setting.HelpCommand)
			continue

		case "clear":
			process.ClearConsole()
			continue

		case "status":
			fmt.Fprintf(os.Stdout, "Whitespace Support:%v\n", setting.SupportWhitespace)
			continue

		case "whitespace":
			setting.SupportWhitespace = !setting.SupportWhitespace
			fmt.Fprintf(os.Stdout, "Set Whitespace Support:%v\n", setting.SupportWhitespace)
			continue

		case "exit":
			fmt.Fprintf(os.Stdout, "Exit System")
			return

		default:
			// Process Input Here
			if err := process.ProcessInput(input); err != nil {
				fmt.Fprintf(os.Stderr, "Error：%v", err)
			}
		}
	}
}
