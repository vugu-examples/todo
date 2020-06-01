package todo

import (
	"database/sql"

	"github.com/julienschmidt/httprouter"
)

type DBConnStr string

func NewDBConnStr() DBConnStr { return "root:root@/db" }
func (dbC *DBConnStr) String() string {
	return string(*dbC)
}

type DBDriverName string

func (dbD *DBDriverName) String() string {
	return string(*dbD)
}
func NewDBDriverName() DBDriverName { return "mysql" }

func NewDBConn(c DBConnStr, d DBDriverName) *sql.DB {
	db, err := sql.Open(c.String(), d.String())
	if err != nil {
		panic(err)
	}
	return db
}

func NewToDoApp(c DBConnStr, d DBDriverName, router *httprouter.Router) (*App, error) {
	db, err := sql.Open(c.String(), d.String())
	if err != nil {
		return nil, err
	}
	return &App{
		DB:     db,
		Router: router,
	}, nil

}

type App struct {
	DB     *sql.DB
	Router *httprouter.Router
}

func main() {

	app, err := Setup()
	if err != nil {
		panic(err)
	}

	_ = app

}
