package main

import (
	"fmt"
	"strings"
)

const repeatCount = 5

func Repeat(s string) string {
	if s == "" {
		return ""
	}
	var res strings.Builder
	for range repeatCount {
		res.WriteString(s)
	}
	return res.String()
}

func main() {
	fmt.Println(Repeat("6"))
}
