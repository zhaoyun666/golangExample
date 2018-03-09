package db

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)
var MysqlEngine *xorm.Engine
func init(){
    dbEngine, _ := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
    MysqlEngine = dbEngine
}
