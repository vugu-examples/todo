package todo_item_store

import "database/sql"

type SelectStmt struct {
	store         *ToDoItemStore
	selectColumns []string //"`widget_id`", "`widget_name`", "`group`"
	whereList     []whereArg
	err           error
	tx            *sql.Tx
	distinct      bool
}