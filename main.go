package main

import (
	"fmt"
	"net/http"
)

// Should take in a slice of strings and return the status code (int)
// Loop through each string in the slice, run a goroutine for each loop using a channel.
// Output the status code for each URL

var URLsToCheck []string = []string{"https://adamherro.dev/", "https://vaudify.com/"}

func URLChecker(urlList []string) map[string]bool {
	urlResult := make(map[string]bool)

	for _, url := range urlList {
		status := checkURL(url)
		if status != 200 {
			fmt.Println("✗ -", url, "-", status)
			urlResult[url] = false
		} else {
			fmt.Println("✓ -", url, "-", status, "(Ok)")
			urlResult[url] = true
		}
	}

	return urlResult
}

func checkURL(url string) int {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return res.StatusCode
}

func main() {
	URLChecker(URLsToCheck)
}
