package main

import (
	"fmt"
)

// Object bogus object
type Object struct {
	SubCategory string
	Years       bool
}

func main() {
	embeddedSlice := make(map[string][]Object)
	embeddedSlice["AKEY"] = []Object{Object{SubCategory: "SUB1", Years: true}, Object{SubCategory: "SUB9", Years: true}}
	embeddedSlice["ANOTHERKEY"] = []Object{Object{SubCategory: "SUB2", Years: true}, Object{SubCategory: "SUB3", Years: true}, Object{SubCategory: "SUB10", Years: false}}
	for key, value := range embeddedSlice {
		fmt.Printf("%s : %#v\n", key, value)
	}
}
