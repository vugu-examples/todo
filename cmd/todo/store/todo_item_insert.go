package store

import (
	"context"
	"database/sql"
	"fmt"
)

type InsertStmt struct {
	store         *ToDoItemStore
	insertColumns []string      //"`id`", "`name`", "`description`"
	insertValues  []interface{} //""
	err           error
	tx            *sql.Tx
}

func (stmt *InsertStmt) Tx(tx *sql.Tx) *InsertStmt {
	stmt.tx = tx
	return stmt
}

func (stmt *InsertStmt) Columns(columns ...string) *InsertStmt {

	if stmt.err != nil {
		return stmt
	}

	if columns == nil {
		stmt.err = fmt.Errorf("columns cannot be nil")
		return stmt
	}

	for _, c := range columns {
		stmt.insertColumns = append(stmt.insertColumns, "`"+c+"`")
	}

	return stmt
}

func (stmt *InsertStmt) Values(values ...interface{}) *InsertStmt {

	if stmt.err != nil {
		return stmt
	}

	if values == nil {
		stmt.err = fmt.Errorf("values cannot be nil")
		return stmt
	}

	for _, v := range values {
		stmt.insertValues = append(stmt.insertValues, v)
	}

	return stmt
}

func (stmt *InsertStmt) Object(o *ToDoItem) *InsertStmt {

	if stmt.err != nil {
		return stmt
	}

	if len(stmt.insertValues) != 0 {
		stmt.err = fmt.Errorf("cannot use object and values in the same insert statement")
		return stmt
	}

	//if .insertColumns are nil then use default cols
	if stmt.insertColumns == nil {
		for _, f := range stmt.store.FieldList {
			stmt.insertColumns = append(stmt.insertColumns, escapeFields(f))
		}
	}

	//match the insert columns to the values of the object
	//which has to be done in the correct sequence
	for _, col := range stmt.insertColumns {

		if stmt.err != nil {
			return stmt
		}

		switch col {
		case "id", "`id`":
			stmt.insertValues = append(stmt.insertValues, &o.ID)
		case "name", "`name`":
			stmt.insertValues = append(stmt.insertValues, &o.Name)
		case "description", "`description`":
			stmt.insertValues = append(stmt.insertValues, &o.Description)
		default:
			stmt.err = fmt.Errorf("unkown column %v in insert statement", col)
		}

	}

	return stmt
}

func (stmt *InsertStmt) SQL() (string, []interface{}, error) {

	if stmt.err != nil {
		return "", nil, stmt.err
	}

	sqlStr := "INSERT INTO " + stmt.store.Table + "("

	for i, c := range stmt.insertColumns {
		sqlStr += c

		if i != len(stmt.insertColumns)-1 {
			sqlStr += ", "
		}
	}

	sqlStr += ") VALUES ("

	for i := range stmt.insertValues {
		sqlStr += "?"
		if i != len(stmt.insertValues)-1 {
			sqlStr += ", "
		}
	}

	sqlStr += ")"

	return sqlStr, stmt.insertValues, nil

}

func (stmt *InsertStmt) Exec() (sql.Result, error) {

	if stmt.err != nil {
		return nil, stmt.err
	}

	if len(stmt.insertColumns) != len(stmt.insertValues) {
		return nil, fmt.Errorf("column names and values are not of same length")
	}

	//TODO:	 any other conditional checking or whatnot

	sqlStr, insertArgs, err := stmt.SQL()
	if err != nil {
		return nil, err
	}

	switch stmt.tx == nil {
	case false:
		//use the provided transaction
		return stmt.tx.Exec(sqlStr, insertArgs...)

	default:
		//call it without a transaction
		return stmt.store.DB.Exec(sqlStr, insertArgs...)
	}

}

func (stmt *InsertStmt) ExecContext(ctx context.Context) (sql.Result, error) {

	if stmt.err != nil {
		return nil, stmt.err
	}

	if len(stmt.insertColumns) != len(stmt.insertValues) {
		return nil, fmt.Errorf("column names and values are not of same length")
	}

	//TODO:	 any other conditional checking or whatnot

	sqlStr, insertArgs, err := stmt.SQL()
	if err != nil {
		return nil, err
	}

	switch stmt.tx == nil {
	case false:
		//use the provided transaction
		return stmt.tx.ExecContext(ctx, sqlStr, insertArgs...)

	default:
		//call it without a transaction
		return stmt.store.DB.ExecContext(ctx, sqlStr, insertArgs...)
	}

}

func (stmt *InsertStmt) MustExecContext(ctx context.Context) sql.Result {

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		panic(err)
	}

	return res

}

func (stmt *InsertStmt) MustExec() sql.Result {

	res, err := stmt.Exec()
	if err != nil {
		panic(err)
	}

	return res

}
