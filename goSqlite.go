package goSqlite

import (
	"database/sql"
)
import _ "github.com/mattn/go-sqlite3"

type GoSqlite struct {
	Db       string
	Sql      string
	Database *sql.DB
	Stmt     *sql.Stmt
}

func (goSqlite *GoSqlite) Connect() *Resp {
	var resp = &Resp{}
	database, err := sql.Open("sqlite3", goSqlite.Db)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "connected"
		goSqlite.Database = database
	}
	return resp
}

func (goSqlite *GoSqlite) Exec() *Resp {
	var resp = &Resp{}
	re, err := goSqlite.Database.Exec(goSqlite.Sql)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		resp.Data = re
	}
	return resp
}

func (goSqlite *GoSqlite) Prepare() *Resp {
	var resp = &Resp{}
	stmt, err := goSqlite.Database.Prepare(goSqlite.Sql)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		goSqlite.Stmt = stmt
	}
	return resp
}

func (goSqlite *GoSqlite) Select() *Resp {
	var resp = &Resp{}
	rows, err := goSqlite.Database.Query(goSqlite.Sql)
	rows.Columns()
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		resp.Rows = rows
	}
	return resp
}

func (goSqlite *GoSqlite) Close() {
	goSqlite.Database.Close()
}
