package main

import (
	"fmt"
	"strings"
)

func main() {

	strs := []string{"peach", "apple", "plum", "pear"}
	fmt.Println(Index(strs, "apple"))
	fmt.Println(Include(strs, "apple"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "p")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
	fmt.Println(Map(strs, strings.ToTitle))
}

func Index(strs []string, t string) int {
	for i, v := range strs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(strs []string, t string) bool {
	return Index(strs, t) >= 0
}

func Any(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(strs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range strs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Map(strs []string, f func(string) string) []string {
	mapped := make([]string, len(strs))
	for i, v := range strs {
		mapped[i] = f(v)
	}
	return mapped
}
