package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"virtual-file-system/setting"
)

func main() {

	// Welcome User
	fmt.Println("Enter a command ('help' for command information, 'exit' to quit)")

	// Create a Infinite Loop to simulate CLI Environment
	for {

		fmt.Print("# ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input Error：", err)
			return
		}

		input = strings.TrimSpace(input)

		switch input {
		case "help":
			fmt.Printf("System Commands\n%s\n", setting.HelpCommand)
			continue
		case "exit":
			fmt.Println("Exit System")
			return
		// Process Input
		default:
			fmt.Printf("Your input: %s\n", input)
		}

	}

}
