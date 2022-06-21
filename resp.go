package goSqlite

import "database/sql"

type Resp struct {
	Code    int
	Message string
	Data    interface{}
	Rows    *sql.Rows
}
