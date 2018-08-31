package excel

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-log"
)

type ObdDevice struct {
	Id             int64  `gorm:"primary_key:id;AUTO_INCREMENT:id"`
	Dealerid       uint32 `gorm:"column:dealerid"`
	Deviceidstring string `gorm:"column:deviceidstring"`
	Ccid           string `gorm:"column:ccid"`
	Vin            string `gorm:"column:vin"`
	PinNumber      string `gorm:"column:pin_number"`
	IsPin          int32  `gorm:"column:is_pin"`
}

func InitDB() *gorm.DB {
	Gdb, err := gorm.Open("mysql", "root:123456@/car_cheler?charset=utf8&parseTime")
	if err != nil {
		log.Fatal("connect mysql database err: %v", err)
	}
	Gdb.LogMode(true)
	Gdb.SingularTable(true)
	return Gdb
}
