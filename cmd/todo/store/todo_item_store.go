package store

import (
	"database/sql"
	"fmt"
)

type ToDoItemStore struct {
	DB        *sql.DB
	Table     string
	FieldList []string
}

func NewToDoItemStore(db *sql.DB) *ToDoItemStore {
	return &ToDoItemStore{
		DB:        db,
		Table:     "todo_items",
		FieldList: []string{"id", "name", "description"},
	}
}

func (s *ToDoItemStore) Insert() *InsertStmt {
	return &InsertStmt{
		store: s,
	}
}

func (s *ToDoItemStore) Select() *SelectStmt {
	return &SelectStmt{
		store: s,
	}
}

func (s *ToDoItemStore) Delete() *DeleteStmt {
	return &DeleteStmt{
		store: s,
	}
}

func (s *ToDoItemStore) Update() *UpdateStmt {
	return &UpdateStmt{
		store: s,
	}
}

type whereArg struct {
	arg   string
	value interface{}
}

type ToDoItem struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func escapeFields(field string) string {
	return fmt.Sprintf("`%v`", field)
}
