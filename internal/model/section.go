package model

type Section struct {
	Act           string `json:"act"`
	ActShort      string `json:"act_short"`
	SectionNumber string `json:"section_number"`
	SectionTitle  string `json:"section_title"`
	Text          string `json:"text"`
	Source        string `json:"source"`
	Language      string `json:"language"`
}
