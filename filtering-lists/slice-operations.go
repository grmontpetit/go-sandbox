package main

import (
	"fmt"
	"math/rand"
)

// Object some object
type Object struct {
	data    string
	counter int
}

// NewObject creates an object
func NewObject() *Object {
	return new(Object)
}

// ObjectList type definition for a list of Objects
type ObjectList []*Object

// NewObjectList constructor for ObjectList's type
func NewObjectList() ObjectList {
	return make([]*Object, 0)
}

// NewBoundObjectList constructor for ObjectList's type with defined lenght
func NewBoundObjectList(size int) ObjectList {
	return make([]*Object, size)
}

// Key Create a key from an object
func (o *Object) Key() string {
	return fmt.Sprintf("%s%d", o.data, o.counter)
}

func main() {
	object1 := new(Object)
	object1.data = "object1"
	object1.counter = 1

	object2 := new(Object)
	object2.data = "object2"
	object2.counter = 2

	object3 := new(Object)
	object3.data = "object3"
	object3.counter = 2

	object4 := new(Object)
	object4.data = "object4"
	object4.counter = 3

	myList := NewObjectList()
	myList = append(myList, object4)
	myList = append(myList, object1)
	myList = append(myList, object3)
	myList = append(myList, object2)

	fmt.Println("Distinct items: ")
	for _, item := range myList.Distinct() {
		fmt.Printf("%#v\n", item)
	}

	fmt.Println("Filtered items: ")
	aFilter := func(o *Object) bool {
		if o.counter != 1 {
			return true
		}
		return false
	}
	for _, item := range myList.Filter(aFilter) {
		fmt.Printf("%#v\n", item)
	}

	fmt.Println("Ordered items: ")
	ordering := func(o1 *Object, o2 *Object) bool {
		if o1.counter > o2.counter {
			return true
		}
		return false
	}
	for _, item := range myList.QuickSort1(ordering) {
		fmt.Printf("%#v\n", item)
	}
}

// Distinct Filter out any duplicates using a map
func (l ObjectList) Distinct() ObjectList {
	m := make(map[string]Object)
	list := NewObjectList()
	for _, object := range l {
		m[object.Key()] = *object
	}
	for _, value := range m {
		content := NewObject()
		*content = value
		list = append(list, content)
	}
	return list
}

// Filter Filters a list based on a predicate
func (l ObjectList) Filter(filter func(object *Object) bool) ObjectList {
	list := NewObjectList()
	for _, o := range l {
		if filter(o) {
			list = append(list, o)
		}
	}
	return list
}

// QuickSortCustom custom ordering
func (l ObjectList) QuickSortCustom(ordering func(o1 *Object, o2 *Object) bool) ObjectList {
	if len(l) < 2 {
		return l
	}

	list := NewBoundObjectList(len(l))

	pivot := rand.Int() % len(l)

	stPivot := func(o *Object) bool {
		if ordering(o, l[pivot]) {
			return true
		}
		return false
	}
	etPivot := func(o *Object) bool {
		if o.counter == l[pivot].counter {
			return true
		}
		return false
	}
	gtPivot := func(o *Object) bool {
		if ordering(l[pivot], o) {
			return true
		}
		return false
	}

	list = append(list, l.Filter(stPivot).QuickSortCustom(ordering)...)
	list = append(list, l.Filter(etPivot)...)
	list = append(list, l.Filter(gtPivot).QuickSortCustom(ordering)...)

	return list.Filter(func(o *Object) bool {
		if o != nil {
			return true
		}
		return false
	})
}

// QuickSort ascending implementation
func (l ObjectList) QuickSort() ObjectList {
	if len(l) < 2 {
		return l
	}

	list := NewBoundObjectList(len(l))

	pivot := rand.Int() % len(l)

	stPivot := func(o *Object) bool {
		if o.counter < l[pivot].counter {
			return true
		}
		return false
	}
	etPivot := func(o *Object) bool {
		if o.counter == l[pivot].counter {
			return true
		}
		return false
	}
	gtPivot := func(o *Object) bool {
		if o.counter > l[pivot].counter {
			return true
		}
		return false
	}

	list = append(list, l.Filter(stPivot).QuickSort()...)
	list = append(list, l.Filter(etPivot)...)
	list = append(list, l.Filter(gtPivot).QuickSort()...)

	return list.Filter(func(o *Object) bool {
		if o != nil {
			return true
		}
		return false
	})
}
