package engine

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Locator struct {
	cache map[string]map[string]string
}

func NewLocator() *Locator {
	return &Locator{
		cache: make(map[string]map[string]string),
	}
}

func (l *Locator) Get(key string) (string, error) {
	parts := strings.Split(key, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid locator key: '%s'. Use 'filename.locatorname' format", key)
	}
	page, name := parts[0], parts[1]

	if _, ok := l.cache[page]; !ok {
		filePath := fmt.Sprintf("locators/%s.yml", page)
		data, err := os.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to read locator file '%s': %w", filePath, err)
		}

		var locs map[string]string
		err = yaml.Unmarshal(data, &locs)
		if err != nil {
			return "", fmt.Errorf("failed to parse yaml file '%s': %ww", filePath, err)
		}
		l.cache[page] = locs
	}

	loc, ok := l.cache[page][name]
	if !ok {
		return "", fmt.Errorf("locator '%s' not found in file '%s.yml'", key, page)
	}
	return loc, nil
}
