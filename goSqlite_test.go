package goSqlite

import (
	"fmt"
	"testing"
	"time"
)

func TestGoSqlite_Connect(t *testing.T) {
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	resp := goSqlite.Connect()
	defer goSqlite.Close()
	fmt.Println(resp)
}

func TestGoSqlite_Table(t *testing.T) {
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	goSqlite.Connect()
	defer goSqlite.Close()
	goSqlite.Sql = `
    CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(128) NULL,
        created DATE NULL
    );
    `
	goSqlite.Exec()
	goSqlite.Close()
}

func TestGoSqlite_Exec(t *testing.T) {
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	goSqlite.Connect()
	defer goSqlite.Close()
	goSqlite.Sql = "INSERT INTO user(name,  created) values(\"guoke\",\"2012-12-09\")"
	resp := goSqlite.Exec()
	fmt.Println(resp)
}

func TestGoSqlite_Insert(t *testing.T) {
	var resp = &Resp{}
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	goSqlite.Connect()
	defer goSqlite.Close()
	goSqlite.Sql = `INSERT INTO user(name,  created) values(?,?)`
	goSqlite.Prepare()
	re, err := goSqlite.Stmt.Exec("test", "2021-02-25")
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		resp.Data, _ = re.LastInsertId()
	}
	fmt.Println(resp)
	//s1 := []interface{}{"guoke", "2012-12-09"}
	//resp := goSqlite.Insert("guoke", "2012-12-09")

}

func TestGoSqlite_Delete(t *testing.T) {
	var resp = &Resp{}
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	goSqlite.Connect()
	defer goSqlite.Close()
	goSqlite.Sql = `delete from user where id=?`
	goSqlite.Prepare()
	re, err := goSqlite.Stmt.Exec(3)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		resp.Data = re.RowsAffected
	}
	fmt.Println(resp)
}

func TestGoSqlite_Update(t *testing.T) {
	var resp = &Resp{}
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	goSqlite.Connect()
	defer goSqlite.Close()
	goSqlite.Sql = `update user set name=? where id=?`
	goSqlite.Prepare()
	re, err := goSqlite.Stmt.Exec("test", 10)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		resp.Data = re.RowsAffected
	}
	fmt.Println(resp)
}

func TestGoSqlite_Select(t *testing.T) {
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	goSqlite.Connect()
	defer goSqlite.Close()
	goSqlite.Sql = `SELECT * FROM user`
	resp := goSqlite.Select()
	for resp.Rows.Next() {
		var uid int
		var name string
		var created time.Time
		err := resp.Rows.Scan(&uid, &name, &created)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(uid)
		fmt.Println(name)
		fmt.Println(created)
	}
}
