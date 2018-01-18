package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		listFiles(os.Args[1])
	}
}

func listFiles(dir string) {

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
			listFiles(dir + "/" + f.Name())
		}
		fmt.Println(dir + "/" + f.Name())
	}

}
