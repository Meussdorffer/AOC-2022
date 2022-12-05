package main

import (
	"fmt"
	"regexp"
)

func main() {
	// move 1 from 2 to 1
	r := regexp.MustCompile(`move (\d) from (\d) to (\d)`)
	fmt.Printf("%#v\n", r.FindStringSubmatch(`move 1 from 2 to 1`))
}
