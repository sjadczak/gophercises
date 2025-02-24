package urlshort

import (
	"encoding/json"
	"os"
)

func parseJSON(json_file string) ([]pathURL, error) {
	f, err := os.Open(json_file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ps := make([]pathURL, 0)
	decoder := json.NewDecoder(f)
	decoder.Decode(&ps)

	return ps, nil
}
