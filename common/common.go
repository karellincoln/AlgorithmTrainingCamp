package common

import "fmt"

func GetArray(num int) ([]int, error) {
	res := make([]int, num)
	var err error
	for i := range res{
		_, err = fmt.Scan(&res[i])
	}
	return res, err
}
