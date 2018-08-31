package model

import (
	"fmt"
	"learning-golang-process/test/procWorker/utils"
)

const (
	DB       = "onlineDB"
	QUERYSQL = `SELECT id, vin FROM %s.car_obd_device WHERE dealerid=275380 LIMIT 0, 6`
	COUNTSQL = `SELECT COUNT(0) AS totalCount FROM %s.car_obd_device WHERE dealerid=275380`
)

type CarObdDevice struct {
	Id             int64  `xorm:"pk autoincr"`
	Deviceidstring string `xorm:"char(16)"`
	Vin            string `xorm:"varchar(50)"`
	Gps            string `xorm:"varchar(32)"`
	Address        string `xorm:"varchar(256)"`
	Rectime        string `xorm:"varchar(32)"`
}

type GeoDevice struct {
	Id             int64  `xorm:"pk autoincr"`
	Deviceidstring string `xorm:"char(16)"`
	Vin            string `xorm:"varchar(50)"`
}

type RowCountDef struct {
	TotalCount uint32 `xorm:"totalCount"`
}

func Gets() ([]*CarObdDevice, int64, error) {
	sun := utils.NewDB()
	sun.ShowSQL()
	obdDevice := CarObdDevice{}
	res := make([]*CarObdDevice, 0)
	count, err := sun.Table("car_obd_device").Where("dealerid= 275380 AND is_pin=? AND vin!=? AND vin != ?", 2, "", "00000000000000000").FindAndCount(&res, obdDevice)
	return res, count, err
}

func (de *CarObdDevice) InsertOne() {
	sun := utils.NewDB()
	sun.ShowSQL()
	sun.Table("export_device_gps").InsertOne(de)
}

func (de *CarObdDevice) Del() {
	sun := utils.NewDB()
	sun.ShowSQL()
	id, err := sun.Table("car_obd_device").Delete(de)
	fmt.Println(id, err)
}

func GetsGeo() ([]*GeoDevice, error) {
	sun := utils.NewDB()
	sun.ShowSQL()
	obdDevice := GeoDevice{}
	res := make([]*GeoDevice, 0)
	err := sun.Table("geo_device").Find(&res, obdDevice)
	return res, err
}
