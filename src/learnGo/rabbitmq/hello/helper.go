/**
 * @author songgl
 * @since 2018-05-25 17:09
 */
package hello

import (
    "fmt"
    "log"
    "strings"
)

func FailOnError(err error, msg string) {

    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf(":%s: %s", msg, err))
    }
}

func BodyFrom(args []string) string {
    return strings.Join(args[1:], " ")
}
