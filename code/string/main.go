package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReplacer(
		"$", "\\$",
		"*", "\\*",
	)

	target := "one*two$two*three"

	rst := r.Replace(target)
	fmt.Println(rst)
}
