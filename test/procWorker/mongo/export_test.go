package mongo

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"learning-golang-process/test/procWorker/utils"
	"os"
	sync2 "sync"
	"testing"
)

var (
	wg     sync2.WaitGroup
	output = make([]string, 0)
)

func readFile(dir string) {
	f, _ := os.Open(dir)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		output = append(output, line)
	}
}

func work() {
	/*for _, v := range output {
		readMongo()
	}*/
}

func TestReadMongo(t *testing.T) {
	//work()
	/*
		LJ8E3A1M2HE006978
		LJ8E3A1M1HE006616
		LJ8E3A1M4HE006397
		LJ8E3A1M1HE006390
		LJ8E3A1M6HE999322
		LJ8E3A1MXHE006002
		LJ8E3A1M7HE999328
		LJ8E3A1M4HE006836
		LJ8E3A1M0HE006753
		LJ8E3A1M4HE008179
		LJ8E3A1M5HE004836
	*/
	// 南昌市的两台测试车辆
	/*
	 * LJ8E3A1M3HE006360
	 * LJ8E3A1M7HE006135
	 */
	vin := "LJ8E3A1M3HE006360"
	_, MG := utils.NewMgo("127.0.0.1:17017", 0)
	readMongo(MG, vin)
}

//根据上报时间获取mongoid
func getMongoId(tm int64) bson.ObjectId {
	logid := bson.NewObjectId()
	s := logid.Hex()
	b, _ := hex.DecodeString(s)
	binary.BigEndian.PutUint32(b[0:], uint32(tm))
	id2 := bson.ObjectIdHex(hex.EncodeToString(b))
	return id2
}

func readMongo(MG *mgo.Session, o string) {

	mgId := getMongoId(1535472000)
	mgLgt := getMongoId(1535681280)
	m := bson.M{
		"vin": o,
		"_id": bson.M{"$gte": mgId, "$lte": mgLgt},
	}
	var limit, skip int
	limit = 0
	skip = 0
	query := utils.Find(MG, "ReportData", limit, skip, m)
	logStr := ""
	filename := fmt.Sprintf("/home/work/go/src/learning-golang-process/test/procWorker/logs/%s.txt", o)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}
	f, _ := os.OpenFile(filename, os.O_WRONLY, 0666)
	defer f.Close()

	for _, v := range query {
		logStr = fmt.Sprintf("%s %02x %s\n", v.Vin, v.Action, hex.EncodeToString(v.RawData))
		io.WriteString(f, logStr)
	}
}

/*
 * LW4522H37014360X
 * LJ8E3A1M5GE001384
 * LW605EI32000090C
 */

// 解析原始报文
func TestParseData(t *testing.T) {
	filename := fmt.Sprintf("/home/work/go/src/learning-golang-process/test/procWorker/logs/%s.txt", "LJ8E3A1M3HE006360")
	readFile(filename)
	fmt.Println(output)
}
