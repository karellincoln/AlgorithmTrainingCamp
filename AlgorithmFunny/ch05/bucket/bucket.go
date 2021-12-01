/*
有三个容积分别是3升、5升和8升的水桶，其中容积为8升的水桶中装满了水，容积为3升和容积为5升的水桶是空的。
三个水桶都没有体积刻度，
现在需要将大水桶中的8升水等分成两份，每份都是4升水，附加条件是只能使用另外两个空水桶，不能借助其他辅助容器。
到底有多少种答案？
*/
package main

import (
	"fmt"
)

const BUCKET_COUNT = 3

var bucketCapicity = [BUCKET_COUNT]int{8, 5, 3}
var bucketInitState = [BUCKET_COUNT]int{8, 0, 0}
var bucketFinalState = [BUCKET_COUNT]int{4, 4, 0}

type Action struct {
	water int
	from  int
	to    int
}

type Bucket struct {
	curState  [BUCKET_COUNT]int
	curAction Action
}

func InitBucketStat() Bucket {
	return Bucket{
		curState: bucketInitState,
	}
}

func (b Bucket) PrintResult() {
	fmt.Printf("Dump %d water from %d to %d, buckets water states is: %v\n",
		b.curAction.water, b.curAction.from, b.curAction.to, b.curState)
}

func (b Bucket) CanJump(from, to int) bool {
	// 边界校验
	if from < 0 || from >= BUCKET_COUNT || to < 0 || to >= BUCKET_COUNT {
		return false
	}
	// 不符合条件校验
	if from == to || b.curState[from] == 0 || b.curState[to] == bucketCapicity[to] {
		return false
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (b Bucket) JumpWater(from, to int) (nextBucket Bucket, canJump bool) {
	if !b.CanJump(from, to) {
		return nextBucket, false
	}
	nextBucket.curAction.water = min(b.curState[from], bucketCapicity[to]-b.curState[to])
	nextBucket.curAction.from = from
	nextBucket.curAction.to = to
	nextBucket.curState = b.curState
	nextBucket.curState[from] -= nextBucket.curAction.water
	nextBucket.curState[to] += nextBucket.curAction.water
	return nextBucket, true
}

var count int

func printResult(bucketQueue []Bucket) {
	count++
	fmt.Printf("%2d: Find result：\n", count)
	for i := range bucketQueue {
		bucketQueue[i].PrintResult()
	}
	fmt.Println()
}

func InQueue(bucket Bucket, bucketQueue []Bucket) bool {
	for i := range bucketQueue {
		if bucketQueue[i].curState == bucket.curState {
			return true
		}
	}
	return false
}
func findAllDumpWater(bucketQueue []Bucket) {
	lastIndex := len(bucketQueue) - 1
	curBucket := bucketQueue[lastIndex]

	// 如果最后一个为最终状态，则打印结果
	if curBucket.curState == bucketFinalState {
		printResult(bucketQueue)
	}
	// 查找倒水的方式。
	for i := 0; i < BUCKET_COUNT; i++ {
		for j := 0; j < BUCKET_COUNT; j++ {
			nextBucket, canJump := curBucket.JumpWater(i, j)
			// 倒水，递归调用find方法。
			if canJump && !InQueue(nextBucket, bucketQueue) {
				bucketQueue = append(bucketQueue, nextBucket)
				findAllDumpWater(bucketQueue)
				bucketQueue = bucketQueue[:len(bucketQueue)-1]
			}
		}
	}
}

func main() {
	bucketQueue := make([]Bucket, 0, 100)
	bucketQueue = append(bucketQueue, InitBucketStat())
	findAllDumpWater(bucketQueue)
}
