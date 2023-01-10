package main

func twoSum(nums []int, target int) []int {
	skiped := make(map[int]int, len(nums))
	for i, num := range nums {
		j, ok := skiped[num]
		if ok {
			return []int{j, i}
		}
		skiped[target-num] = i
	}
	return nil
}
