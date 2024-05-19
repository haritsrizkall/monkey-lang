package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/haritsrizkall/monkey-lang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is goddamn Monkey Language. Enjoyyy\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
