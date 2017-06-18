package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	teststring := "123456"
	md5Instnace := md5.New()
	md5Instnace.Write([]byte(teststring))
	result := md5Instnace.Sum([]byte(""))
	fmt.Printf("%x\r\n", result)

	sha1Instance := sha1.New()
	sha1Instance.Write([]byte(teststring))
	result = sha1Instance.Sum([]byte(""))
	fmt.Printf("%x\r\n", result)

	//file
	testfile := "hash.go"
	gofile, err := os.Open(testfile)
	if err == nil {
		md5h := md5.New()
		io.Copy(md5h, gofile)
		fmt.Printf("%x %s\n", md5h.Sum([]byte("")), testfile)

		sha1 := sha1.New()
		io.Copy(sha1, gofile)
		fmt.Printf("%x %s\n", sha1.Sum([]byte("")), testfile)

	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
