package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	if len(os.Args) >= 2 {

		listFilesParaller(os.Args[1])

		wg.Wait()

	}
}

func listFilesParaller(dir string) {
	file, err := os.Open(dir)

	if err != nil {

		fmt.Println("open dir faild!", err)

	}

	defer file.Close()

	files, err := file.Readdir(-1)

	if err != nil {
		fmt.Println("read dir faild!", err)
	}

	for _, f := range files {
		if f.IsDir() {
			wg.Add(1)
			go listFilesParaller(dir + "/" + f.Name())
		}
		fmt.Println(dir + "/" + f.Name())
	}
	defer wg.Done()

}
