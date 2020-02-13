package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kieranajp/monkey/cmd"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to the Monkey REPL, %s\n", user.Username)

	cmd.StartRepl(os.Stdin, os.Stdout)
}
