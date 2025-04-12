package main

import (
	"flag"
	"fmt"
	"strings"
)

// Usage:
// ./main -n hello world
// ./main -sep ", " hello world
// ./main -help
var (
	n   = flag.Bool("n", false, "omit trailing newlines")
	sep = flag.String("sep", " ", "separator")
)

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
