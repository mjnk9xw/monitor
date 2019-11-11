package models

import (
	"sync"
	"time"
)

type MonitorModel struct {
	ActionDailyLog  *sync.Map //map[string]*DayStat
	RequestDailyLog *sync.Map //map[string]*DayStat
	LastHour        int
	LastTime        time.Time
}

func NewMonitorModel() MonitorModelIf {
	return &MonitorModel{
		ActionDailyLog:  &sync.Map{}, //make( map[string]*DayStat),
		RequestDailyLog: &sync.Map{}, //make ( map[string]*DayStat) ,
		LastHour:        time.Now().Hour(),
		LastTime:        time.Now(),
	}
}
