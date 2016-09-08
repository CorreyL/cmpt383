package main

import (
	"fmt"
	"strings"
)

func main() {
	var words []string
	words = strings.Fields("  foo bar  baz   ")
	fmt.Printf("Fields are: %q", words)
}