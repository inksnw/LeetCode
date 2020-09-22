package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for n, v := range nums {
		another := target - v
		if _, ok := m[another]; ok {
			return []int{m[another], n}
		}
		m[v] = n
	}
	return nil

}

func main() {

	a := twoSum([]int{2, 4, 5}, 6)
	fmt.Print(a)

}
