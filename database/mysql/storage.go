package mysql

import "errors"

// Storage is the db access implementation for MySQL
type Storage struct {
}

// GetUsers returns all the existing users
func (s *Storage) GetUsers() (users []string, err error) {
	err = errors.New("Just a dummy implemetation")
	return
}
