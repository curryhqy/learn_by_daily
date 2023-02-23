package max

import "github.com/curryhqy/learn_by_daily/algorithm/constraints"

func Bitwise(a int, b int, base int) int {
	z := a - b
	i := (z >> base) & 1
	return a - (i * z)
}

func Int[T constraints.Integer](values ...T) T {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
