/**
 * @author songgl
 * @since 2018-05-14 17:13
 */
package main

import (
    "fmt"
)

const (
    StatusOpen       = iota + 1
    _
    StatusDeprecated
)

func main() {
    fmt.Println(StatusOpen)
    //fmt.Println(StatusClose)
    fmt.Println(StatusDeprecated)
}
