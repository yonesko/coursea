package coursea

func isCyclicRotation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := len(b) - 1; i > 0; i-- {
		p2 := b[i:]
		p1 := b[:i]
		if a == p2+p1 {
			return true
		}
	}
	return false
}

func tandemRepeat(a string) string {
	return ""
}
