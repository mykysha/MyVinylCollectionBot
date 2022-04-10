package stringanalyser

func IsOneOf(goal string, options []string) bool {
	for _, option := range options {
		if goal == option {
			return true
		}
	}

	return false
}
