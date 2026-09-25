package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming langage !\n", user.Username)
	fmt.Println("Feel free to type in commands\nlet x = 3;\nx+2;")

	repl.Start(os.Stdin, os.Stdout)
}
