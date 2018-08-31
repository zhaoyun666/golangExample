package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var RC *redis.Client
var GEORC *redis.Client

func New() {
	RC = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "Carlt123",
		DB:       0,
	})
	pong, err := RC.Ping().Result()
	fmt.Println("redis", pong, err)
}

func NewGEO() {
	GEORC = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "Carlt123",
		DB:       2,
	})
	pong, err := GEORC.Ping().Result()
	fmt.Println("redis", pong, err)
}

func NewDB() *xorm.Engine {
	//masterDns := "cheler6frme2yza:rds4MnS6YOEhou@(127.0.0.1:3306)/car_cheler?charset=utf8"
	//devDns := "ruser:Passwd.123@tcp(121.40.133.157:3306)/17car_clw2?charset=utf8"
	dns := "root:123456@tcp(127.0.0.1:3306)/onlineDB?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dns)
	if err != nil {
		fmt.Errorf("Mysql Init failed:%v", err)
		return nil
	}
	return engine
}

func testOne() {
	s := 9.0

	fmt.Println(checkType(s))

}

func checkType(v interface{}) string {
	switch t := v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case []interface{}:
	default:
		_ = t
	}
	return "unkown"
}
