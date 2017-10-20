package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	p := fmt.Println

	match, _ := regexp.MatchString("p[a-z]+ch", "patch")
	p(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	p(r.MatchString("patch"))

	p(r.FindString("peach punch"))
	p(r.FindStringIndex("peach punch patch"))

	p(r.FindAllString("peach punch patch", -1))
	p(r.FindAllStringIndex("peach punch patch", -1))

	p(r.FindStringSubmatch("peach punch patch"))
	p(r.FindAllStringSubmatch("peach punch patch", -1))

	p(r.FindAllStringSubmatchIndex("peach punch patch", -1))

	r = regexp.MustCompile("p([a-z]+)ch")
	p(r)

	p(r.ReplaceAllString("str:peach punch patch", "oo"))
	p(r.ReplaceAllStringFunc("str:peach punch patch", func(str string) string { return "--" + str }))
	p(r.ReplaceAllStringFunc("str:peach punch patch", strings.ToUpper))
}
