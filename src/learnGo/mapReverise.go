package main

import (
	"fmt"
)

/**
 * map key, value 反转
 */
func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	m2 := make(map[string]int, len(m1))
	for key, value := range m1 {
		m2[value] = key
	}
	fmt.Println(m2)
}
