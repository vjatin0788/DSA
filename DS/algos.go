package DS

import (
	"fmt"
	"math"
	"sort"
)

func FindPermutation(prefix, str []string) {
	if len(str) == 0 {
		printString(prefix)
		return
	}
	for idx := 0; idx < len(str); idx++ {
		var tempStr []string
		tempStr = append(tempStr, str[:idx]...)
		tempStr = append(tempStr, str[idx+1:]...)
		FindPermutation(append(prefix, str[idx]), tempStr)
	}
}

func printString(prefix []string) {
	for idx := range prefix {
		fmt.Print(prefix[idx])
	}
	fmt.Println()
}

func coverPoints(A []int, B []int) int {
	var count int
	for idx := 1; idx < len(A); idx++ {
		count += FindMin(float64(A[idx]), float64(B[idx]), float64(A[idx-1]), float64(B[idx-1]))
	}
	return count
}

func FindMin(a, b, c, d float64) int {
	e := math.Abs(a - c)
	f := math.Abs(b - d)
	return int(math.Max(e, f))
}

//------------------> Add one to number <--------------

func plusOne(A []int) []int {
	var ar []int
	var carry bool

	carry = true
	for idx := len(A) - 1; idx >= 0; idx-- {
		var temp int
		temp = A[idx]
		if carry {
			if (A[idx] + 1) >= 10 {
				temp = 0
				carry = true
			} else {
				temp = A[idx] + 1
				carry = false
			}
		}
		ar = append(ar, temp)
	}

	if carry {
		ar = append(ar, 1)
	}

	return invertArray(ar)
}

func invertArray(a []int) []int {
	br := make([]int, 0)

	var pos int
	for idx := len(a) - 1; idx >= 0; idx-- {
		br = append(br, a[idx])
	}
	//fmt.Println(br)
	for idx := 0; idx < len(a); idx++ {
		if br[idx] != 0 {
			pos = idx
			break
		}
	}

	return br[pos:len(a)]
}

func maxSubArray(A []int) int {
	var maxSoFar, maxEndHere int

	maxSoFar = A[0]
	for idx := 0; idx < len(A); idx++ {
		maxEndHere = max(A[idx], A[idx]+maxEndHere)

		maxSoFar = max(maxSoFar, maxEndHere)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAbsDiff(A []int) int {

	cache := make([][]int, len(A))
	var maxSoFar int
	for idx := range cache {
		cache[idx] = make([]int, idx+1)
	}

	for idx := range cache {
		for jdx := 0; jdx <= idx; jdx++ {
			cache[idx][jdx] = -1
		}
	}

	for idx := range A {
		for jdx := range A {
			max := 0
			if jdx < len(cache[idx]) && cache[idx][jdx] != -1 {
				max = cache[idx][jdx]
			} else {
				max = int(abs(A, idx, jdx))
			}
			cache[jdx][idx] = max

			if maxSoFar < max {
				maxSoFar = max
			}
		}
	}

	return maxSoFar
}

func abs(A []int, i, j int) float64 {
	return math.Abs(float64(A[i]-A[j])) + math.Abs(float64(i-j))
}

func repeatedNumber(A []int) (result []int) {
	sort.Ints(A)
	// fmt.Println(A)
	missing := -1

	for idx := 1; idx < len(A); idx++ {
		if A[idx] == A[idx-1] {
			result = append(result, A[idx])
		}
		if A[idx]-2 == A[idx-1] {
			missing = A[idx] - 1
		}

		if missing != -1 && len(result) == 1 {
			result = append(result, missing)
		}
		if len(result) == 2 {
			break
		}
	}
	if missing == -1 {
		result = append(result, A[0]-1)
	}

	//fmt.Println(result)
	return result
}
