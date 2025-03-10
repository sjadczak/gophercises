package story

import (
	"encoding/json"
	"fmt"
	"os"
)

var ErrParseStory = fmt.Errorf("could not parse story json")

type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ParseStory(path string) (map[string]Arc, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, ErrParseStory
	}
	defer f.Close()

	story := make(map[string]Arc)
	decoder := json.NewDecoder(f)
	decoder.Decode(&story)

	return story, nil
}
