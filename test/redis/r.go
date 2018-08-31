package main

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func New() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "Carlt123",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println("redis", pong, err)
	return client
}

func NewDB() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "cheler6frme2yza:rds4MnS6YOEhou@127.0.0.1:13306/?db=car_cheler&charset=utf8")
	if err != nil {
		fmt.Errorf("Mysql Init failed:%v", err)
		return nil
	}
	return engine
}

func main() {
	//testOne()
	key := "gbzmplant"
	vin := "LJ866666366666666"
	c := New()
	h := c.HGet(key, vin)
	fmt.Println(h.Bytes())

	sun := NewDB()
	sun.Table("car_obd_device").Where("vin=?", vin).Find()
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
