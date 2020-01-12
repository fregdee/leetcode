func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)

	for _, n := range arr {
		count[n] += 1
	}

	result := make(map[int]int)

	for _, v := range count {
		result[v] += 1
		if result[v] > 1 {
			return false
		}
	}

	return true
}
