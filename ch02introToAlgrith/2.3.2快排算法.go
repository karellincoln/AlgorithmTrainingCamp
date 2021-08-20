package main

import (
	"fmt"
	"github.com/karellincoln/AlgorithmTrainingCamp/common"
)
func partition(a []int) (mid int) {
	i, j := 0, len(a)-1
	for i < j {
		for i < j && a[i] <= a[0] {
			i++
		}
		for i <j && a[j] > a[0] {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	if a[i] > a[0] {
		a[i-1], a[0] = a[0], a[i-1]
		return i-1
	}
	a[i], a[0] = a[0], a[i]
	return i;
}

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	mid := partition(a)
	quickSort(a[:mid])
	quickSort(a[mid+1:])
}

func main() {
	var n int
	var w []int
	_, err := fmt.Scanln(&n)
	w, err = common.GetArray(n)

	for err == nil {
		quickSort(w)
		fmt.Println(w)
		_, err = fmt.Scanln(&n)
		w, err = common.GetArray(n)
	}
}