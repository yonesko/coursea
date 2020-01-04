package coursea

func isCyclicRotation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := len(b) - 1; i > 0; i-- {
		if a == b[i:]+b[:i] {
			return true
		}
	}
	return false
}

func tandemRepeat(a string) string {
	return ""
}
