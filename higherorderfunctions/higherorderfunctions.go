package main

import (
	"fmt"
)

func main() {
	f := func(msg string) {
		fmt.Println("f says: " + msg)
	}

	f("hello")

	do("world", f)
}

func do(s string, f func(s string)) {
	f("(do) " + s)
}
