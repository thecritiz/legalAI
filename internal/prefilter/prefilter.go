package prefilter

import "strings"

func Clean(text string) string {
	start := strings.Index(text, "1. Title and extent")
	if start == -1 {
		return text // fail-safe
	}

	text = text[start:]

	noise := []string{
		"WritingLaw.com",
		"INDIAN PENAL CODE, 1860",
		"CHAPTER",
		"•",
	}

	for _, n := range noise {
		text = strings.ReplaceAll(text, n, "")
	}

	return text
}
