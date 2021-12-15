/*
题目是这样的:
据说有五个不同颜色的房间排成一排，每个房间里分别住着一个不同国籍的人，每个人都喝一种特定品牌的饮料，抽一种特定品牌的烟，养一种宠物，
没有任意两个人抽相同品牌的香烟，或喝相同品牌的饮料，或养相同的宠物。问题是谁在养鱼作为宠物？
为了寻找答案，爱因斯坦给出了以下15条线索。
1. 英国人住在红色的房子里；
2. 瑞典人养狗作为宠物；
3. 丹麦人喝茶；
4. 绿房子紧挨着白房子，在白房子的左边；
5. 绿房子的主人喝咖啡；
6. 抽Pall Mall牌香烟的人养鸟；
7. 黄色房子里的人抽Dunhill牌香烟；
8. 住在中间那个房子里的人喝牛奶；
9. 挪威人住在第一个房子里面；
10. 抽Blends牌香烟的人和养猫的人相邻；
11. 养马的人和抽Dunhill牌香烟的人相邻；
12. 抽BlueMaster牌香烟的人喝啤酒；
13. 德国人抽Prince牌香烟；
14. 挪威人和住在蓝房子的人相邻；
15. 抽Blends牌香烟的人和喝矿泉水的人相邻。
*/
package main

import (
	"fmt"
)

const GROUPS_ITEMS = 5
const GROUPS_COUNT = 5

var itemName = [GROUPS_ITEMS]string{"房子", "国家", "饮料", "宠物", "烟"}
var valueName = [GROUPS_ITEMS][GROUPS_COUNT]string{
	{"蓝色", "红色", "绿色", "黄色", "白色"},
	{"挪威", "丹麦", "瑞士", "英国", "德国"},
	{"茶", "水", "咖啡", "啤酒", "牛奶"},
	{"马", "猫", "鸟", "鱼", "狗"},
	{"Blends", "Dunhill", "Prince", "PallMall", "BlueMaster"},
}

const (
	COLOR = iota
	NATION
	DRINK
	PET
	CIGARET
)

const (
	COLOR_BLUE = iota + 1
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_WHITE
)
const (
	NATION_NORWAY = iota + 1
	NATION_DANMARK
	NATION_SWEDEND
	NATION_ENGLAND
	NATION_GERMANY
)
const (
	DRINK_TEA = iota + 1
	DRINK_WATER
	DRINK_COFFEE
	DRINK_BEER
	DRINK_MILK
)
const (
	PET_HORSE = iota + 1
	PET_CAT
	PET_BIRD
	PET_FISH
	PET_DOG
)
const (
	CIGARET_BLENDS = iota + 1
	CIGARET_DUNHILL
	CIGARET_PRINCE
	CIGARET_PALLMALL
	CIGARET_BLUEMASTER
)

func check1(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) bool {
	if checkBind(graph, COLOR, COLOR_RED, NATION, NATION_ENGLAND) &&
		checkBind(graph, NATION, NATION_SWEDEND, PET, PET_DOG) &&
		checkBind(graph, NATION, NATION_DANMARK, DRINK, DRINK_TEA) &&
		checkBind(graph, COLOR, COLOR_GREEN, DRINK, DRINK_COFFEE) &&
		checkBind(graph, PET, PET_BIRD, CIGARET, CIGARET_PALLMALL) &&
		checkBind(graph, COLOR, COLOR_YELLOW, CIGARET, CIGARET_DUNHILL) &&
		checkBind(graph, DRINK, DRINK_BEER, CIGARET, CIGARET_BLUEMASTER) &&
		checkBind(graph, NATION, NATION_GERMANY, CIGARET, CIGARET_PRINCE) {
		return true
	}
	return false
}
func checkBind(graph *[GROUPS_ITEMS][GROUPS_COUNT]int, firstType, firstValue, secondType, secondValue int) bool {
	index := -1
	for i := range graph[firstType] {
		if graph[firstType][i] == firstValue {
			index = i
		}
	}
	if index == -1 || graph[secondType][index] == 0 || graph[secondType][index] == secondValue {
		return true
	}
	return false
}

func check4(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) bool {
	for i := range graph[COLOR] {
		if graph[COLOR][i] == COLOR_GREEN {
			if (i+1 >= GROUPS_COUNT) || (graph[COLOR][i+1] != 0 && graph[COLOR][i+1] != COLOR_WHITE) {
				return false
			}
		}
	}
	return true
}

func check8(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) bool {
	if graph[DRINK][2] == 0 || graph[DRINK][2] == DRINK_MILK {
		return true
	}
	return false
}
func check9(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) bool {
	if graph[NATION][0] == 0 || graph[NATION][0] == NATION_NORWAY {
		return true
	}
	return false
}
func check10(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) bool {
	if checkNeighbor(graph, PET, PET_CAT, CIGARET, CIGARET_BLENDS) &&
		checkNeighbor(graph, PET, PET_HORSE, CIGARET, CIGARET_DUNHILL) &&
		checkNeighbor(graph, COLOR, COLOR_BLUE, NATION, NATION_NORWAY) &&
		checkNeighbor(graph, DRINK, DRINK_WATER, CIGARET, CIGARET_BLENDS) {
		return true
	}
	return false
}
func checkNeighbor(graph *[GROUPS_ITEMS][GROUPS_COUNT]int, firstType, firstValue, secondType, secondValue int) bool {
	index := -1
	for i := range graph[firstType] {
		if graph[firstType][i] == firstValue {
			index = i
		}
	}
	if index == -1 {
		return true
	}
	left, right := false, false
	if index-1 >= 0 {
		if graph[secondType][index-1] == 0 || graph[secondType][index-1] == secondValue {
			left = true
		}
	}
	if index+1 < GROUPS_COUNT {
		if graph[secondType][index+1] == 0 || graph[secondType][index+1] == secondValue {
			right = true
		}
	}
	if left || right {
		return true
	}
	return false
}

func checkfullGraph(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) bool {
	for i := GROUPS_ITEMS - 1; i >= 0; i-- {
		for j := GROUPS_COUNT - 1; j >= 0; j-- {
			if graph[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
func printGraph(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) {
	for i := range itemName {
		fmt.Print(itemName[i], "\t")
	}
	fmt.Println()
	for i := range graph {
		for j := range graph[i] {
			fmt.Print(valueName[j][graph[j][i]-1], "\t")
		}
		fmt.Println()
	}
}

var checkCount = 0

func checkValidGraph(graph *[GROUPS_ITEMS][GROUPS_COUNT]int) bool {
	checkCount++
	result := false
	if check1(graph) && check4(graph) && check8(graph) && check9(graph) && check10(graph) {
		result = true
	}
	if result && checkfullGraph(graph) {
		printGraph(graph)
	}
	return result
}

func notInGraph(graph *[GROUPS_ITEMS][GROUPS_COUNT]int, itemIdx int, idx int, con int) bool {
	for i := 0; i < idx; i++ {
		if graph[itemIdx][i] == con {
			return false
		}
	}
	return true
}

func generateRank(graph *[GROUPS_ITEMS][GROUPS_COUNT]int, itemIdx int, idx int) {
	if idx == GROUPS_COUNT {
		generateRank(graph, itemIdx+1, 0)
		return
	}
	if itemIdx == GROUPS_ITEMS {
		return
	}
	for i := 1; i <= GROUPS_COUNT; i++ {
		if notInGraph(graph, itemIdx, idx, i) {
			graph[itemIdx][idx] = i
			if checkValidGraph(graph) {
				generateRank(graph, itemIdx, idx+1)
			}
			graph[itemIdx][idx] = 0
		}
	}
}

func main() {
	var graph [GROUPS_ITEMS][GROUPS_COUNT]int
	// 1. 生成每个类别的排列
	// 2. 调用check函数做剪枝
	generateRank(&graph, 0, 0)
	fmt.Println("final checkCount: ", checkCount)
}
