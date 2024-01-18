package main

import (
	"fmt"
	"log"
	"virtual-file-system/user"
)

func main() {
	username := "Tom"
	fmt.Println("Starting...")
	if err := user.SetName(username); err != nil {

		log.Println(err)

	} else {
		fmt.Printf("Add %s successfully\n", username)
	}

}
