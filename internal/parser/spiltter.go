package parser

func SplitSections(text string) []string {
	matches := SectionHeader.FindAllStringIndex(text, -1)
	var sections []string

	for i, m := range matches {
		start := m[0]
		end := len(text)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		sections = append(sections, text[start:end])
	}

	return sections
}
