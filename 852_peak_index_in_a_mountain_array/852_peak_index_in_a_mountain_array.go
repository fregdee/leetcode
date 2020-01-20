func peakIndexInMountainArray(A []int) int {
	for i := 0; i < len(A); i++ {
		if A[i] > A[i+1] {
			return i
		}
	}
	return len(A) - 1
}
