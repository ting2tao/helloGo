package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	var group = new(errgroup.Group)
	var urls = []string{
		"http://www.baidu.com/",
		"https://golang2.eddycjy.com/",
		"https://eddycjy.com/",
	}
	for _, url := range urls {
		url := url
		group.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Body)
				resp.Body.Close()
			}
			return err
		})
	}
	if err := group.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Printf("Errors: %+v", err)
	}
}
