package karatsuba

import (
	"math"
)

func Multiply(x, y int64) int64 {
	if x < 10 || y < 10 {
		return x * y
	}

	totalNumberOfDigits := totalNumberOfDigits(x)

	a, b := split(x, totalNumberOfDigits)
	c, d := split(y, totalNumberOfDigits)

	p := a + b
	q := c + d

	ac := Multiply(a, c)
	bd := Multiply(b, d)
	pq := Multiply(p, q)

	adbc := pq - ac - bd

	return int64(math.Pow(10, float64(totalNumberOfDigits)))*ac + int64(math.Pow(10, math.Ceil(float64(totalNumberOfDigits)/2)))*adbc + bd
}

func totalNumberOfDigits(number int64) int64 {
	if number == 0 {
		return 1
	}

	var digits int64

	for number > 0 {
		digits++
		number /= 10
	}

	return digits
}

func split(number, totalNumberOfDigits int64) (int64, int64) {
	exponent := math.Ceil(float64(totalNumberOfDigits) / 2)
	divisor := int64(math.Pow(10, exponent))

	if number < divisor {
		return 0, number
	}

	return number / divisor, number % divisor
}
