package pointer

import (
	"github.com/tealeg/xlsx"
	"log"
)

func Create() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell
	var err error
	var nav = [6]string{"车架号", "设备号", "总里程", "所属经销商", "下线日期", "活跃状态"}
	var content = map[int32]string{
		1: "Y",
		2: "x",
		3: "Y", 4: "x", 5: "Y", 6: "Z",
	}

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		log.Println(err.Error())
	}
	// 创建头部
	row = sheet.AddRow()
	row.SetHeightCM(1)
	for _, v := range nav {
		cell = row.AddCell()
		cell.Value = v
	}

	// 创建body

	for _, val := range content {
		row1 = sheet.AddRow()
		row1.SetHeightCM(1)
		cell = row1.AddCell()
		cell.Value = val
		cell = row1.AddCell()
		cell.Value = val
		cell = row1.AddCell()
		cell.Value = val
		cell = row1.AddCell()
		cell.Value = val
		cell = row1.AddCell()
		cell.Value = val
		cell = row1.AddCell()
		cell.Value = val
	}

	for _, val := range content {
		row2 = sheet.AddRow()
		row2.SetHeightCM(1)
		cell = row2.AddCell()
		cell.Value = val
		cell = row2.AddCell()
		cell.Value = val
		cell = row2.AddCell()
		cell.Value = val
		cell = row2.AddCell()
		cell.Value = val
		cell = row2.AddCell()
		cell.Value = val
		cell = row2.AddCell()
		cell.Value = val
	}

	// 保存文件
	err = file.Save("/home/zc/Desktop/Excel/test.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
}

/*func createHeader(nav [6]string) {
	row = sheet.AddRow()
	row.SetHeightCM(1)
	for _, v := range nav {
		cell = row.AddCell()
		cell.Value = v
	}
}

func createBody() {
	row1 = sheet.AddRow()
	row1.SetHeightCM(1)
	cell = row1.AddCell()
	cell.Value = "狗子"
	cell = row1.AddCell()
	cell.Value = "18"

	row2 = sheet.AddRow()
	row2.SetHeightCM(1)
	cell = row2.AddCell()
	cell.Value = "蛋子"
	cell = row2.AddCell()
	cell.Value = "28"

	row3 = sheet.AddRow()
	row3.SetHeightCM(1)
	cell = row3.AddCell()
	cell.Value = "狗子"
	cell = row3.AddCell()
	cell.Value = "18"

	row4 = sheet.AddRow()
	row4.SetHeightCM(1)
	cell = row4.AddCell()
	cell.Value = "蛋子"
	cell = row4.AddCell()
	cell.Value = "28"


	row5 = sheet.AddRow()
	row5.SetHeightCM(1)
	cell = row5.AddCell()
	cell.Value = "狗子"
	cell = row5.AddCell()
	cell.Value = "18"

	row6 = sheet.AddRow()
	row6.SetHeightCM(1)
	cell = row6.AddCell()
	cell.Value = "蛋子"
	cell = row6.AddCell()
	cell.Value = "28"
}

func saveFile(filename string) {
	err = file.Save(filename)
	if err != nil {
		log.Fatalln(err)
	}
}*/
