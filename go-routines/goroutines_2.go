package goroutines

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Result struct {
	URL  string
	Size int
}

type FetchError struct {
	URL string
	Err error
}

func fetch(url string, ch chan<- Result, errCh chan<- FetchError, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		errCh <- FetchError{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errCh <- FetchError{URL: url, Err: err}
		return
	}

	ch <- Result{URL: url, Size: len(body)}
}

func GoRoutinesExample2() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.medium.com",
		"https://www.catholicanswers.com",
		"https://www.otto.de",
		"https://www.phind.com",
	}

	ch := make(chan Result)
	errCh := make(chan FetchError)
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go fetch(url, ch, errCh, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
		close(errCh)
	}()

	for {
		select {
		case result, ok := <-ch:
			if !ok {
				ch = nil
			} else {
				fmt.Printf("URL: %s, Größe: %d Bytes\n", result.URL, result.Size)
			}
		case errResult, ok := <-errCh:
			if !ok {
				errCh = nil
			} else {
				fmt.Println("Fehler beim Abrufen der URL:", errResult.URL, errResult.Err)
			}
		}

		if ch == nil && errCh == nil {
			break
		}
	}
}
