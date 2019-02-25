package main

import (
	"errors"
	"fmt"
	"reflect"
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

	aMap := createFunctionMap()
	fmt.Printf("%#v\n", aMap["function1"])
	fmt.Printf("%#v\n", aMap["function2"])

	b, _ := Call(aMap, "function1", "s")
	fmt.Printf("%v\n", b)

	bb, _ := Call(aMap, "function2", 1)
	fmt.Printf("%v\n", bb)
	fmt.Printf("%v\n", bb[0].Bool())
}

// Call function map invoker
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("Invalid Params Count")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func createFunctionMap() map[string]interface{} {
	functionMap := make(map[string]interface{})
	functionMap["function1"] = func(s string) bool {
		if s == "s" {
			return true
		}
		return false
	}
	functionMap["function2"] = func(i int) bool {
		if i == 1 {
			return true
		}
		return false
	}
	return functionMap
}
