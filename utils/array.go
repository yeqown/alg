// 数组常用操作

package utils

// EqualArray ...
func EqualArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for pos, val := range arr1 {
		if arr2[pos] != val {
			return false
		}
	}
	return true
}

// MaxOfArray ...
func MaxOfArray(arr []int) (max int) {
	max = arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// MidOfArrary ...
// TODO:
func MidOfArrary(arr []int) (mid int) {
	return arr[0]
}

// MaxOfThree ...
func MaxOfThree(a, b, c int) (max int) {
	max = a
	if max < b {
		max = b
	}
	if max < c {
		max = c
	}
	return
}

// MinOfThree ...
func MinOfThree(a, b, c int) (min int) {
	min = a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return
}

// MidOfThree ...
func MidOfThree(a, b, c int) int {
	max := MaxOfThree(a, b, c)
	min := MinOfThree(a, b, c)
	return a + b + c - max - min
}
