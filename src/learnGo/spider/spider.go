package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
)

const (
	DOMAIN = "https://javasgl.github.io"
)

var (
	wg sync.WaitGroup
)

func main() {

	respon, err := http.Get("https://javasgl.github.io/archives/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer respon.Body.Close()
	body, err := ioutil.ReadAll(respon.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	re, err := regexp.Compile("<a class=\"post-title-link\" href=\"(.*)\" itemprop=\"url\">")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	links2 := re.FindAllStringSubmatch(string(body), -1)

	fmt.Println(links2)
	fmt.Println("-------")

	wg.Add(len(links2) * 10)
	for _, link := range links2 {
		url := DOMAIN + link[1]
		for i := 0; i <= 10; i++ {
			fmt.Println(url)
			go access(url)
		}
	}
	wg.Wait()
}

func access(url string) {
	defer wg.Done()
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("done")
	}
}
