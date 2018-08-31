package model

// 整车数据 0x1
type P01 struct {
	Vehicle_status    byte   `xorm:"tinyint(3) "`
	Charge_status     byte   `xorm:"tinyint(3) "`
	Speed             uint16 `xorm:"smallint(4) "`
	Miles             int64  `xorm:"int(10) "`
	Voltage           uint16 `xorm:"smallint(5) "`
	Current           int64  `xorm:"int(5)"`
	SOC               byte   `xorm:"tinyint(3)  'SOC'"`
	Dc_dc             byte   `xorm:"tinyint(3) "`
	Gears             byte   `xorm:"tinyint(2) "`
	Gears_braking     byte   `xorm:"tinyint(1) "`
	Gears_drive       byte   `xorm:"tinyint(1) "`
	Resistance        uint16 `xorm:"smallint(5) "`
	Run_mode          byte   `xorm:"tinyint(3) "`
	Acc_pedal_scale   byte   `xorm:"tinyint(3) "`
	Brake_pedal_scale byte   `xorm:"tinyint(3) "`
}

// 驱动电机数据2
type P02 struct {
	E1 byte // 驱动点击序号
	E2 byte // 驱动点击状态
	E3 byte // 驱动电机控制器温度
	E4 uint16 // 驱动电机转速
	E5 int16 // 驱动电机转矩
	E6 int8 // 驱动电机温度
	E7 uint16 // 电机控制器输入电压 单位0.1
	E8 int16 // 电机控制器直流母线电流
}

// 燃料电池数据
type P03 struct {
	E1 uint16 // 燃料电池电压 0.1
	E2 uint16 // 燃料电池电流 0.1
	E3 uint16 // 燃料消耗率 0.01
	E4 uint16 // 燃料电池温度探针总数
	E5 int16 // 探针温度值
	E6 int16 // 氢系统中最高温度 0.1
	E7 byte // 氢系统中最高温度探针代号
	E8 uint16 // 氢气最高浓度
	E9 byte // 氢气最高浓度传感器代号
	E10 uint16 // 氢气最高压力
	E11 byte // 氢气最高压力传感器代号
	E12 byte // 高压 DC/DC状态
}

// 发动机数据 0x4
type P04 struct {
	E1 byte // 发动机状态
	E2 uint16 // 曲轴转速
	E3 uint16 // 燃料消耗率
}

// 车辆位置数据 0x5
type P05 struct {
	E1 byte // 定位状态
	E2 string // 经度
	E3 string // 纬度
}

// 极值数据 0x6
type P06 struct {
	E1 byte // 最高电压电池子系统号
	E2 byte // 最高电压电池单体代号
	E3 uint16 // 电池单体电压最高值 单位0.001
	E4 byte // 最低电压电池子系统号
	E5 byte // 最低电压电池单体代号
	E6 uint16 // 电池单体电压最低值 单位0.001
	E7 byte // 最高温度子系统号
	E8 byte // 最高温度探针序号
	E9 int16 // 最高温度值
	E10 byte // 最低温度子系统号
	E11 byte // 最低温度探针序号
	E12 int16 // 最低温度值
}

// 报警数据 0x7
type P07 struct {
	E1 byte // 最高报警等级
	E2 string // 通用报警标志
	E3 byte // 可充电储能装置故障总数N1
	E4 string // 可充电储能装置故障代码列表
	E5 byte // 驱动点击故障总数N2
	E6 string // 驱动电机故障代码列表
	E7 byte // 发动机故障总数N3
	E8 string // 发动机故障列表
	E9 byte // 其他故障总数N4
	E10 string // 其他故障代码列表
}
