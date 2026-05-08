package main

import (
	"fmt"
	"net/http"
)

// Should take in a slice of strings and return the status code (int)
// Loop through each string in the slice, run a goroutine for each loop using a channel.
// Output the status code for each URL

var URLsToCheck []string = []string{"https://adamherro.dev/", "https://vaudify.com/", "https://adamherro.dev/contact", "https://adamhro.dev/"}

type Result struct {
	resultList map[string]int
}

func (r *Result) makeResultMap() map[string]int {
	r.resultList = make(map[string]int)
	return r.resultList
}

type URLList []string

func URLChecker(list URLList) *Result {
	result := &Result{}
	result.makeResultMap()

	for _, url := range list {
		status, err := checkURL(url)
		if err != nil {
			fmt.Println("✗ -", url, "-", "Server Does Not Exist")
			result.resultList[url] = 0
		} else if status != 200 {
			fmt.Println("✗ -", url, "-", status)
			result.resultList[url] = status
		} else {
			fmt.Println("✓ -", url, "-", status)
			result.resultList[url] = status
		}
	}

	return result
}

func checkURL(a string) (int, error) {
	res, err := http.Get(a)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}

func main() {
	URLChecker(URLsToCheck)
}
