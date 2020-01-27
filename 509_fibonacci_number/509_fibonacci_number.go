func fib(N int) int {
	var fibSec []int
	fibSec = append(fibSec, []int{0,1}...)

	for i := 2; i <= N; i++ {
		fibSec = append(fibSec, fibSec[i-1] + fibSec[i-2])
	}

	return fibSec[N]
}
