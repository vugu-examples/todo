package store

import (
	"database/sql"
	"errors"
)

type SelectStmt struct {
	store         *ToDoItemStore
	selectColumns []string //"`widget_id`", "`widget_name`", "`group`"
	whereList     []whereArg
	err           error
	tx            *sql.Tx
	distinct      bool
}

func (stmt *SelectStmt) Where(arg string, value interface{}) *SelectStmt {
	if stmt.err != nil {
		return stmt
	}

	stmt.whereList = append(stmt.whereList, whereArg{
		arg:   arg,
		value: value,
	})

	return stmt

}

func (stmt *SelectStmt) Tx(tx *sql.Tx) *SelectStmt {
	if stmt.err != nil {
		return stmt
	}
	stmt.tx = tx
	return stmt
}

func (stmt *SelectStmt) SQL() (string, []interface{}, error) {
	if stmt.err != nil {
		return "", nil, stmt.err
	}

	var q string
	var wheres []interface{}

	if len(stmt.selectColumns) == 0 || (len(stmt.selectColumns) == 1 && stmt.selectColumns[0] == "*") {
		q = "SELECT * FROM " + stmt.store.Table
	} else {
		q = "SELECT "
		for i, c := range stmt.selectColumns {
			if i == (len(stmt.selectColumns) - 1) {
				q += c
			} else {
				q += c + ", "
			}
		}
		q += " FROM " + stmt.store.Table
	}

	if len(stmt.whereList) != 0 {

		for _, w := range stmt.whereList {
			q = "WHERE " + w.arg
		}

	}

	return q, wheres, nil
}

func (stmt *SelectStmt) GetOne() (*ToDoItem, error) {

	if stmt.err != nil {
		return nil, stmt.err
	}

	query, args, err := stmt.SQL()
	if err != nil {
		return nil, err
	}

	rows, err := stmt.store.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var ret ToDoItem

	i := 0
	for rows.Next() {
		if i == 1 {
			return nil, errors.New("sqlgen error: unexpected amount of rows")
		}

		for _, r := range stmt.selectColumns {
			switch r {
			case "id", "`id`":
				err = rows.Scan(&ret.ID)
			case "name", "`name`":
				err = rows.Scan(&ret.Name)
			case "description", "`description`":
				err = rows.Scan(&ret.Description)
			}
			if err != nil {
				return nil, err
			}
		}
		i++
	}

	return &ret, nil

}
