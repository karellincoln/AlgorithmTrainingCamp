// 训练1-42：蛇形填数，输入一个整数n，按照蛇形填写n×n的矩阵。
package main

import "fmt"


// 1. 一个复杂的事情，需要先拆分成一个一个的步骤。再思考每一个步骤的启示条件，终止添加。
// 2. 可以借助原数组中的数据做一些状态标记。
// 3. 为了能够使每一步都能够比较有规律的处理，可以提前设置一个值，或者一些列。
func main() {
	var n int
	_, err := fmt.Scanln(&n)
	for err == nil {
		var matrix [][]int
		matrix = make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, n)
		}
		var x, y int
		total := 1
		matrix[0][0] = total
		for total < (n*n) {
			fmt.Println(total)
			// 右
			for y+1 < n && matrix[x][y+1] == 0 {
				total++
				y++
				matrix[x][y] = total
			}
			// 下
			for x+1 < n && matrix[x+1][y] == 0 {
				total++
				x++
				matrix[x][y] = total
			}
			// 左
			for y-1 >= 0 && matrix[x][y-1] == 0 {
				total++
				y--
				matrix[x][y] = total
			}
			// 上
			for x-1 >= 0 && matrix[x-1][y] == 0 {
				total++
				x--
				matrix[x][y] = total
			}
		}
		for i := range matrix {
			fmt.Println(matrix[i])
		}
		_, err = fmt.Scanln(&n)
	}

}