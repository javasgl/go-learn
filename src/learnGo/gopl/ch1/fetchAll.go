package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	}

	fmt.Printf("%.2fs\texpend\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	start := time.Now()

	if !strings.HasPrefix(url, "http://") {

		url = "http://" + url

	}

	resp, err := http.Get(url)

	if err != nil {

		ch <- fmt.Sprint(err)

		return
	}

	count, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {

		ch <- fmt.Sprint(err)

		return

	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, count, url)

}
