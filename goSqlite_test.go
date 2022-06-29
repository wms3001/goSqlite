package goSqlite

import (
	"encoding/json"
	"fmt"
	"github.com/wms3001/goCommon"
	"testing"
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
	var resp = &goCommon.Resp{}
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
		lastId, _ := re.LastInsertId()
		roweffct, _ := re.RowsAffected()
		data := make(map[string]int64)
		data["lastInsertId"] = lastId
		data["rowsAffected"] = roweffct
		dataType, _ := json.Marshal(data)
		dataString := string(dataType)
		resp.Data = dataString
	}
	fmt.Println(resp)
	//s1 := []interface{}{"guoke", "2012-12-09"}
	//resp := goSqlite.Insert("guoke", "2012-12-09")

}

func TestGoSqlite_Delete(t *testing.T) {
	var resp = &goCommon.Resp{}
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
		lastId, _ := re.LastInsertId()
		roweffct, _ := re.RowsAffected()
		data := make(map[string]int64)
		data["lastInsertId"] = lastId
		data["rowsAffected"] = roweffct
		dataType, _ := json.Marshal(data)
		dataString := string(dataType)
		resp.Data = dataString
	}
	fmt.Println(resp)
}

func TestGoSqlite_Update(t *testing.T) {
	var resp = &goCommon.Resp{}
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
		lastId, _ := re.LastInsertId()
		roweffct, _ := re.RowsAffected()
		data := make(map[string]int64)
		data["lastInsertId"] = lastId
		data["rowsAffected"] = roweffct
		dataType, _ := json.Marshal(data)
		dataString := string(dataType)
		resp.Data = dataString
	}
	fmt.Println(resp)
}

type Tttttt struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

func TestGoSqlite_Select(t *testing.T) {
	goSqlite := &GoSqlite{}
	goSqlite.Db = "D:\\db\\test.db"
	goSqlite.Connect()
	defer goSqlite.Close()
	goSqlite.Sql = `SELECT * FROM user`
	resp := goSqlite.Select()
	var ttt []Tttttt
	fmt.Println(resp.Data)
	json.Unmarshal([]byte(resp.Data), &ttt)
	fmt.Println(ttt)
}
