package business

import (
	"LittleBeeMark/disneyland/model"
	"LittleBeeMark/disneyland/util"
)

func FilterTimeAfter(timeHour int, retMap map[string][]*model.WaitTimeData) {
	for name, ddd := range retMap {
		// 过滤没开门的数据
		var delIndex []int
		for i, data := range ddd {
			if data.TimeStr == "Average" {
				continue
			}
			if util.GetTimeStrHour(data.TimeStr) >= timeHour {
				delIndex = append(delIndex, i)
			}
		}

		retMap[name] = removeElements(ddd, delIndex)
	}
}
