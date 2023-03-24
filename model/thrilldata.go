package model

import "fmt"

type WaitTimeData struct {
	Name     string
	TimeStr  string
	WaitTime float64
}

type StrategyDesc struct {
	Name              string
	BonusesWait       float64
	NormalWait        float64
	SavingTime        float64
	CostEffectiveness float64
}

func (sd *StrategyDesc) Desc(index int) string {
	return fmt.Sprintf("第%d名：%s || 早享时间：%.2f | 普通时间：%.2f | 节省时间：%.2f | 分数：%.2f ",
		index+1, GetProjectName(sd.Name), sd.BonusesWait, sd.NormalWait, sd.SavingTime, sd.CostEffectiveness)
}
