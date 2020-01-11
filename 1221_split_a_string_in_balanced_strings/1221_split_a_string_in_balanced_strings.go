func balancedStringSplit(s string) int {
	var L, R, count int

	for _, slice := range strings.Split(s, "") {
		if slice == "L" {
			L++
		} else {
			R++
		}

		if L == R {
			count++
			L = 0
			R = 0
		}
	}

	return count
}
