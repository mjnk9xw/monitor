package data

import "time"

// SecondStat tổng số trong 1 s
type SecondStat struct {
	TotalCount  uint64        `json:"totalcount"`
	TotalTime   time.Duration `json:"totaltime"`
	BiggestTime time.Duration `json:"biggesttime"`
}

// HourStat 1 h
type HourStat struct {
	SecondData [3600]SecondStat `json:"data_insecond"`
}

// DayStat 1 ngày
type DayStat struct {
	// HourData [24]HourStat
	SecondData [24 * 3600]SecondStat `json:"data_insecond"`
	MinuteData [24 * 60]SecondStat   `json:"datainminute"` //
	HourData   [24]SecondStat        `json:"hourlydata"`   // reset hàng ngày
	WholeDay   SecondStat
}
