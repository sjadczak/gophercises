package urlshort

import (
	"os"

	"gopkg.in/yaml.v3"
)

func parseYAML(path string) ([]pathURL, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ps := make([]pathURL, 0)
	decoder := yaml.NewDecoder(f)
	decoder.Decode(&ps)

	return ps, nil
}
