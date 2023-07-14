package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {

	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.microsoft.com",
		"http://www.google.com",
		"http://www.facebook.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {

			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("A maior página é:", biggest)

}
