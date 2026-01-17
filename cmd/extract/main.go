package main

import (
	"encoding/json"
	"os"

	"extractor/internal/cleaner"
	"extractor/internal/parser"
	"extractor/internal/prefilter"
)

func main() {
	raw, _ := os.ReadFile("data/ipc_raw.txt")

	text := prefilter.Clean(string(raw))
	text = cleaner.Normalize(text)

	chunks := parser.SplitSections(text)

	var sections []interface{}
	for _, c := range chunks {
		sections = append(sections, parser.ParseSection(c))
	}

	out, _ := json.MarshalIndent(sections, "", "  ")
	os.WriteFile("output/ipc_sections.json", out, 0644)
}
