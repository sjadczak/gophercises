package story

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sjadczak/gophercises/cyoa/views"
)

var ErrParseStory = fmt.Errorf("could not parse story json")

type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"chapter"`
}

type StoryService struct {
	story map[string]Arc
	Tpl   *views.HTMLTemplate
}

func NewStoryService(path string) (*StoryService, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, ErrParseStory
	}
	defer f.Close()

	service := &StoryService{}
	service.story = make(map[string]Arc)
	decoder := json.NewDecoder(f)
	decoder.Decode(&service.story)

	return service, nil
}

func (s *StoryService) GetArc(arc string) (Arc, error) {
	arcData, ok := s.story[arc]
	if !ok {
		return Arc{}, fmt.Errorf("arc not found")
	}
	return arcData, nil
}

func (s *StoryService) WebHandler(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Path[len("/story/"):]
	if a == "" {
		a = "intro"
	}

	arc, ok := s.story[a]
	if !ok {
		http.Error(w, "Story not found", http.StatusNotFound)
		return
	}

	s.Tpl.Execute(w, r, arc)
}
