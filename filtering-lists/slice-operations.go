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
	for _, item := range myList.QuickSort(ordering) {
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

// OrderBy Orders a list by ordering the items
func (l ObjectList) OrderBy(ordering func(o1 *Object, o2 *Object) bool) ObjectList {
	return l.QuickSort(ordering)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

// QuickSort implementation
func (l ObjectList) QuickSort(ordering func(o1 *Object, o2 *Object) bool) ObjectList {
	if len(l) < 2 {
		return l
	}

	list := NewBoundObjectList(len(l))

	left, right := 0, len(l)-1

	pivot := rand.Int() % len(l)

	l[pivot], l[right] = l[right], l[pivot]

	for i := range l {
		if ordering(l[i], l[right]) {
			list[left], list[i] = l[i], l[left]
			left++
		}
		if ordering(l[right], l[i]) {
			list[left], list[i] = l[left], l[i]
			left++
		}
	}

	l[left], l[right] = l[right], l[left]

	list = append(list, l[:left].QuickSort(ordering)...)
	list = append(list, l[left+1:].QuickSort(ordering)...)

	return list
}
