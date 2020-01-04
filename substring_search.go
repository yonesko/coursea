package coursea

func isCyclicRotation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := len(b) - 1; i > 0; i-- {
		if a[:1] == b[i:i+1] && a == b[i:]+b[:i] {
			return true
		}
	}
	return false
}

func tandemRepeat(a string) string {
	return ""
}
