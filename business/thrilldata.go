package business

import (
	"LittleBeeMark/disneyland/client"
	"LittleBeeMark/disneyland/model"
	"LittleBeeMark/disneyland/util"
	"fmt"
	"sort"
	"strings"
)

func GetAllDataOfProjectByDate(date string) (map[string][]*model.WaitTimeData, error) {
	res5Map, err := client.GetEvery5minWaitTime(date)
	if err != nil {
		return nil, err
	}

	res15Map, err := client.GetEvery15minWaitTime(date)
	if err != nil {
		return nil, err
	}

	reshMap, err := client.GetEveryHourWaitTime(date)
	if err != nil {
		return nil, err
	}

	for name, _ := range res5Map {
		if data15s, ok := res15Map[name]; ok {
			newData5s := []*model.WaitTimeData{data15s[0]}
			res5Map[name] = append(newData5s, res5Map[name]...)
		}

		if datahs, ok := reshMap[name]; ok {
			newData5s := []*model.WaitTimeData{datahs[0]}
			res5Map[name] = append(newData5s, res5Map[name]...)
		}
	}

	return res5Map, nil
}

func GetAllDataOfProjectByDateByFilter(date string,
	filter func(map[string][]*model.WaitTimeData)) (map[string][]*model.WaitTimeData, error) {
	retMap, err := GetAllDataOfProjectByDate(date)
	if err != nil {
		return nil, err
	}
	if filter != nil {
		filter(retMap)
	}

	return retMap, nil
}

func GetAllDataOfProjectByDateFusion(dateList []string, filter func(map[string][]*model.WaitTimeData)) (map[string][]*model.WaitTimeData, error) {
	retMapList := make([]map[string][]*model.WaitTimeData, 0, len(dateList))
	for _, date := range dateList {
		retMap, err := GetAllDataOfProjectByDateByFilter(date, filter)
		if err != nil {
			return nil, err
		}

		retMapList = append(retMapList, retMap)
	}

	retMapFinal := make(map[string][]*model.WaitTimeData)
	for _, retMap := range retMapList {
		for name, datas := range retMap {

			ddd, ok := retMapFinal[name]
			if !ok {
				retMapFinal[name] = datas
				continue
			}
			retMapFinal[name] = mergeData(ddd, datas)
		}
	}

	return retMapFinal, nil
}

// 融合数据
func mergeData(dst, src []*model.WaitTimeData) []*model.WaitTimeData {
	sourceMap := make(map[string]*model.WaitTimeData)
	for _, srcData := range src {
		sourceMap[srcData.TimeStr] = srcData
	}
	for _, data := range dst {
		if v, ok := sourceMap[data.TimeStr]; ok {
			if v.WaitTime >= 5 {
				data.WaitTime = (data.WaitTime + v.WaitTime) / 2
			}
		}
	}

	return dst
}

// GetWaitTimeByTime 获取某时间点各项目的等待时间
func GetWaitTimeByTime(date, tiemStr string) error {
	resultMap, err := GetAllDataOfProjectByDate(date)
	if err != nil {
		return err
	}

	for name, datas := range resultMap {
		fmt.Printf("=========== name : %s ==============\n", model.GetProjectName(name))
		for _, data := range datas {
			if tiemStr == data.TimeStr {
				fmt.Printf("timeStr : %s, waittime : %v \n", data.TimeStr, data.WaitTime)
			}
		}
	}

	return nil
}

func removeElements(slice []*model.WaitTimeData, indices []int) []*model.WaitTimeData {
	// 将要删除的索引从小到大排序
	sort.Ints(indices)

	// 遍历要删除的索引，每次删除一个元素
	for i := len(indices) - 1; i >= 0; i-- {
		index := indices[i]
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}

func DescSortListData(resultMap map[string][]*model.WaitTimeData, datelist, projects []string, num int) ([]string, error) {
	needProjects := make(map[string]struct{}, len(projects))
	for _, project := range projects {
		needProjects[project] = struct{}{}
	}

	resultDesc := make([]string, 0, len(projects)+1)
	resultDesc = append(resultDesc, fmt.Sprintf("%s 项目等待时间点排名: ", strings.Join(datelist, ",")))
	for name, datas := range resultMap {
		if _, ok := needProjects[name]; !ok {
			continue
		}

		// 过滤没开门的数据
		var delIndex []int
		for i, data := range datas {
			if data.WaitTime == -1 {
				delIndex = append(delIndex, i)
			}
		}
		datas = removeElements(datas, delIndex)

		sort.SliceStable(datas, func(i, j int) bool {
			return datas[i].WaitTime < datas[j].WaitTime
		})

		// 截取
		var realDatas []*model.WaitTimeData
		if num >= 100 {
			realDatas = datas
		} else {
			realDatas = datas[:num]
		}
		// 形成描述
		var desc string = model.GetProjectName(name) + ": \n"
		for i, data := range realDatas {
			desc += fmt.Sprintf(" %d. %s : %v \n", i+1, data.TimeStr, data.WaitTime)
		}

		resultDesc = append(resultDesc, desc)
	}

	return resultDesc, nil
}

func GetWaitTimeSortListByProjectName(date string, projects []string, num int,
	filter func(map[string][]*model.WaitTimeData)) ([]string, error) {
	resultMap, err := GetAllDataOfProjectByDateByFilter(date, filter)
	if err != nil {
		return nil, err
	}

	return DescSortListData(resultMap, []string{date}, projects, num)
}

func GetWaitTimeSortListByProjectNameFusion(datelist []string, projects []string, num int,
	filter func(map[string][]*model.WaitTimeData)) ([]string, error) {
	resultMap, err := GetAllDataOfProjectByDateFusion(datelist, filter)
	if err != nil {
		return nil, err
	}

	return DescSortListData(resultMap, datelist, projects, num)
}

// 红利为期8：00-9：00，游玩期为 9：10 - 18：00

// GenerateStrategy  生成游玩策略
func GenerateStrategy(resultMap map[string][]*model.WaitTimeData, projects []string) []*model.StrategyDesc {
	needProjects := make(map[string]struct{}, len(projects))
	for _, project := range projects {
		needProjects[project] = struct{}{}
	}

	var descInfo []*model.StrategyDesc
	for name, datas := range resultMap {
		if _, ok := needProjects[name]; !ok {
			continue
		}

		var (
			bonusesWait float64
			bonusesNum  float64
			normalWait  float64
			normalNum   float64
		)
		for _, data := range datas {
			if data.TimeStr == "Average" {
				continue
			}
			// 红利期平均时间
			if util.GetStrInTimeZone(data.TimeStr, "08:00 AM", "09:10 AM") && data.WaitTime != -1 {
				bonusesWait += data.WaitTime
				bonusesNum++
			}
			// 普通平均时间
			if util.GetStrInTimeZone(data.TimeStr, "09:15 AM", "05:30 PM") && data.WaitTime != -1 {
				normalWait += data.WaitTime
				normalNum++
			}
		}

		newStrategy := &model.StrategyDesc{
			Name:        name,
			BonusesWait: bonusesWait / bonusesNum,
			NormalWait:  normalWait / normalNum,
		}
		newStrategy.SavingTime = newStrategy.NormalWait - newStrategy.BonusesWait
		newStrategy.CostEffectiveness = newStrategy.SavingTime / (model.GetProjectSpeendTime(name) + bonusesWait)
		descInfo = append(descInfo, newStrategy)
	}

	sort.SliceStable(descInfo, func(i, j int) bool {
		return descInfo[i].CostEffectiveness > descInfo[j].CostEffectiveness
	})
	return descInfo
}
