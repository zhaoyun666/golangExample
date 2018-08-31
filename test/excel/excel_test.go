package excel

import (
	"github.com/jinzhu/gorm"
	"testing"
	"github.com/micro/go-log"
)

var GDB *gorm.DB

func TestStart(t *testing.T) {
	// 初始化数据库
	GDB = InitDB()
	ProData()
}

func db() *gorm.DB {
	return GDB.Table("car_obd_device")
}

func ProData() {
	obd := make([]*ObdDevice, 0)
	cond := &ObdDevice{Dealerid: 275380, IsPin: 2}
	if err := db().Offset(0).Limit(1000).Find(&obd, cond).Error; err != nil {
		log.Fatalf("query car_obd_device table err:%v", err)
	}
	Csv(obd)
}

// 命令标识	应答标识	车辆状态	充电状态	运行模式	车速	累计里程	总电压	总电流	SOC	DC-DC状态	挡位	绝缘电阻	加速踏板行程值	制动踏板状态	驱动电机个数	驱动电机序号	驱动电机状态	驱动电机控制器温度	驱动电机转速	驱动电机转矩	驱动电机温度	电机控制器输入电压	电机控制器直流母线电流	驱动电机序号	驱动电机状态	驱动电机控制器温度	驱动电机转速	驱动电机转矩	驱动电机温度	电机控制器输入电压	电机控制器直流母线电流	发动机状态	曲轴转速	燃料消耗率	定位状态	经度	纬度	最高电压电池子系统号	最高电压电池单体代号	电池单体电压最高值	最低电压电池子系统号	最低电压电池单体代号	电池单体电压最低值	最高温度子系统号	最高温度探针单体代号	最高温度值	最低温度子系统号	最低温度探针子系统代号	最低温度值	最高报警等级	通用报警标志	报警类型	可充电储能装置故障总数	N1	驱动电机故障总数	N2	发动机故障总数	N3	其他故障总数	N4	可充电储能子系统个数	可充电储能子系统号	可充电储能装置电压	可充电储能装置电流	单体电池总数	本帧起始电池序号	本帧单体电池总数	单体电池电压	可充电储能子系统个数	可充电储能子系统号	可充电储能温度探针个数	温度值

