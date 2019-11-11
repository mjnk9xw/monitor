package monitorcontroller

import (
	"monitor/data"
	"monitor/models"
	"monitor/responses"
)

// @Title GetListActions
// @Description Get total action names
// @Success 200 {json} 	responses.GetListAction
// @router /getactions [get]
func (c *MonitorController) GetListActions() {
	m := models.GetMonitorModel()
	if m != nil {
		dataResp, _ := m.GetAllActionNames()
		c.Data["json"] = responses.GetListAction{
			Data: dataResp,
		}
		//m.FinishAction(cp2)
		c.ServeJSON()
		return
	}

}

// @Title GetListRequests
// @Description Get total request names
// @Success 200 {json} 	responses.GetListAction
// @router /getrequests [get]
func (c *MonitorController) GetListRequests() {
	m := models.GetMonitorModel()
	if m != nil {
		dataResp, _ := m.GetAllRequestNames()
		c.Data["json"] = responses.GetListAction{
			Data: dataResp,
		}
		//m.FinishAction(cp2)
		c.ServeJSON()
		return
	}

}

// @Title GetActionsDataByHours
// @Description Get total hourly action data
// @Success 200 {object}	responses.ActionsDataByHours
// @router /totalactionshourly [get]
func (c *MonitorController) GetActionsDataByHours() {
	m := models.GetMonitorModel()
	if m != nil {
		aRes := responses.ActionsDataByHours{
			HourData: make(map[string][24]data.SecondStat),
		}
		m.GetActionDataByHour(&aRes)
		c.Data["json"] = aRes
		c.ServeJSON()
		return
	}
}

// @Title GetRequestDataByHour
// @Description Get total hourly action data
// @Success 200 {object}	responses.ActionsDataByHours
// @router /totalrequestshourly [get]
func (c *MonitorController) GetRequestDataByHour() {
	m := models.GetMonitorModel()
	if m != nil {

		aRes := responses.ActionsDataByHours{
			HourData: make(map[string][24]data.SecondStat),
		}
		m.GetRequestDataByHour(&aRes)
		c.Data["json"] = aRes
		c.ServeJSON()
		return
	}
}

// @Title GetAllInfoForCharting
// @Description Get all info for charting action data
// @Success 200 {object} responses.GetAllInfoForCharting
// @router /getall [get]
func (c *MonitorController) GetAllInfoForCharting() {
	m := models.GetMonitorModel()
	if m != nil {
		resp := m.GetAllInfoForCharting()
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
}
