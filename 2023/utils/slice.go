package utils

import (
	"fmt"
	"math"
)

func Min(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("the slice is empty")
	}

	min := math.MaxInt

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min, nil
}

func Max(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("the slice is empty")
	}

	max := -math.MaxInt

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max, nil
}
