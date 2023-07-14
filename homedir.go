package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	fmt.Println("homedir")
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
	fmt.Println(usr.HomeDir)
}
