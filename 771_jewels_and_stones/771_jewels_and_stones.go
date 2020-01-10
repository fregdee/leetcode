func numJewelsInStones(J string, S string) int {
	count := 0
	js := strings.Split(J, "")

	for _, j := range js {
	count += strings.Count(S, j)
	}

	return count
}
