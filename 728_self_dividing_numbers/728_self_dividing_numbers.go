func selfDividingNumbers(left int, right int) []int {
	var result []int

	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) == true {
			fmt.Println(i)
			result = append(result, i)
		}
	}

	return result
}

func isSelfDividingNumber(n int) bool {
	origin := n
	for n > 0 {
		mod := n % 10
		if mod == 0 || origin % mod != 0 {
			return false
		}
		n = n / 10
	}
	return true
}
