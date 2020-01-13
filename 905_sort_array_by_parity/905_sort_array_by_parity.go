func sortArrayByParity(A []int) []int {
	var even, odd []int

	for i := 0; i < len(A); i++ {
		if A[i] % 2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}

	return append(even, odd...)
}
