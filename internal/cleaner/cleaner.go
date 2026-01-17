package cleaner

import (
	"regexp"
	"strings"
)

func Normalize(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")
	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")
	return strings.TrimSpace(text)
}
