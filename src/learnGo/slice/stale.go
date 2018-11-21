package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //3 3 [1,2,3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //2 2 [2,3]
	s2[0] += 10
	fmt.Println(s1, s2) //[1,12,3] [12,3]

	s3 := s1[:1]
	fmt.Println(len(s3), cap(s3), s3) //1 3 [1]
	s3[0] += 10
	fmt.Println(s1, s3) //[11,12,3] [11]

	s4 := append(s3, 4)
	fmt.Println(s1, s2, s3, s4) //[11 4 3] [4,3] [11] [11 4]
}
