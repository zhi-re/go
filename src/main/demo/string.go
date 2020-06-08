package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "223.71.97.14:17159"
	cut := strings.Index(s, ":")
	content := s[0:cut]
	fmt.Println(content)
}
