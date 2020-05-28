// Write a function `expand(s string, f func(string) string) string`
// that replace each substring "$foo" within `s` by the text returned by f("foo").

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "$fooasdfasdfasdf$foobar$fff$foobar2"
	fmt.Println(expand(s, strings.ToUpper))
}

// expand replace each substring within `s`
func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}
