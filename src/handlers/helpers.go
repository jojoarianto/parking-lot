package handlers

// DeleteLastChar is method to delete last character with key ", "
func DeleteLastChar(s string) string {
	sz := len(s)
	if sz > 0 && s[sz-1] == ' ' {
		s = s[:sz-1]
	}

	sz = len(s)
	if sz > 0 && s[sz-1] == ',' {
		s = s[:sz-1]
	}

	return s
}
