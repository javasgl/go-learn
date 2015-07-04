package main;

import "fmt";

func main() {
	arrMap:=make(map[int]string);
	print(arrMap);
}
func print(arrMap map[int]string){
	fmt.Printf("map::%v\n",arrMap);
}