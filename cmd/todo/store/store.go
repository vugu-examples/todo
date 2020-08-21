package store

import "database/sql"

type Store struct {
	DB *sql.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) ToDoItem() *ToDoItemStore {
	return newToDoItemStore(s.DB)
}
