package urlshort

import (
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attemp to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all relate to having
// invalid YAML data.
//
// See `MapHandler` to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(path string, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := parseYAML(path)
	if err != nil {
		return nil, fmt.Errorf("YAMLHandler: %w", err)
	}

	pathsToURLs := buildMap(paths)

	return MapHandler(pathsToURLs, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFuc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. IF the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//	{
//	    {"path": "/some-path", "url": "https://google.com"},
//	}
//
// The only errors that will be returned all relate to having
// invalid JSON data.
//
// See `MapHandler` to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(path string, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := parseJSON(path)
	if err != nil {
		return nil, fmt.Errorf("JSONHandler: %w", err)
	}

	pathsToURLs := buildMap(paths)

	return MapHandler(pathsToURLs, fallback), nil
}
