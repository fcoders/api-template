package database

// Storage is the generic interface with db access functions
type Storage interface {
	GetUsers() ([]string, error) // Dummy function
}
