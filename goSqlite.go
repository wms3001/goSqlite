package goSqlite

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/goTool"
)

type GoSqlite struct {
	Db       string
	Sql      string
	Database *sql.DB
	Stmt     *sql.Stmt
}

func (goSqlite *GoSqlite) Connect() *goCommon.Resp {
	var resp = &goCommon.Resp{}
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

func (goSqlite *GoSqlite) Exec() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	re, err := goSqlite.Database.Exec(goSqlite.Sql)
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
	return resp
}

func (goSqlite *GoSqlite) Prepare() *goCommon.Resp {
	var resp = &goCommon.Resp{}
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

func (goSqlite *GoSqlite) Select() *goCommon.Resp {
	var goTool = goTool.GoTool{}
	var resp = &goCommon.Resp{}
	rows, err := goSqlite.Database.Query(goSqlite.Sql)
	rows.Columns()
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		columns, _ := rows.Columns()
		count := len(columns)
		values := make([]interface{}, count)
		scanArgs := make([]interface{}, count)
		for i := range values {
			scanArgs[i] = &values[i]
		}
		var ttt []map[string]interface{}
		for rows.Next() {
			err := rows.Scan(scanArgs...)
			if err != nil {
				resp.Code = -1
				resp.Message = err.Error()
				break
			}
			mm := make(map[string]interface{})
			for i, v := range values {
				nn := goTool.Strval(v)
				mm[columns[i]] = nn
				//x := nn.([]byte)
				//if nx, ok := strconv.ParseFloat(string(x), 64); ok == nil {
				//	//masterData[columns[i]] = append(masterData[columns[i]], nx)
				//	mm[columns[i]] = nx
				//} else if b, ok := strconv.ParseInt(string(x), 10, 64); ok == nil {
				//	mm[columns[i]] = b
				//} else if b, ok := strconv.ParseBool(string(x)); ok == nil {
				//	//masterData[columns[i]] = append(masterData[columns[i]], b)
				//	mm[columns[i]] = b
				//} else if "string" == fmt.Sprintf("%T", string(x)) {
				//	//masterData[columns[i]] = append(masterData[columns[i]], string(x))
				//	mm[columns[i]] = string(x)
				//	//} else if reflect.TypeOf(x) {
				//} else {
				//	fmt.Printf("Failed on if for type %T of %v\n", x, x)
				//}
			}
			ttt = append(ttt, mm)
		}
		dataType, _ := json.Marshal(ttt)
		dataString := string(dataType)
		resp.Code = 1
		resp.Message = "success"
		resp.Data = dataString
	}
	return resp
}

func (goSqlite *GoSqlite) Close() {
	goSqlite.Database.Close()
}
