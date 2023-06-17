package luhn

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}

	return true
}
