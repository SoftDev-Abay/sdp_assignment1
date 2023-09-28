package main

import (
	"fmt"
	"reflect"
)

type Sorter interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (b *BubbleSort) Sort(array []int) []int {
	arrayLength := len(array)
	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	return array
}

type SelectionSort struct{}

func (q *SelectionSort) Sort(array []int) []int {
	arrayLength := len(array)
	for i := 0; i < arrayLength; i++ {
		min := i
		for j := i; j < arrayLength-1; j++ {
			if array[min] > array[j+1] {
				min = j + 1
			}
		}
		temp := array[min]
		array[min] = array[i]
		array[i] = temp
	}
	return array
}

type InsertionSort struct{}

func (i *InsertionSort) Sort(array []int) []int {

	arrayLength := len(array)
	for i := 0; i < arrayLength; i++ {
		for j := i - 1; j >= 0; j-- {
			if array[i] < array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
				i = j
			} else {
				break
			}
		}
	}
	return array
}

func SortArray(s Sorter, array []int) {
	name := ""
	if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	} else {
		name = t.Name()
	}

	fmt.Printf("Sorted with %v : %v \n", name, s.Sort(array))

}

func main() {
	array1 := []int{10, 2, 15, 4, 11, 25, 8, 9}
	array2 := []int{10, 2, 15, 4, 11, 25, 8, 9}
	array3 := []int{10, 2, 15, 4, 11, 25, 8, 9}
	b := BubbleSort{}
	SortArray(&b, array1)
	i := InsertionSort{}
	SortArray(&i, array2)
	s := SelectionSort{}
	SortArray(&s, array3)
}

//
