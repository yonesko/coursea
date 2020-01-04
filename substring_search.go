package coursea

import "strings"

func isCyclicRotation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	index := strings.LastIndex(b, a[:1])
	if index < 0 {
		return false
	}

	p2 := b[index:]
	p1 := b[:index]

	return a == p2+p1

}
