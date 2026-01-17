package parser

import "regexp"

var SectionHeader = regexp.MustCompile(
	`(\d+)\.\s+([^—–-]+)[—–-]`,
)
