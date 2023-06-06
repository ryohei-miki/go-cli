package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// options
	var (
		n = flag.Bool("n", false, "omit trailing new line")
		s = flag.String("s", " ", "separator")
	)
	// parse
	flag.Parse()
	// join args
	output := strings.Join(flag.Args(), *s)
	if *n {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}
