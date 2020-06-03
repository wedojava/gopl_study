package main

import "fmt"

func main() {
	fmt.Println(VariadicJoin(","))               // ""
	fmt.Println(VariadicJoin(",", "foo"))        // "foo"
	fmt.Println(VariadicJoin(",", "foo", "bar")) // "foo,bar"
	fmt.Println(VariadicJoin(" ", "foo", "bar")) // "foo bar"
	fmt.Println(VariadicJoin(", ", "we", "are", "the", "world!"))
}

func ExampleVariadicJoin() {
	fmt.Println(VariadicJoin(", ", "we", "are", "the", "world!"))
	// Output:we, are, the, world!!
}

func VariadicJoin(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var str = strs[0]
	for i, s := range strs {
		if i == 0 {
			continue
		}
		str += sep + s
	}
	return str
}
