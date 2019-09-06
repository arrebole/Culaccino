package session

import "time"

var _instance = &Manager{
	storage: make(map[string]*Session),
	maxAge:  time.Hour * 3,
}

// NewStore ...
func NewStore() Store {
	return _instance
}
