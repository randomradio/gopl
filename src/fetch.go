// fetch demonstrates use of net/http package
// as well as machenism of goroutine
// fetch multiple website at the same time
// according to argument passed in from command line
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

// fetching url and putting all error and fetching stats into channel
func fetch(url string, ch chan string) {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err.Error())
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprint(err.Error())
		return
	}

	secs := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%2.f, %7d, %s", secs, nbytes, url)
}
