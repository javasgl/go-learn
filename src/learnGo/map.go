package main

import "fmt"

func main() {
	arrMap := make(map[int]string)
	print(arrMap)

	arrMap[1] = "一"
	arrMap[2] = "二"

	print(arrMap)
	for index, val := range arrMap {

		if index == 1 {
			fmt.Println(index)
			fmt.Println(val)
			continue
		}
	}

}
func print(arrMap map[int]string) {
	fmt.Printf("%#v\n", arrMap)
}
