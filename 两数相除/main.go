package main

import "fmt"

func divide(dividend int, divisor int) int {
	IsNega := false
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		IsNega = true
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	count := 0
	for dividend >= divisor {
		count++
		dividend -= divisor
	}

	if IsNega {
		if -count >= 2147483648 || -count < -2147483648 {
			return 2147483647
		} else {
			return -count
		}
	} else {
		if count >= 2147483648 || count < -2147483648 {
			return 2147483647
		} else {
			return count
		}
	}
}

func main() {
	fmt.Println(divide(-2147483648, -1))
}
