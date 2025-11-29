// Package main provides the entry point for the greeter application.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Aga602/reddit-wallstreets-bets/internal/greeter"
)

func main() {
	name := flag.String("name", "", "Name of the person to greet")
	flag.Parse()

	g := greeter.New(*name)
	fmt.Println(g.Greet())
	os.Exit(0)
}
