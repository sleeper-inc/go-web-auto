package engine

import (
	"encoding/json"
	"fmt"
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
		return "", fmt.Errorf("invalid locator key: %s", key)
	}
	page, name := parts[0], parts[1]

	if _, ok := l.cache[page]; !ok {
		data, err := os.ReadFile(fmt.Sprintf("locators/%s.json", page))
		if err != nil {
			return "", err
		}
		var locs map[string]string
		err = json.Unmarshal(data, &locs)
		if err != nil {
			return "", err
		}
		l.cache[page] = locs
	}

	loc, ok := l.cache[page][name]
	if !ok {
		return "", fmt.Errorf("locator not found: %s", key)
	}
	return loc, nil
}
