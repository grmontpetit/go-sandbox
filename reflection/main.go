package main

import (
	"fmt"
	"reflect"
)

// Object some random object
type Object struct {
	Field1 string
	Field2 int
	Field3 []string
}

func main() {
	object := new(Object)
	object.Field1 = "field1"
	object.Field2 = 2
	object.Field3 = []string{"aaa", "bbb"}

	matchFunc := func(s string) bool {
		if s == "Field2" {
			return true
		}
		return false
	}
	valueOf := reflect.ValueOf(*object)
	//ref := reflect.New(typeOf)
	structField := valueOf.FieldByNameFunc(matchFunc)

	//structField, _ := dd.FieldByNameFunc(matchFunc)
	fmt.Printf("%v\n", structField)

}
