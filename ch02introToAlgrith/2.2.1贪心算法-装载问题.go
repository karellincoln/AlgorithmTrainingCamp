package main

import (
	"fmt"
	"github.com/karellincoln/AlgorithmTrainingCamp/common"
	"sort"
)

// 最优子结构性质。当一个问题的最优解包含其子问题的最优解时，称此问题具有最优子结构性质。问题的最优子结构性质是该问题是否可以用贪心算法求解的关键。
// 例如原问题S={a1,a2,…,ai,…,an}，通过贪心选择选出一个当前最优解{ai}之后，转化为求解子问题S-{ai}，
// 如果原问题的最优解包含子问题的最优解，则说明该问题满足最优子结构性质。

// //  贪心算法的求解步骤：
// 1. 定好贪心策略。
// 2. 根据最优子结构性质，找出局部最优解。
// 3. 将局部最优解合成，全局最优解。

// // 问题：
// 有一天，海盗们截获了一艘装满各种各样古董的货船，每件古董都价值连城，一旦打碎就失去了价值。虽然海盗船足够大，但载重为c，每件古董的重量为wi，海盗们绞尽脑汁要把尽可能多的宝贝装上海盗船，该怎么办呢？

func bestLoad(c int, w []int) {
	sort.Ints(w)
	ret := make([]int, 0)
	for _, t := range w {
		if t < c {
			ret = append(ret, t)
			c -= t
		}
	}
	fmt.Printf("len: %d, value %v \n", len(ret), ret)
}
func main() {
	var c, n int
	var w []int
	_, err := fmt.Scanln(&c, &n)
	w, err = common.GetIntArray(n)

	for err == nil {
		bestLoad(c, w)
		_, err = fmt.Scanln(&c, &n)
		w, err = common.GetIntArray(n)
	}
}