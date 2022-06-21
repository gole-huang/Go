package main

import (
	"fmt"
	"math"
)

func divide(dividend, divisor int) int {
	if divisor == 0 {
		return int(math.NaN())
	}
	if dividend == 0 {
		return 0
	}
	sign := dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	result := 0
	for dividend >= divisor {
		div := divisor
		res := 1
		for dividend >= div+div {
			div <<= 1
			res <<= 1
		}
		dividend -= div
		result += res

	}
	if !sign {
		result = -result
	}
	if result > math.MaxInt32 {
		return 2147483647
	} else if result < math.MinInt32 {
		return -2147483648
	} else {
		return result
	}
}

func main() {
	fmt.Println(divide(math.MaxInt32, 1))
}
