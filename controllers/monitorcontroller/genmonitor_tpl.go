package monitorcontroller

import (
	"monitor/models"
)

// @Title GenMonitor
// @Description view statistics request
// @Success 200 {object} responses.BoolResponse
// @router /genmonitor [get]
func (this *MonitorController) GenMonitor() {

	m := models.GetMonitorModel()
	if m != nil {

		resp := m.GetAllInfoForCharting()

		this.Data["the_average_time_action"] = resp.TheAverageTimeMonitorData.TimeAction
		this.Data["the_average_time_request"] = resp.TheAverageTimeMonitorData.TimeRequest

		this.Data["best_action_label_count"] = resp.BestCountMonitorData.BestActionLabel
		this.Data["best_action_count"] = resp.BestCountMonitorData.BestActionCount
		this.Data["best_request_label_count"] = resp.BestCountMonitorData.BestRequestLabel
		this.Data["best_request_count"] = resp.BestCountMonitorData.BestRequestCount

		this.Data["all_action"] = resp.AllAction
		this.Data["all_request"] = resp.AllRequest

		this.Data["action_data_hours"] = resp.MonitorData.ActionDataHours
		this.Data["action_data_count"] = resp.MonitorData.ActionDataCount
		this.Data["request_data_hours"] = resp.MonitorData.RequestDataHours
		this.Data["request_data_count"] = resp.MonitorData.RequestDataCount

		labelLoadMinute := make([]int, 0)
		for i := 0; i < 1440; i++ {
			labelLoadMinute = append(labelLoadMinute, i)
		}
		labelLoadHour := make([]int, 0)
		for i := 0; i < 24; i++ {
			labelLoadHour = append(labelLoadHour, i)
		}

		this.Data["labelLoadMinute"] = labelLoadMinute
		this.Data["labelLoadHour"] = labelLoadHour
		this.Data["action_data_minutes"] = resp.MonitorData.ActionDataByMinutes
		this.Data["request_data_minutes"] = resp.MonitorData.RequestDataByMinutes
		this.Data["action_data_hours_line"] = resp.MonitorData.ActionDataByHours
		this.Data["request_data_hours_line"] = resp.MonitorData.RequestDataByHours

		// option ajax
		this.Data["url_data"] = "/monitor/getall"
		this.Data["time_load"] = 60

		this.TplName = "genmonitor.tpl"

	}
}
