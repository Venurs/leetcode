package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap :=make(map[int] int)
	for index, num := range nums{
		numsMap[num] = index
	}
	for nowIndex, num := range nums{
		var result = target - num
		index, exist := numsMap[result]
		if exist && nowIndex != index{
			return []int{nowIndex, index}
		}
	}
	return [] int{}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
