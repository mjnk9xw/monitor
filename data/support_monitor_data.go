package data

type TheAverageTimeMonitorData struct {
	TimeAction  []int64 `json:"the_average_time_action"`
	TimeRequest []int64 `json:"the_average_time_request"`
}
type BestCountMonitorData struct {
	BestActionLabel  []string `json:"best_action_label_count"`
	BestActionCount  []int64  `json:"best_action_count"`
	BestRequestLabel []string `json:"best_request_label_count"`
	BestRequestCount []int64  `json:"best_request_count"`
}

type MonitorData struct {
	ActionDataByHours    map[string][24]SecondStat      `json:"action_data_by_hours"`
	ActionDataByMinutes  map[string][24 * 60]SecondStat `json:"action_data_by_minutes"`
	ActionDataHours      []int64                        `json:"action_data_hours"`
	ActionDataCount      []int64                        `json:"action_data_count"`
	RequestDataByHours   map[string][24]SecondStat      `json:"request_data_by_hours"`
	RequestDataByMinutes map[string][24 * 60]SecondStat `json:"request_data_by_minutes"`
	RequestDataHours     []int64                        `json:"request_data_hours"`
	RequestDataCount     []int64                        `json:"request_data_count"`
}

type AllInfoForChartingData struct {
	AllAction                 []string                   `json:"all_action"`
	AllRequest                []string                   `json:"all_request"`
	TheAverageTimeMonitorData *TheAverageTimeMonitorData `json:"the_average_time_monitor"`
	BestCountMonitorData      *BestCountMonitorData      `json:"best_count_monitor"`
	MonitorData               *MonitorData               `json:"monitor_data"`
}
