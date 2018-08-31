package service

import (
	"bytes"
	"fmt"
	"learning-golang-process/test/procWorker/helper"
	"learning-golang-process/test/procWorker/model"
	"learning-golang-process/test/procWorker/utils"
	"sync"
)

type TaskSlice struct {
	BatchId int
	Data    []*model.CarObdDevice
}

func ProDeviceGPS() {
	utils.New()
	res, count, err := model.Gets()
	if err != nil {
		fmt.Errorf("Query SQL Error: %s", err.Error())
		return
	}

	if count > 0 {
		result := make([]*model.CarObdDevice, 0)
		// 剔除无效车架号
		for _, v := range res {
			if v.Vin != fmt.Sprintf("%s0", v.Deviceidstring) {
				result = append(result, v)
			} else {
				count = count - 1
			}
		}

		var (
			nParseRoutine, n int
			wg               = &sync.WaitGroup{}
		)

		nParseRoutine = 100
		if nParseRoutine > len(result) {
			nParseRoutine = len(result)
			n = 1
		} else {
			n = len(result)/nParseRoutine + 1
		}

		length := len(result)
		resultCh := make(chan TaskSlice, nParseRoutine)
		for i := 0; i < nParseRoutine; i++ {
			if i*n > length {
				break
			}
			goParseRoutine(wg, result[i*n:min(int(count), (i+1)*n)], i, resultCh)
		}
		wg.Wait()
		close(resultCh)
		result = mergeResult(resultCh)

		for _, v := range result {
			v.InsertOne()
		}
	}
}

func goParseRoutine(wg *sync.WaitGroup, data []*model.CarObdDevice, n int, out chan<- TaskSlice) {
	wg.Add(1)
	go func() {
		for _, car := range data {
			//转换gps坐标
			l, addr, time := convertGPS(car.Vin)
			car.Gps = l
			car.Address = addr
			car.Rectime = time
		}

		out <- TaskSlice{BatchId: n, Data: data}

		wg.Done()
	}()
}

func convertGPS(vin string) (location, address, rectime string) {
	key := "gbzmplant"

	res, err := utils.RC.HGet(key, vin).Bytes()
	if err != nil {
		return "", "", ""
	}
	r := bytes.NewBuffer(res)
	buf6 := make([]byte, 6)
	r.Read(buf6)
	rectime = fmt.Sprintf("20%02d-%02d-%02d %02d:%02d:%02d", buf6[0], buf6[1], buf6[2], buf6[3], buf6[4], buf6[5])
	addr := helper.ParseProto(r)

	p := &helper.GpsPosition{
		Position: addr,
	}
	return addr, helper.ConvertAndResolve(p), rectime
}

// 合并结果
func mergeResult(ch <-chan TaskSlice) []*model.CarObdDevice {
	ret := make([]*model.CarObdDevice, 0)
	// TODO: odo排序
	for data := range ch {
		ret = append(ret, data.Data...)
	}

	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// LSA121BL5J2001080
// LSA121BL6J2001086
