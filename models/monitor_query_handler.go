package models

import (
	"fmt"
	"monitor/data"
	"monitor/responses"
	"sort"
)

// GetAllActionNames return all action
func (md *MonitorService) GetAllActionNames() ([]string, error) {
	res := []string{}
	// var res []string;
	fmt.Println("GetAllActionNames ")
	md.ActionDailyLog.Range(func(key, _ interface{}) bool {
		sKey := key.(string)
		res = append(res, sKey)
		// dailyStat := value.(*DayStat)
		return true
	})

	sort.Strings(res)
	return res, nil
}

// GetAllRequestNames return all request
func (md *MonitorService) GetAllRequestNames() ([]string, error) {
	res := []string{}
	md.RequestDailyLog.Range(func(key, _ interface{}) bool {
		sKey := key.(string)
		res = append(res, sKey)
		// dailyStat := value.(*DayStat)
		return true
	})

	sort.Strings(res)
	return res, nil
}

// GetActionDataByHour return time action
func (md *MonitorService) GetActionDataByHour(out *responses.ActionsDataByHours) {

	md.ActionDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		out.HourData[sKey] = aDayStat.HourData //need to copy?
		fmt.Printf("action data of %s, %v \n", sKey, aDayStat.HourData)
		return true
	})
}

// GetRequestDataByHour return time request
func (md *MonitorService) GetRequestDataByHour(out *responses.ActionsDataByHours) {
	md.RequestDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		out.HourData[sKey] = aDayStat.HourData //need to copy?

		fmt.Printf("action data of %s, %v \n", sKey, aDayStat.HourData)
		return true
	})
}

// GetActionDataByMinute return time action
func (md *MonitorService) GetActionDataByMinute(out *responses.ActionsDataByMinutes) {
	md.ActionDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		out.HourData[sKey] = aDayStat.MinuteData //need to copy?

		fmt.Printf("action data of %s, %v \n", sKey, aDayStat.HourData)
		return true
	})
}

// GetRequestDataByMinute return time request
func (md *MonitorService) GetRequestDataByMinute(out *responses.ActionsDataByMinutes) {
	md.RequestDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		out.HourData[sKey] = aDayStat.MinuteData //need to copy?
		fmt.Printf("action data of %s, %v \n", sKey, aDayStat.HourData)
		return true
	})
}

// GetAllInfoForCharting get tất cả thông tin cơ bản cho vẽ biểu đồ
func (md *MonitorService) GetAllInfoForCharting() (resp *data.AllInfoForChartingData) {
	fmt.Printf("[GetAllInfoForCharting] call \n")

	resp = &data.AllInfoForChartingData{
		AllAction:  make([]string, 0),
		AllRequest: make([]string, 0),
		TheAverageTimeMonitorData: &data.TheAverageTimeMonitorData{
			TimeAction:  make([]int64, 0),
			TimeRequest: make([]int64, 0),
		},
		BestCountMonitorData: &data.BestCountMonitorData{
			BestActionCount:  make([]int64, 0),
			BestActionLabel:  make([]string, 0),
			BestRequestCount: make([]int64, 0),
			BestRequestLabel: make([]string, 0),
		},
		MonitorData: &data.MonitorData{
			ActionDataCount:      make([]int64, 0),
			ActionDataHours:      make([]int64, 0),
			ActionDataByMinutes:  make(map[string][24 * 60]data.SecondStat),
			ActionDataByHours:    make(map[string][24]data.SecondStat),
			RequestDataCount:     make([]int64, 0),
			RequestDataHours:     make([]int64, 0),
			RequestDataByMinutes: make(map[string][24 * 60]data.SecondStat),
			RequestDataByHours:   make(map[string][24]data.SecondStat),
		},
	}

	resp.AllAction, _ = md.GetAllActionNames()
	resp.AllRequest, _ = md.GetAllRequestNames()

	//aResActionsMinuteData := make(map[string][24 * 60]data.SecondStat)
	md.ActionDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		resp.MonitorData.ActionDataByMinutes[sKey] = aDayStat.MinuteData //need to copy?
		return true
	})

	//aResRequestsMinuteData := make(map[string][24 * 60]data.SecondStat)
	md.RequestDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		resp.MonitorData.RequestDataByMinutes[sKey] = aDayStat.MinuteData //need to copy?
		return true
	})

	// aResActionsHour := make(map[string][24]data.SecondStat)
	md.ActionDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		resp.MonitorData.ActionDataByHours[sKey] = aDayStat.HourData //need to copy?
		return true
	})

	//aResRequestsHour := make(map[string][24]data.SecondStat)
	md.RequestDailyLog.Range(func(key, value interface{}) bool {
		// add key - value to
		sKey := key.(string)
		aDayStat := value.(*data.DayStat)
		resp.MonitorData.RequestDataByHours[sKey] = aDayStat.HourData //need to copy?
		return true
	})

	for _, v := range resp.AllAction {
		lstHours := resp.MonitorData.ActionDataByHours[v]
		total := int64(0)
		totalTime := int64(0)
		for _, hours := range lstHours {
			total += int64(hours.TotalCount)
			totalTime += int64(hours.TotalTime)
		}
		resp.MonitorData.ActionDataCount = append(resp.MonitorData.ActionDataCount, total)
		resp.MonitorData.ActionDataHours = append(resp.MonitorData.ActionDataHours, totalTime)

		if total > 0 {
			resp.TheAverageTimeMonitorData.TimeAction = append(resp.TheAverageTimeMonitorData.TimeAction, totalTime/total)
		} else {
			resp.TheAverageTimeMonitorData.TimeAction = append(resp.TheAverageTimeMonitorData.TimeAction, 0)
		}
	}

	// lstDataRequestHours := make([]int64, 0)
	// lstDataRequest := make([]int64, 0)
	for _, v := range resp.AllRequest {
		lstHours := resp.MonitorData.RequestDataByHours[v]
		total := int64(0)
		totalTime := int64(0)
		for _, hours := range lstHours {
			total += int64(hours.TotalCount)
			totalTime += int64(hours.TotalTime)
		}
		resp.MonitorData.RequestDataCount = append(resp.MonitorData.RequestDataCount, total)
		resp.MonitorData.RequestDataHours = append(resp.MonitorData.RequestDataHours, totalTime)

		if total > 0 {
			resp.TheAverageTimeMonitorData.TimeRequest = append(resp.TheAverageTimeMonitorData.TimeRequest, totalTime/total)
		} else {
			resp.TheAverageTimeMonitorData.TimeRequest = append(resp.TheAverageTimeMonitorData.TimeRequest, 0)
		}
	}
	resp.BestCountMonitorData.BestActionLabel = make([]string, len(resp.AllAction))
	resp.BestCountMonitorData.BestActionCount = make([]int64, len(resp.MonitorData.ActionDataCount))
	resp.BestCountMonitorData.BestRequestLabel = make([]string, len(resp.AllRequest))
	resp.BestCountMonitorData.BestRequestCount = make([]int64, len(resp.MonitorData.RequestDataCount))
	copy(resp.BestCountMonitorData.BestActionLabel, resp.AllAction)
	copy(resp.BestCountMonitorData.BestActionCount, resp.MonitorData.ActionDataCount)
	copy(resp.BestCountMonitorData.BestRequestLabel, resp.AllRequest)
	copy(resp.BestCountMonitorData.BestRequestCount, resp.MonitorData.RequestDataCount)
	for i := 0; i < len(resp.BestCountMonitorData.BestActionCount)-1; i++ {
		for j := i + 1; j < len(resp.BestCountMonitorData.BestActionCount); j++ {
			if resp.BestCountMonitorData.BestActionCount[j] > resp.BestCountMonitorData.BestActionCount[i] {
				resp.BestCountMonitorData.BestActionCount[j], resp.BestCountMonitorData.BestActionCount[i] = resp.BestCountMonitorData.BestActionCount[i], resp.BestCountMonitorData.BestActionCount[j]
				resp.BestCountMonitorData.BestActionLabel[j], resp.BestCountMonitorData.BestActionLabel[i] = resp.BestCountMonitorData.BestActionLabel[i], resp.BestCountMonitorData.BestActionLabel[j]
			}
		}
	}
	for i := 0; i < len(resp.BestCountMonitorData.BestRequestCount)-1; i++ {
		for j := i + 1; j < len(resp.BestCountMonitorData.BestRequestCount); j++ {
			if resp.BestCountMonitorData.BestRequestCount[j] > resp.BestCountMonitorData.BestRequestCount[i] {
				resp.BestCountMonitorData.BestRequestCount[j], resp.BestCountMonitorData.BestRequestCount[i] = resp.BestCountMonitorData.BestRequestCount[i], resp.BestCountMonitorData.BestRequestCount[j]
				resp.BestCountMonitorData.BestRequestLabel[j], resp.BestCountMonitorData.BestRequestLabel[i] = resp.BestCountMonitorData.BestRequestLabel[i], resp.BestCountMonitorData.BestRequestLabel[j]
			}
		}
	}

	return
}
