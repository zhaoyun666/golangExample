package utils

import (
	"time"
	"gopkg.in/mgo.v2"
)

const (
	gbinfokey = "GBDeviceInfo"
)

type ReportData struct {
	Vin     string `bson:"vin"`
	Action  byte `bson:"action"`
	RawData []byte `bson:"rawData"`
}

func NewMgo(addr string, timeout int64) (err error, session *mgo.Session) {
	//TODO
	session, err = mgo.DialWithTimeout(addr, time.Duration(timeout) * time.Second)

	if err != nil {
		return err, nil
	}
	session.SetSyncTimeout(time.Duration(timeout) * time.Second)
	session.SetSocketTimeout(time.Duration(timeout) * time.Second)

	return nil, session
}

// helper func
func Find(db *mgo.Session, col string, lit, skip int, query interface{}) []*ReportData {
	r := make([]*ReportData, 0)
	if lit > 0 {
		db.DB(gbinfokey).C(col).Find(query).Limit(lit).Skip(skip).All(&r)
	}else{
		db.DB(gbinfokey).C(col).Find(query).All(&r)
	}
	return r
}