package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Object some random object
type Object struct {
	Field1       string
	Field2       int
	Field3       []string
	BooleanField bool
	Data         *AStruct
	StructSlice  []*Elem
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
	object.Field2 = 1
	object.Field3 = []string{"aaa", "bbb"}
	structField := new(AnotherStruct)
	structField.Field5 = 9999
	elems := make([]*Elem, 0)
	elems = append(elems, &Elem{Field99: 1})
	elems = append(elems, &Elem{Field99: 2})
	elems = append(elems, &Elem{Field99: 3})
	object.StructSlice = elems
	//object.Data = &AStruct{}
	object.Data = &AStruct{Field4: "ooo", StructField: structField}

	// DO NOT DELETE
	// val := reflect.ValueOf(object).Elem()
	// for i := 0; i < val.NumField(); i++ {
	// 	valueField := val.Field(i)
	// 	typeField := val.Type().Field(i)
	// 	tag := typeField.Tag
	// 	fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\t IsEmpty: %t \n", typeField.Name, valueField.Interface(), tag.Get("tag_name"), isEmpty(valueField.Interface()))
	// }

	valuePath := "Object/BooleanField"
	b, v := valueExistInStructPath(object, valuePath)
	found := fmt.Sprintf("Not found: no data exist at %s", valuePath)
	if b {
		found = fmt.Sprintf("Found: data exist at %s: value=%v", valuePath, v)
	}
	fmt.Println(found)

}

func valueExistInStructPath(currentValue interface{}, path string) (bool, interface{}) {

	fullPath := strings.Split(path, "/")
	var empty bool
	var value interface{}

	if len(fullPath) == 1 {
		empty, value = isEmpty(currentValue)
		return !empty, value
	}

	fullPath = fullPath[1:]

	val := reflect.ValueOf(currentValue).Elem()
	empty, value = isEmpty(val)
	if empty {
		return false, nil
	}

	valueField := val.FieldByName(fullPath[0])
	empty, value = isEmpty(valueField)
	if empty {
		return false, nil
	}

	return valueExistInStructPath(valueField.Interface(), strings.Join(fullPath, "/"))
}

func isEmpty(data interface{}) (bool, interface{}) {
	empty := false
	var value interface{}

	v := reflect.ValueOf(data)
	switch reflect.ValueOf(data).Kind() {
	case reflect.Ptr:
		val := reflect.ValueOf(data).Elem()
		if reflect.ValueOf(data).IsNil() {
			empty = true
			return empty, nil
		}
		for i := 0; i < val.NumField(); i++ {
			empty, value = isEmpty(val.Field(i).Interface())
		}
	case reflect.Int:
		value = v.Int()
		if value == 0 {
			empty = true
		}
	case reflect.Bool:
		value = v.Bool()
		if !v.Bool() {
			empty = true
		}
	case reflect.String:
		value = v.String()
		if value == "" {
			empty = true
		}
	case reflect.Slice:
		if v.Len() == 0 {
			value = v.Slice
			empty = true
		}
	case reflect.Struct:
		if reflect.Zero(reflect.TypeOf(v)).Interface() == v.Interface() {
			value = v.Interface()
			empty = true
		}
	default:
		return empty, value
	}
	return empty, value
}
