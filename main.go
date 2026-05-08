package main

import (
	"fmt"
	"net/http"
)

// Should take in a slice of strings and return the status code (int)
// Loop through each string in the slice, run a goroutine for each loop using a channel.
// Output the status code for each URL

var URLsToCheck []string = []string{"https://adamherro.dev", "https://vaudify.com", "https://adamherro.dev/contact", "https://adamhro.dev", "https://google.com", "https://hockey-ecommerce-store.onrender.com"}

type receivedURL struct {
	url    string
	status int
	error  error
}

type URLList []string

func URLChecker(list URLList) map[string]int {
	result := make(map[string]int)

	resultChan := make(chan receivedURL)

	for _, url := range list {
		go func() {
			status, err := checkURL(url)
			resultChan <- receivedURL{url, status, err}
		}()
	}
	for i := 0; i < len(list); i++ {
		received := <-resultChan
		if received.error != nil {
			fmt.Println("✗ -", received.url, "-", "Server Does Not Exist")
			result[received.url] = 0
		} else if received.status != 200 {
			fmt.Println("✗ -", received.url, "-", received.status)
			result[received.url] = received.status
		} else {
			fmt.Println("✓ -", received.url, "-", received.status)
			result[received.url] = received.status
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
