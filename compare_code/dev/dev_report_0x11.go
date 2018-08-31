package report_handler

import (
	"fmt"
	"net/url"
	"strings"

	"git.linewin.cc/msvr-cluster/comm/logger"
	"git.linewin.cc/msvr-cluster/comm/protocol/util"
	"git.linewin.cc/msvr-cluster/dbworker/models/obd"
	"git.linewin.cc/msvr-cluster/dbworker/models/obdlog"
	//	"git.linewin.cc/msvr-cluster/reportworker/config"
	"git.linewin.cc/msvr-cluster/dbworker/models/member"
	"git.linewin.cc/msvr-cluster/reportworker/report_models"
	"git.linewin.cc/msvr-cluster/reportworker/utils"
)

type Report_0x11 struct {
}

func NewReport_0x11() HandlerInterface {
	report_0x11 := &Report_0x11{}
	return report_0x11
}

// 0x11:17 开机上报
// testing 2016.3.10 ok
// update 2016.6.21 new table
// update 2016.10.18 add electric
func (a *Report_0x11) Report(r *ReportPacket) error {
	var err error

	// 参数处理
	reportData, _ := url.ParseQuery(string(r.P.Data))
	ccid := reportData.Get("CCID")
	version := reportData.Get("version")
	rectime := utils.ReportRectime(reportData.Get("recTime"))
	imei := reportData.Get("IMEI")
	vpin := reportData.Get("pin")
	esk := reportData.Get("ESK")
	vin := reportData.Get("VIN")
	vin = utils.FilterVin([]byte(vin))

	// Report data
	if r.Logid > 0 {
		logBoot := &obdlog.CarObdLogX11{
			Logid:      r.Logid,
			FirmwareID: reportData.Get("firmwareID"),
			Version:    version,
			VC:         reportData.Get("VC"),
			AccePark:   reportData.Get("accePark"),
		}

		if logBoot.Add() != nil {
			logger.Errorf("[%s] 0x%x userid:%d logid:%d Write the data in the log_boot table to failed,err:%v",
				r.P.DeviceID, r.P.Action, r.Car.Uid, r.Logid, err.Error())
		}
	}

	// CleanAll VIN and CCID
	// adolph.liu 2017.8.23
	/*err = report_models.CleanDeviceVin(r.Device.Id, vin)
	if err != nil {
		logger.Errorf("[%s] 0x%x uid=%d lid=%d Clean VIN failed,err:%v",
			r.P.Id(), r.P.Action, r.Car.Uid, r.Logid, err.Error())
	}*/
	err = report_models.CleanDeviceCCID(r.Device.Id, ccid)
	if err != nil {
		logger.Errorf("[%s] 0x%x uid=%d lid=%d Clean CCID failed,err:%v",
			r.P.Id(), r.P.Action, r.Car.Uid, r.Logid, err.Error())
	}

	// Update the data in the log_obd_device table.
	device := &obd.Device{
		Vehicle_status: 1, //1-醒着，2-睡眠，3-掉线
		Hardversion:    version,
		Imei:           imei,
		Cancel:         "2", // 1-未出厂 2-已出厂 3-已作废
		Updatetime:     r.LogTime,
		Isupgrade:      0, // 升级结束
		//Upgradetime:    rectime.Unix(), // 升级结束时间
		//Vin: vin,
		Esk: esk,
	}
	// 当设备和车架号属于新关系时更新设备车架号
	// zhaoce@linewin.cc 2018/05/04
	isUpdate := report_models.IsUpdateVin(r.Device.Id, vin)
	if isUpdate && len(vin) == 17 {
		device.Vin = vin
	}

	// Update the data in the member_car table.
	if len(vin) == 17 {
		memberCar := &member.CarMemberCar{
			Standcarno: vin,
			Updatedate: rectime.Unix(),
		}
		modelsMemberCar := &report_models.MemberCar{
			DB: memberCar,
		}
		modelsMemberCar.Update(r.Car.Id, r.Car.Cuscarid, r.P.DeviceID)
	}

	// Database PIN is empty and report PIN is correct.
	// adolph.liu 2017.8.23
	if r.Device.Pin_number == "" && vpin != "" && vpin != "ffffffff" {
		if len(vpin) < 8 {
			vpin = fmt.Sprintf("%08s", vpin)
		}

		device.Pin_number = strings.ToUpper(vpin)
		device.Pin_number_time = r.LogTime
		device.Is_pin = 2
	}

	report_models.AddDeviceExt(r.Device.Id, r.Device.Pin_number, rectime)

	// CCID
	if ccid != "" && len(ccid) == 20 {
		// CCID是否有变化,第一次更新不算变化
		var chanageCcid bool

		if r.Device.Ccid != "" && ccid != r.Device.Ccid {
			chanageCcid = true
			device.Simupdate = r.Device.Simupdate + 1 // Update the data in the log_obd_device table.
			// Clean CCID
			simTemp := &obd.TempSim{
				IsBind:    2,
				Devstring: "",
			}
			simTemp.UpdateByCCid(r.Device.Ccid)
		}

		// 众泰短信卡激活: 未激活 || CCID变更时均需要激活
		var smsMsg string
		//if chanageCcid && r.Device.Ccidtime > 0 {
		//	// update by 2016.04.26
		//	if r.Device.Sms > 0 && r.Device.Sms != 3 {
		//		sms := &utils.Sms{
		//			Way:            r.Device.Sms,
		//			CardIccid:      r.Device.Ccid,
		//			Agentuser:      "zlink",
		//			Agentpwd:       "123456789",
		//			UserName:       "zlinkUser",
		//			Phone:          "18691635352",
		//			CertificateNum: "640323198711213018",
		//			DeviceId:       r.P.DeviceID,
		//			DeviceType:     "zlink device",
		//		}
		//		msg, err := sms.UploadUserInfo(config.ReportConf.String("sms::sms.address"))
		//		if err != nil {
		//			smsMsg = "CCID卡激活失败"
		//			logger.Errorf("[%s] 0x%x userid=%d logid=%d M2M unicom SMS activation failed,err:%v",
		//				r.P.Id(), r.P.Action, r.Car.Uid, r.Logid, err.Error())
		//		} else {
		//			device.Ccidtime = r.LogTime
		//			smsMsg = "CCID卡激活成功"
		//			logger.Infof("[%s] 0x%x userid=%d logid=%d M2M unicom SMS activation success,resultCode=%d,resultMsg=%s,object=%s",
		//				r.P.Id(), r.P.Action, r.Car.Uid, r.Logid, msg.ResultCode, msg.ResultMsg, msg.Object)
		//		}
		//	}
		//}

		// Update data in the temp_sim table
		sim := &obd.TempSim{Ccid: ccid}
		sim.Get()
		if sim.Id == 0 {
			logger.Infof("[%s] 0x%x CCID=%s car_obd_temp_sim config is empty. err:%v",
				r.Device.Deviceidstring, r.P.Action, ccid, err.Error())
		}

		newSim := new(obd.TempSim)
		newSim.Bindtime = r.Device.Bindtime //绑定时间
		newSim.IsBind = 1                   //是否绑定
		newSim.Devstring = r.P.DeviceID     //设备串
		if sim.Starttime == 0 {
			newSim.Starttime = rectime.Unix()
		}
		_, err = newSim.Update(sim.Id)
		if err != nil {
			logger.Errorf("[%s] 0x%x userid=%d logid=%d Write the data in the temp_sim table to failed. err:%v",
				r.P.Id(), r.P.Action, r.Car.Uid, r.Logid, err.Error())
		}

		device.Ccid = ccid
		device.Ccidtime = r.LogTime
		if sim.Id > 0 {
			device.Simid = sim.Simid
		}

		// SIM&CCID变更日志
		if chanageCcid {
			deviceLog := &obd.CarObdDeviceLog{
				Deviceid:    r.Device.Id,
				Dealerid:    r.Car.Dealerid,
				Operatetype: 1,
				Operateid:   r.Car.Uid,
				Action:      "update",
				Msg:         "SIM:" + r.Device.Simid + "变更" + sim.Simid + ";CCID:" + r.Device.Ccid + "变更" + ccid + "," + smsMsg,
				Rectime:     rectime.Unix(),
			}
			deviceLog.Add()
		}
	}

	// 记录车架号变更信息
	err = report_models.AddDeviceVinHistory(r.Device.Deviceidstring, vin, r.P.Action)
	if err != nil {
		logger.Errorf("[%s] 0x%x uid=%d lid=%d ADD log_obd_device_vin_history table to failed:err:%s", r.P.Id(), r.P.Action, r.Car.Uid, r.Logid, err.Error())
	}

	// 临时存储重复设备上报不同车架号问题
	err = report_models.StoreDuplicationDevice(r.Device.Deviceidstring, vin)
	if err != nil {
		logger.Errorf("[%s] 0x%x lid=%d ADD log_obd_device_duplication table to failed:err:%s", r.P.Id(), r.P.Action, r.Logid, err.Error())
	}

	err = report_models.SetDevice(r.Device.Id, device)
	if err != nil {
		logger.Errorf("[%s] 0x%x userid=%d logid=%d update log_obd_device table to failed,err:%v",
			r.P.Id(), r.P.Action, r.Car.Uid, r.Logid, err.Error())
	}
	return nil
}

func init() {
	Register(util.CS0001, 0x11, NewReport_0x11)
}
