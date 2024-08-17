package service

func FindPairs(numbers []int, target int) [][]int {

	var solutions [][]int

	seen := make(map[int][]int)

	for i, num := range numbers {
		complement := target - num
		if j, ok := seen[complement]; ok {
			for _, idx := range j {
				solutions = append(solutions, []int{idx, i})
			}
		}
		seen[num] = append(seen[num], i)
	}
	return solutions
}
