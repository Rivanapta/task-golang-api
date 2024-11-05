package utils

import "math"

func MagicSum(n int) int {
	return n + n
}

func MagicPow(n int) int {
	return int(math.Pow(float64(n), float64(n)))
}

func MagicOdd(n int) bool {
	return n%2 != 0
}

func MagicGrade(n int) string {
	switch n {
	case 0:
		return "zero"
	case 1:
		return "bad"
	case 2:
		return "Ok"
	case 3:
		return "Nice"
	case 4:
		return "Awesome"
	case 5:
		return "Exceptional"
	default:
		return "unknown"

	}
}

func MagicName(n int) []string {
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = "Rivan"
	}
	return names
}

func MagicTria(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func MagicChange(n *int) {
	*n = *n * 2
}

// Struct untuk MagicNumber
type MagicNumber struct {
	Number int
}

// Method untuk MagicNumber
func (m *MagicNumber) Multiply(n int) {
	m.Number *= n
}
