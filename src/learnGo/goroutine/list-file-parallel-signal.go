package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var signalChan = make(chan struct{}, runtime.GOMAXPROCS(runtime.NumCPU()))

func main() {

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	if len(os.Args) >= 2 {

		listFilesWithSignal(os.Args[1])

		wg.Wait()

	}

}

func listFilesWithSignal(dir string) {

	signalChan <- struct{}{}

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
			go listFilesWithSignal(dir + "/" + f.Name())
		}
		fmt.Println(dir + "/" + f.Name())
	}
	defer func() {
		<-signalChan
		wg.Done()
	}()
}
