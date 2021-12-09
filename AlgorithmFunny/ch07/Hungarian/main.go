/* 最大匹配与匈牙利算法
匈牙利算法的实现以顶点集合V 为基础，每次X 集合中选一个顶点Xi 做增广路径的起点搜索增广路径。重点在找增广路径上;
基本流程：
将图G最大匹配初始化为空
while(从Xi点开始在图G中找到新的增广路径) {
    将增广路径假如到最大匹配中;
}
输出图G的最大匹配;

增广路径流程：
while(从Xi的邻接表中找到下一个关联顶点Yj)
{
    if(顶点Yj不在增广路径上)
    {
        将Yj加入增广路径;
        if(Yj是未覆盖点或者从与Yj相关连的顶点（Xk）能找到增广路径)
        {
            将Yj的关联顶点修改为Xi；
            从顶点Xi开始有增广路径，返回true；
        }
    }
}
从顶点Xi开始没有增广路径，返回false；
*/

package main

import "fmt"

type TagPartner struct {
	Name string
}

func HungaryMatch(x, y []TagPartner, edge [][]bool) (max int) {
	visit := make([]bool, len(y)) // 记录右侧元素是否已被访问过
	p := make(map[int]int)        // 记录当前右侧元素所对应的左侧元素

	var FindAugmentPath func(i int) bool
	FindAugmentPath = func(i int) bool {
		for j := range edge[i] {
			if edge[i][j] && !visit[j] {
				visit[j] = true
				if _, ok := p[j]; !ok || FindAugmentPath(p[j]) {
					p[j] = i
					return true
				}
			}
		}
		return false
	}

	for i := range x {
		if FindAugmentPath(i) {
			max++
		}
		visit = make([]bool, len(y)) // 重置vis
	}

	fmt.Println("最大匹配：", max)
	for i := range p {
		if _, ok := p[i]; ok {
			fmt.Printf("%v  <--->  %v\n", x[p[i]], y[i])
		}
	}
	return
}

func main() {
	girls := []TagPartner{
		{"girl0"},
		{"girl1"},
		{"girl2"},
		{"girl3"},
		{"girl4"},
	}
	boys := []TagPartner{
		{"boy0"},
		{"boy1"},
		{"boy2"},
		{"boy3"},
		{"boy4"},
	}
	edge := [][]bool{
		{true, false, false, false, false},
		{false, true, false, false, false},
		{true, false, true, false, false},
		{false, false, true, true, false},
		{false, true, false, false, false},
	}
	HungaryMatch(boys, girls, edge)
}
