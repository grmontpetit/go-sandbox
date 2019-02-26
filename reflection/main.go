package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Object some random object
type Object struct {
	Field1      string
	Field2      int
	Field3      []string
	Data        *AStruct
	StructSlice []*Elem
}

// AStruct a struct
type AStruct struct {
	Field4      string
	StructField *AnotherStruct
}

// AnotherStruct another struct
type AnotherStruct struct {
	Field5 int
}

// Elem an element inside a slice
type Elem struct {
	Field99 int
}

func main() {
	object := new(Object)
	object.Field1 = "field1"
	object.Field3 = []string{"aaa", "bbb"}
	structField := new(AnotherStruct)
	structField.Field5 = 9999
	elems := make([]*Elem, 0)
	elems = append(elems, &Elem{Field99: 1})
	elems = append(elems, &Elem{Field99: 2})
	elems = append(elems, &Elem{Field99: 3})
	object.StructSlice = elems
	//object.Data = &AStruct{}
	//object.Data = &AStruct{Field4: "ooo", StructField: structField}

	// DO NOT DELETE
	// val := reflect.ValueOf(object).Elem()
	// for i := 0; i < val.NumField(); i++ {
	// 	valueField := val.Field(i)
	// 	typeField := val.Type().Field(i)
	// 	tag := typeField.Tag
	// 	fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\t IsEmpty: %t \n", typeField.Name, valueField.Interface(), tag.Get("tag_name"), isEmpty(valueField.Interface()))
	// }

	valuePath := "Object/Data/Field1"
	b := valueExistInStructPath(object, valuePath)
	exist := "does NOT exist"
	if b {
		exist = "exists"
	}
	fmt.Printf("Value %s in %s\n", exist, valuePath)

}

func valueExistInStructPath(parentStruct interface{}, path string) bool {
	fmt.Println(path)
	fullPath := strings.Split(path, "/")

	if len(fullPath) == 1 {
		return isEmpty(reflect.ValueOf(parentStruct))
	}

	if strings.Split(reflect.TypeOf(parentStruct).String(), ".")[1] == fullPath[0] {
		fullPath = fullPath[1:]
	}
	val := reflect.ValueOf(parentStruct).Elem()
	exist := false
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		// fmt.Printf("%v == %v\n", typeField.Name, fullPath[0])
		// if typeField.Name == fullPath[0] {
		// 	fmt.Println(valueField.Interface())
		// }

		//fmt.Println(valueField)
		if len(fullPath) == 1 && fullPath[0] == typeField.Name {
			exist = !isEmpty(valueField.Interface())
		} else {
			exist = valueExistInStructPath(valueField.Interface(), strings.Join(fullPath[1:], "/"))
		}
	}
	return exist
}

func isEmpty(data interface{}) bool {
	empty := false
	v := reflect.ValueOf(data)
	switch reflect.ValueOf(data).Kind() {
	case reflect.Ptr:
		val := reflect.ValueOf(data).Elem()
		if reflect.ValueOf(data).IsNil() {
			empty = true
			return empty
		}
		for i := 0; i < val.NumField(); i++ {
			empty = isEmpty(val.Field(i).Interface())
		}
	case reflect.Int:
		if v.Int() == 0 {
			empty = true
		}
	case reflect.Bool:
		if !v.Bool() {
			empty = true
		}
	case reflect.String:
		if v.String() == "" {
			empty = true
		}
	case reflect.Slice:
		if v.Len() == 0 {
			empty = true
		}
	case reflect.Struct:
		if reflect.Zero(reflect.TypeOf(v)).Interface() == v.Interface() {
			empty = true
		}
	default:
		return empty
	}
	return empty
}
