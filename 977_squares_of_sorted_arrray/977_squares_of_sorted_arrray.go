func sortedSquares(A []int) []int {
	var squares []int

	for i := 0; i < len(A); i++ {
		squares = append(squares, A[i] * A[i])
	}

	sort.Ints(squares)

	return squares
}
