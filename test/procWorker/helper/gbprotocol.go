package helper

import "bytes"
import (
	"encoding/binary"
	"fmt"
	. "learning-golang-process/test/procWorker/model"
	"math"
)

func ParseProto(r *bytes.Buffer) *ProtoData {
	pro := &ProtoData{}
	runHandler := func(action byte) {
		switch action {
		case 0x1:
			pro.X1 = handler_0x01(r)
		case 0x2:
			pro.X2 = handler_0x02(r)
		case 0x3:
			pro.X3 = handler_0x03(r)
		case 0x4:
			pro.X4 = handler_0x04(r)
		case 0x5:
			pro.X5 = handler_0x05(r)
		case 0x6:
			pro.X6 = handler_0x06(r)
		case 0x7:
			pro.X7 = handler_0x07(r)
		case 0x8:
		case 0x9:
		}
	}
	for {
		action, err := r.ReadByte()
		if err != nil {
			break
		}
		if action == 0 || action > 9 {
			break
		}
		runHandler(action)
	}
	return pro
}

func handler_0x01(r *bytes.Buffer) *P01 {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	// 1Byte 车辆状态
	vehicle_status, _ := r.ReadByte()
	// 1Byte 充电状态
	charge_status, _ := r.ReadByte()
	// 1Byte 运行模式
	run_mode, _ := r.ReadByte()
	// 2Byte 车数
	r.Read(buf2)
	speed := binary.BigEndian.Uint16(buf2)
	// 4Byte 累计里程
	r.Read(buf4)
	miles := int64(binary.BigEndian.Uint32(buf4))
	// 2Byte 总电压
	r.Read(buf2)
	voltage := binary.BigEndian.Uint16(buf2)
	// 2Byte 总电流
	r.Read(buf2)
	current := int64(binary.BigEndian.Uint16(buf2)) - 10000 // Offest 1000A
	// 1Byte SOC
	soc, _ := r.ReadByte()
	// 1Byte DC-DC状态
	dc_dc, _ := r.ReadByte()

	// 1Byte, gears
	gearsByte, _ := r.ReadByte()
	gears := gearsByte & 0x0f           // gears 档位
	gears_braking := gearsByte >> 4 & 1 // braking 刹车
	gears_drive := gearsByte >> 5 & 1   // drive 驱动力
	// 2Byte 绝缘电阻
	r.Read(buf2)
	resistance := binary.BigEndian.Uint16(buf2)
	// 1Byte
	acc_pedal_scale, _ := r.ReadByte()
	// 1Byte
	brake_pedal_scale, _ := r.ReadByte()
	p01 := &P01{
		Vehicle_status:    vehicle_status,
		Charge_status:     charge_status,
		Run_mode:          run_mode,
		Speed:             speed,
		Miles:             miles,
		Voltage:           voltage,
		Current:           current,
		SOC:               soc,
		Dc_dc:             dc_dc,
		Gears:             gears,
		Gears_braking:     gears_braking,
		Gears_drive:       gears_drive,
		Resistance:        resistance,
		Acc_pedal_scale:   acc_pedal_scale,
		Brake_pedal_scale: brake_pedal_scale,
	}
	return p01
}

func handler_0x02(r *bytes.Buffer) *P02 {
	drive, _ := r.ReadByte()
	p2 := &P02{
		Drive: drive,
	}
	for i := 0; i < int(drive); i++ {
		p2.DriveList = append(p2.DriveList, handler_0x02_drive(r))
	}
	return p2
}

func handler_0x02_drive(r *bytes.Buffer) *P002 {
	// Declare an empty slice
	buf_2 := make([]byte, 2)

	// 1Byte, 1~253
	e1, _ := r.ReadByte()
	// 1Byte
	e2, _ := r.ReadByte()
	// 1Byte, offest 40°C
	e3, _ := r.ReadByte()

	// 2Byte, 0 ~65531,offest 20000
	e4, _ := r.Read(buf_2)

	// 2Byte, 0 ~ 65531
	e5, _ := r.Read(buf_2)

	// 1Byte, 0 ~ 250, offest 40
	e6, _ := r.ReadByte()

	// 2Byte, 0 ~ 60000
	e7, _ := r.Read(buf_2)
	// 2Byte, 0 ~ 20000,offest 10000
	e8, _ := r.Read(buf_2)
	return &P002{
		E1: e1,
		E2: e2,
		E3: e3,
		E4: uint16(e4),
		E5: uint16(e5),
		E6: int8(e6),
		E7: int16(e7),
		E8: uint16(e8),
	}
}

func handler_0x03(r *bytes.Buffer) *P03 {
	// Declare an empty slice
	buf_2 := make([]byte, 2)

	// 2Byte
	r.Read(buf_2)
	e1 := binary.BigEndian.Uint16(buf_2)
	// 2Byte
	r.Read(buf_2)
	e2 := binary.BigEndian.Uint16(buf_2)
	// 2Byte
	r.Read(buf_2)
	e3 := binary.BigEndian.Uint16(buf_2)
	// 2Byte
	r.Read(buf_2)
	fuel_probe_total := binary.BigEndian.Uint16(buf_2)

	// 1 * N
	var e5 uint16
	if fuel_probe_total != 65534 && fuel_probe_total != 65535 {
		buf_n := make([]byte, fuel_probe_total)
		r.Read(buf_n)
		e5 = binary.BigEndian.Uint16(buf_n)
	}
	// 2Byte
	r.Read(buf_2)
	e6 := binary.BigEndian.Uint16(buf_2)
	// 1Byte
	e7, _ := r.ReadByte()
	// 2Byte
	r.Read(buf_2)
	e8 := binary.BigEndian.Uint16(buf_2)
	// 1Byte
	e9, _ := r.ReadByte()
	// 2Byte
	r.Read(buf_2)
	e10 := binary.BigEndian.Uint16(buf_2)
	// 1Byte
	e11, _:= r.ReadByte()
	// 1Byte
	e12, _ := r.ReadByte()
	return &P03{
		E1: uint16(e1),
		E2: uint16(e2),
		E3: uint16(e3),
		E4: fuel_probe_total,
		E5: e5,
		E6: int16(e6),
		E7: e7,
		E8: e8,
		E9: e9,
		E10: e10,
		E11: e11,
		E12: e12,
	}
}

func handler_0x04(r *bytes.Buffer) *P04 {
	// Declare an empty slice
	buf_2 := make([]byte, 2)

	// 1Byte
	e1, _ := r.ReadByte()
	// 2Byte
	e2, _ := r.Read(buf_2)
	// 2Byte
	e3, _ := r.Read(buf_2)
	return &P04{
		E1:e1,
		E2: uint16(e2),
		E3: uint16(e3),
	}
}

func handler_0x05(r *bytes.Buffer) *P05 {

	// Declare an empty slice
	buf_4 := make([]byte, 4)

	// 1Byte
	state,_ := r.ReadByte()

	// 4Byte, longitude
	r.Read(buf_4)
	longitude := binary.BigEndian.Uint32(buf_4)
	// 4Byte, Latitude
	r.Read(buf_4)
	latitude := binary.BigEndian.Uint32(buf_4)

	logitude_str := fmt.Sprintf("%.6f", float64(longitude)/math.Pow(10, 6))
	latitude_str := fmt.Sprintf("%.6f", float64(latitude)/math.Pow(10, 6))

	return &P05{
		E1:state,
		E2: logitude_str,
		E3: latitude_str,
	}
}

func handler_0x06(r *bytes.Buffer) *P06 {

	// Declare an empty slice
	buf_2 := make([]byte, 2)

	// 1Byte
	max_system_num, _ := r.ReadByte()
	// 1Byte
	max_mon_num, _ := r.ReadByte()
	// 2Byte
	r.Read(buf_2)
	max_mon_voltage := binary.BigEndian.Uint16(buf_2)
	// 1Byte

	min_system_num, _ := r.ReadByte()
	// 1Byte
	min_mon_num, _ := r.ReadByte()
	// 2Byte
	r.Read(buf_2)
	min_mon_vltage := binary.BigEndian.Uint16(buf_2)
	// 1Byte
	max_temp_num, _ := r.ReadByte()
	// 1Byte
	max_probe_num, _ := r.ReadByte()
	// 1Byte
	max_num_byte, _ := r.ReadByte()
	max_num := int16(max_num_byte)
	// 1Byte
	min_temp_num, _ := r.ReadByte()
	// 1Byte
	min_probe_num, _ := r.ReadByte()
	// 1Byte
	min_num_byte, _ := r.ReadByte()
	min_num := int16(min_num_byte)

	//x02x06
	return &P06{
		E1:  max_system_num,
		E2:     max_mon_num,
		E3: max_mon_voltage,
		E4:  min_system_num,
		E5:     min_mon_num,
		E6: min_mon_vltage,
		E7:    max_temp_num,
		E8:   max_probe_num,
		E9:         max_num,
		E10:    min_temp_num,
		E11:   min_probe_num,
		E12:         min_num,
	}
}

func handler_0x07(r *bytes.Buffer) *P07 {

	// Declare an empty slice
	buf_4 := make([]byte, 4)

	// 1Byte
	state,_ := r.ReadByte()

	// 4Byte, longitude
	r.Read(buf_4)
	longitude := binary.BigEndian.Uint32(buf_4)
	// 4Byte, Latitude
	r.Read(buf_4)
	latitude := binary.BigEndian.Uint32(buf_4)

	logitude_str := fmt.Sprintf("%.6f", float64(longitude)/math.Pow(10, 6))
	latitude_str := fmt.Sprintf("%.6f", float64(latitude)/math.Pow(10, 6))

	return &P07{
		E1:state,
		E2: logitude_str,
		E3: latitude_str,
	}
}


