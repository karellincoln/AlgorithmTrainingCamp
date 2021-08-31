package common

import "fmt"

func GetIntArray(num int) ([]int, error) {
	res := make([]int, num)
	var err error
	for i := range res{
		_, err = fmt.Scan(&res[i])
	}
	return res, err
}

func GetStringArray(num int) ([]string, error) {
	res := make([]string, num)
	var err error
	for i := range res{
		_, err = fmt.Scan(&res[i])
	}
	return res, err
}