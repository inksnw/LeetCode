package leetcode

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
