package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, v := range nums {
		another := target - v
		if _, ok := m[another]; ok {
			return []int{m[another], index}
		}
		m[v] = index
	}
	return nil
}
