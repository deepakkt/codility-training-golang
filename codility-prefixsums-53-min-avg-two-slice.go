package solution

func Solution(A []int) int {
	inputLength := len(A)
	prefixSum := make([]int, inputLength + 1)

	currentSum := 0

	for i := 1; i <= inputLength; i++ {
		currentSum += A[i - 1]
		prefixSum[i] = currentSum
	}

	prefixSum[0] = 0

	minAverage := 10001.0
	minSlice := 0

	for o := 0; o <=inputLength - 2; o++ {
		for n := o + 1; n <= o + 2; n++ {
			if n >= inputLength {
				continue
			}

			currentSum = prefixSum[n + 1] - prefixSum[o]

			currentAverage := float64(currentSum) / float64(n - o + 1)

			if currentAverage < minAverage {
				minAverage = currentAverage
				minSlice = o
			}
		}
	}

	return minSlice
}
