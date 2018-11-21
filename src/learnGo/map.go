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
	m := make(map[int]uint)
	for i := 0; i <= 10; i++ {
		if i%2 > 0 {
			m[i%2] += 1
		} else {
			m[i%2] += 2
		}
	}
	fmt.Println(m)

}
func print(arrMap map[int]string) {
	fmt.Printf("%#v\n", arrMap)
}
