package main

import "fmt"

func main() {

	res, err := f(10)
	fmt.Println(err)
	fmt.Println(res)
	if argerr, ok := err.(*ArgsError); ok {
		fmt.Println(argerr.msg)
	}
}

func f(arg int) (int, error) {
	if arg == 10 {
		return arg, &ArgsError{arg: arg, msg: "don`t support this arg!"}
	}
	return arg, nil
}

type ArgsError struct {
	arg int
	msg string
}

func (this *ArgsError) Error() string {
	return fmt.Sprintf("%d %s", this.arg, this.msg)
}
