func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for 0 < n {
		rem := n % 10
		product *= rem
		sum += rem
		n /= 10
	}

	return product - sum
}
