package store

import (
	"context"
	"database/sql"
	"fmt"
)

type DeleteStmt struct {
	store  *ToDoItemStore
	err    error
	tx     *sql.Tx
	record *ToDoItem
}

func (stmt *DeleteStmt) Tx(tx *sql.Tx) *DeleteStmt {
	stmt.tx = tx
	return stmt
}

func (stmt *DeleteStmt) Record(o *ToDoItem) *DeleteStmt {
	stmt.record = o
	return stmt
}

func (stmt *DeleteStmt) SQL() (string, []interface{}, error) {

	if stmt.err != nil {
		return "", nil, stmt.err
	}

	sqlStr := "DELETE FROM " + stmt.store.Table + " WHERE "

	for _, f := range stmt.store.FieldList {
		sqlStr += f + " = ?"
	}

	sqlStr += ";"

	params := []interface{}{
		stmt.record.ID,
		stmt.record.Name,
		stmt.record.Description,
	}

	return sqlStr, params, nil

}

func (stmt *DeleteStmt) ExecContext(ctx context.Context) (sql.Result, error) {

	if stmt.err != nil {
		return nil, stmt.err
	}

	if stmt.record == nil {
		return nil, fmt.Errorf("record is required")
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

func (stmt *DeleteStmt) MustExecContext(ctx context.Context) sql.Result {

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		panic(err)
	}

	return res

}

func (stmt *DeleteStmt) Exec() (sql.Result, error) {

	if stmt.err != nil {
		return nil, stmt.err
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

func (stmt *DeleteStmt) MustExec() sql.Result {

	res, err := stmt.Exec()
	if err != nil {
		panic(err)
	}

	return res

}
