package main

import (
	"fmt"
	"monkey/repr"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", usr.Name)
	repr.Start(os.Stdin, os.Stdout)
}
