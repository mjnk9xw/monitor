package responses

import "monitor/data"

type ActionsDataByHours struct {
	HourData map[string][24]data.SecondStat `json:"data"` // reset hàng ngày
}

type GetListAction struct {
	Data []string `json:"data"`
}

type ActionsDataByMinutes struct {
	HourData map[string][24 * 60]data.SecondStat `json:"data"` // reset hàng ngày
}
