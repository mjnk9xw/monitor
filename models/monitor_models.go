package models

import (
	"sync"
	"time"
)

func onceInitMonitorModel() {
	modelInstance = NewMonitorService()
}

type MonitorService struct {
	ActionDailyLog  *sync.Map //map[string]*DayStat
	RequestDailyLog *sync.Map //map[string]*DayStat
	LastHour        int
	LastTime        time.Time
}

func NewMonitorService() MonitorServiceIf {
	return &MonitorService{
		ActionDailyLog:  &sync.Map{}, //make( map[string]*DayStat),
		RequestDailyLog: &sync.Map{}, //make ( map[string]*DayStat) ,
		LastHour:        time.Now().Hour(),
		LastTime:        time.Now(),
	}
}
