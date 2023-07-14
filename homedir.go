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

	fmt.Println(usr)
	fmt.Println(usr.HomeDir)
}
