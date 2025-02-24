package urlshort

type pathURL struct {
	Path string `yaml:"path" json:"path" sql:"path"`
	URL  string `yaml:"url" json:"url" sql:"url"`
}

func buildMap(pathURLs []pathURL) map[string]string {
	pathsToURLs := make(map[string]string)
	for _, pu := range pathURLs {
		pathsToURLs[pu.Path] = pu.URL
	}

	return pathsToURLs
}
