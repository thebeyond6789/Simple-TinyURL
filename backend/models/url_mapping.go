package models

import (
	"sync"
)

type URLMapping struct {
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

var (
	mappings = make(map[string]string)
	mutex    = &sync.RWMutex{}
)

func GetURL(slug string) (string, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	
	url, exists := mappings[slug]
	return url, exists
}

func SetURL(slug, url string) {
	mutex.Lock()
	defer mutex.Unlock()
	
	mappings[slug] = url
}