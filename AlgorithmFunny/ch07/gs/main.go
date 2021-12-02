/*Gale-Shapley算法原理
盖尔和沙普利的策略是一种寻找稳定婚姻的策略，不管男女之间有何种偏好，这种策略总可以得到一个稳定的婚姻匹配。
初始化所有的m∈ M，w∈ W，所有的m和w都是自由状态;
while (存在男人是自由的，并且他还没有对每个女人都求过婚) {
    选择一个这样的男人m;
    w = m的优先选择表中还没有求过婚的排名最高的女人;
    if (w 是自由状态) {
        将(m, w)的状态设置为约会状态;
    }
    else { //w 已经和其他男人约会了
		m' = w当前约会的男人;
		if (w 更喜爱m' 而不是m) {
			m 保持单身状态（w不更换约会对象）;
		}
		else { // w 更喜爱m 而不是m
			将(m, w)的状态设置为约会状态;
			将m' 设置为自由状态;
		}
	}
}
输出已经匹配的集合S;
*/
package main

import "fmt"

type TagPartner struct {
	Name    string
	Next    int   // 下一个邀请对象
	Current int   // 当前舞伴，-1表示还没有舞伴
	Perfect []int // 偏爱列表
}

func (t TagPartner) IsBestLove(other int) bool {
	if t.Current == -1 {
		return true
	}
	// 采用这种方式判断，感觉比原书中的更好一些
	for _, v := range t.Perfect {
		if v == other {
			return true
		}
		if v == t.Current {
			return false
		}
	}
	return false
}

type GroupPartner struct {
	Girls []*TagPartner
	Boys  []*TagPartner
}

func (g *GroupPartner) GetFreedomMan() (int, *TagPartner) {
	for i, v := range g.Boys {
		// 第二个判断避免越界
		if v.Current == -1 && v.Next < len(g.Boys[i].Perfect) {
			return i, v
		}
	}
	return -1, nil
}

func (g *GroupPartner) PrintResult() {
	for _, boy := range g.Boys {
		// 可能存在没有匹配上的情况
		girlName := "no partner"
		if boy.Current != -1 {
			girlName = g.Girls[boy.Current].Name
		}
		fmt.Printf("%s <--> %s \n", boy.Name, girlName)
	}
}


func main() {
	partner := GroupPartner{
		Girls: []*TagPartner{
			{"girl0", 0, -1, []int{2, 1, 4, 0, 3}},
			{"girl1", 0, -1, []int{1, 3, 0, 4, 2}},
			{"girl2", 0, -1, []int{0, 1, 2, 4, 3}},
			{"girl3", 0, -1, []int{2, 0, 3, 1, 4}},
			{"girl4", 0, -1, []int{4, 1, 2, 3, 0}},
		},
		Boys: []*TagPartner{
			{"boy0", 0, -1, []int{4, 1, 0, 3, 2}},
			{"boy1", 0, -1, []int{3, 1, 4, 0, 2}},
			{"boy2", 0, -1, []int{2, 3, 4, 0, 1}},
			{"boy3", 0, -1, []int{4, 0, 3, 1, 2}},
			{"boy4", 0, -1, []int{3, 4, 2, 0, 1}},
		},
	}
	for i, boy := partner.GetFreedomMan(); boy != nil; i, boy = partner.GetFreedomMan() {
		girlIndex := boy.Perfect[boy.Next]
		girl := partner.Girls[girlIndex]
		if girl.Current == -1 {
			girl.Current = i
			boy.Current = girlIndex
		} else {
			if girl.IsBestLove(i) {
				// 原来的恋人置为单身
				otherBoy := partner.Boys[girl.Current]
				otherBoy.Current = -1

				// 现在的成为恋人
				girl.Current = i
				boy.Current = girlIndex
			}
		}
		boy.Next++
	}
	partner.PrintResult()
}
