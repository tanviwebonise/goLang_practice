package dao

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

var db *sql.DB 

func InitDB(dataSourceName string){
	var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Panic(err)
    }
	//defer db.Close()
    err = db.Ping()
    if err != nil {
        log.Panic(err)
    }
}