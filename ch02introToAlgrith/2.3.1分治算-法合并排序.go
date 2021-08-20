package main

import "fmt"

// 分治秘籍
//（1）分解：将原问题分解为若干规模较小、相互独立且与原问题形式相同的子问题。
//（2）治理：求解各个子问题。由于各个子问题与原问题形式相同，只是规模较小，所以当子问题划分得足够小时，就可以用较简单的方法解决。
//（3）合并：按原问题的要求，将子问题的解逐层合并成原问题的解。

func Merge(a []int,mid int) {
	b := make([]int, len(a))
	copy(b, a)
	// b -> a
	i,j := 0, mid
	for k:=0; k < len(a); k++ {
		if i >= mid {
			a[k] = b[j]
			j++
		} else if j >= len(a) {
			a[k] = b[i]
			i++
		} else if b[i] <= b[j] {
			a[k] = b[i]
			i++
		} else {
			a[k] = b[j]
			j++
		}
	}
}
func MergeSort(A[]int) {
	if len(A) <= 1 {
		return
	}
	mid := len(A) / 2
	MergeSort(A[:mid])
	MergeSort(A[mid:])
	Merge(A,mid)
}
func main() {
	var n int
	var w []int
	_, err := fmt.Scanln(&n)
	w, err = getArray(n)

	for err == nil {
		MergeSort(w)
		fmt.Println(w)
		_, err = fmt.Scanln(&n)
		w, err = getArray(n)
	}
}