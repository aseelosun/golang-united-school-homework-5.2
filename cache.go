package cache

import "time"

type Cache struct {
	cElement       string
	cDeadlineExist bool
	cDeadline      time.Time
}

func NewCache() Cache {
	return Cache{
		cElement:       make(map[string]string),
		cDeadline:      make(map[string]time.Time),
		cDeadlineExist: make(map[string]bool),
	}
}

func (receiver) Get(key string) (string, bool) {
}

func (receiver) Put(key, value string) {

}

func (receiver) Keys() []string {
}

func (receiver) PutTill(key, value string, deadline time.Time) {
}
