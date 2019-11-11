package models

import (
	"monitor/data"
	"monitor/responses"
	"sync"
	"time"
)

var (
	once          sync.Once
	modelInstance MonitorServiceIf
)

// GetQB86APIModel return interface models
func GetMonitorModel() MonitorServiceIf {
	once.Do(func() {
		onceInitMonitorModel()
	})
	return modelInstance
}

type MonitorServiceIf interface {
	GetActionStat(actionName string) *data.DayStat
	GetRequestStat(requestName string) *data.DayStat
	CheckReset()
	LogAction(actionName string, runTime time.Duration)
	LogRequest(requestName string, runTime time.Duration)
	StartAction(actionName string) MonitorCheckpoint
	FinishAction(cp MonitorCheckpoint)
	StartRequest(reqName string) MonitorCheckpoint
	FinishRequest(cp MonitorCheckpoint)
	GetAllActionNames() ([]string, error)
	GetAllRequestNames() ([]string, error)
	GetActionDataByHour(out *responses.ActionsDataByHours)
	GetRequestDataByHour(out *responses.ActionsDataByHours)
	GetActionDataByMinute(out *responses.ActionsDataByMinutes)
	GetRequestDataByMinute(out *responses.ActionsDataByMinutes)
	GetAllInfoForCharting() (resp *data.AllInfoForChartingData)
	MakeMonitor()
}
