package main

import (
	"fmt"
	"github.com/karellincoln/AlgorithmTrainingCamp/common"
	"strings"
)

func find(str []string, v string) bool {
	for _, s := range str {
		if s == v {
			return true
		}
	}
	return false
}
/* input:
8 4 3
zhan qian sun li zhou wu zhen wang
zhao qian sun li
zhao zhou zheng
2 2 2
zhao qian
zhao qian
zhao qian
 */
func main() {
	var a, b, c int;
	var x, y, z []string;
	_, err := fmt.Scanln(&a, &b, &c)
	for err == nil {
		x, err = common.GetStringArray(a)
		y, err = common.GetStringArray(b)
		z, err = common.GetStringArray(c)
		var res []string
		for _, v := range y {
			if find(x, v) {
				if !find(z, v) {
					res = append(res, v)
				}
			}
		}
		if len(res) == 0 {
			fmt.Println("No enemy spy")
		} else {
			fmt.Println(strings.Join(res, " "))
		}
		_, err = fmt.Scanln(&a, &b, &c)
	}
}