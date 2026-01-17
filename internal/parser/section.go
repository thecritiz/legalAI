package parser

import (
	"extractor/internal/model"
	"strings"
)

func ParseSection(chunk string) model.Section {
	m := SectionHeader.FindStringSubmatch(chunk)

	header := m[0]
	number := m[1]
	title := strings.TrimSpace(m[2])
	body := strings.TrimSpace(chunk[len(header):])

	return model.Section{
		Act:           "Indian Penal Code",
		ActShort:      "IPC",
		SectionNumber: number,
		SectionTitle:  title,
		Text:          body,
		Source:        "WritingLaw.com",
		Language:      "en",
	}
}
