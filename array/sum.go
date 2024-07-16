package array

func sum(numArr []int) int {
	sum := 0

	for _, num := range numArr {
		sum += num
	}

	return sum
}
