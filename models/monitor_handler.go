package models

import (
	"fmt"
	"monitor/data"
	"monitor/utils/set_monitor"
	"sync"
	"time"

	"github.com/astaxie/beego"
)

// GetActionStat get time action
func (md *MonitorService) GetActionStat(actionName string) *data.DayStat {
	res, _ := md.ActionDailyLog.LoadOrStore(actionName, &data.DayStat{})
	return res.(*data.DayStat)
}

// GetRequestStat get time request
func (md *MonitorService) GetRequestStat(requestName string) *data.DayStat {
	res, _ := md.RequestDailyLog.LoadOrStore(requestName, &data.DayStat{})
	return res.(*data.DayStat)
}

// CheckReset check reset
func (md *MonitorService) CheckReset() {
	// currentHour := time.Now().Hour()
	// if md.LastHour > currentHour {
	md.LastHour = time.Now().Hour()
	md.ActionDailyLog = &sync.Map{} //make( map[string]*data.DayStat),
	md.RequestDailyLog = &sync.Map{}
	md.LastTime = time.Now()
	//}
}

// LogAction log action
func (md *MonitorService) LogAction(actionName string, runTime time.Duration) {
	// fmt.Println("LogAction ", actionName, runTime)
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentSecondsInDay := currentTime.Hour()*3600 + currentTime.Minute()*60 + currentTime.Second()

	actionStat := md.GetActionStat(actionName)

	actionStat.SecondData[currentSecondsInDay].TotalCount++
	actionStat.SecondData[currentSecondsInDay].TotalTime += runTime

	currentMinsInday := currentTime.Hour()*60 + currentTime.Minute()
	actionStat.MinuteData[currentMinsInday].TotalCount++
	actionStat.MinuteData[currentMinsInday].TotalTime += runTime
	if runTime > actionStat.MinuteData[currentMinsInday].BiggestTime {
		actionStat.MinuteData[currentMinsInday].BiggestTime = runTime
	}

	actionStat.HourData[currentHour].TotalCount++
	actionStat.HourData[currentHour].TotalTime += runTime
	if runTime > actionStat.HourData[currentHour].BiggestTime {
		actionStat.HourData[currentHour].BiggestTime = runTime
	}

	actionStat.WholeDay.TotalCount++
	actionStat.WholeDay.TotalTime += runTime
}

// LogRequest log request
func (md *MonitorService) LogRequest(requestName string, runTime time.Duration) {

	// fmt.Println("LogRequest ", requestName, runTime)
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentSecondsInDay := currentTime.Hour()*3600 + currentTime.Minute()*60 + currentTime.Second()

	actionStat := md.GetRequestStat(requestName)

	actionStat.SecondData[currentSecondsInDay].TotalCount++
	actionStat.SecondData[currentSecondsInDay].TotalTime += runTime

	currentMinsInday := currentTime.Hour()*60 + currentTime.Minute()
	actionStat.MinuteData[currentMinsInday].TotalCount++
	actionStat.MinuteData[currentMinsInday].TotalTime += runTime
	if runTime > actionStat.MinuteData[currentMinsInday].BiggestTime {
		actionStat.MinuteData[currentMinsInday].BiggestTime = runTime
	}

	actionStat.HourData[currentHour].TotalCount++
	actionStat.HourData[currentHour].TotalTime += runTime
	if runTime > actionStat.HourData[currentHour].BiggestTime {
		actionStat.HourData[currentHour].BiggestTime = runTime
	}

	actionStat.WholeDay.TotalCount++
	actionStat.WholeDay.TotalTime += runTime

}

// MonitorCheckpoint struct checkpoint time
type MonitorCheckpoint struct {
	StartTime time.Time
	Name      string
}

// StartAction start action
func (md *MonitorService) StartAction(actionName string) MonitorCheckpoint {
	return MonitorCheckpoint{Name: actionName, StartTime: time.Now()}
}

// FinishAction finish action
func (md *MonitorService) FinishAction(cp MonitorCheckpoint) {
	md.LogAction(cp.Name, time.Since(cp.StartTime))
}

// StartRequest start request
func (md *MonitorService) StartRequest(reqName string) MonitorCheckpoint {
	return MonitorCheckpoint{Name: reqName, StartTime: time.Now()}
}

// FinishRequest finish request
func (md *MonitorService) FinishRequest(cp MonitorCheckpoint) {
	fmt.Println("Finish Request ", cp)
	md.LogRequest(cp.Name, time.Since(cp.StartTime))
}
func (md *MonitorService) MakeMonitor() {
	address := beego.AppConfig.String("monitor_option::etcd")
	name := beego.AppConfig.String("monitor_option::name")
	url := beego.AppConfig.String("monitor_option::url")
	fmt.Println(address, name, url)
	set_monitor.SetURLToMonitorAll(address, name, url)
	return
}
