package main

import (
	"fmt"
	"net/http"
)

// func URLChecker(a, b string) []bool {
func URLChecker(a string) int {
	aStatus := func() int {
		res, err := http.Get(a)
		if err != nil {
			fmt.Println("Error")
		}
		return res.StatusCode
	}

	aStatusCode := aStatus()

	if aStatusCode != 200 {
		fmt.Println("✗ -", aStatusCode)
	} else {
		fmt.Println("✓ -", aStatusCode, "(Ok)")
	}

	return aStatusCode
}

func main() {
	URLChecker("https://adamherro.dev/")
}
