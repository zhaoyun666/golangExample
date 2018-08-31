package excel

import "github.com/tealeg/xlsx"
import "github.com/micro/go-log"

func Excel(obd []*ObdDevice) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1 *xlsx.Row
	var cell *xlsx.Cell
	var err error
	var nav = [6]string{"车架号", "设备号", "总里程", "所属经销商", "下线日期", "活跃状态"}
	var width = [6]float64{16, 20, 20, 16, 45, 16}

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		log.Log(err.Error())
	}

	// 创建头部
	row = sheet.AddRow()
	row.SetHeightCM(1)
	for k, v := range nav {
		log.Log(k, k+1, width[k])
		sheet.SetColWidth(k, k+1, width[k])
		cell = row.AddCell()
		cell.Value = v
	}

	// 创建body

	for _, val := range obd {
		row1 = sheet.AddRow()
		row1.SetHeightCM(1)
		cell = row1.AddCell()
		cell.Value = val.Vin
		cell = row1.AddCell()
		cell.Value = val.Deviceidstring
		cell = row1.AddCell()
		cell.Value = val.Ccid
		cell = row1.AddCell()
		cell.Value = val.PinNumber
	}

	// 保存文件
	err = file.Save("/home/zc/Desktop/Excel/test.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
