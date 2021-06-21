

package solution

// you can also use imports, for example:
import "sort"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	distinct := 1

	for i := 1; i < len(A); i++ {
		if A[i] != A[i - 1] {
			distinct += 1
		}
	}

	return distinct
}

