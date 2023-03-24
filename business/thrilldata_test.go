package business

import (
	"LittleBeeMark/disneyland/model"
	"LittleBeeMark/disneyland/util"
	"fmt"
	"testing"
	"time"
)

func TestGetAllDataOfProjectByDate(t *testing.T) {
	res5Map, err := GetAllDataOfProjectByDate("2023-03-19")
	if err != nil {
		t.Log("err : ", err)
		return
	}

	for name, datas := range res5Map {
		fmt.Printf("=========== name : %s ==============\n", model.GetProjectName(name))
		for _, data := range datas {
			fmt.Printf("timeStr : %s, waittime : %v \n", data.TimeStr, data.WaitTime)
		}
	}

}
func TestGetAllDataOfProjectByDateFusion(t *testing.T) {
	res5Map, err := GetAllDataOfProjectByDateFusion(util.GetLatestWeekDayList(time.Sunday, 2), nil)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	for name, datas := range res5Map {
		fmt.Printf("=========== name : %s ==============\n", model.GetProjectName(name))
		for _, data := range datas {
			fmt.Printf("timeStr : %s, waittime : %v \n", data.TimeStr, data.WaitTime)
		}
	}

}
func TestGetWaitTimeByTime(t *testing.T) {
	err := GetWaitTimeByTime("2023-03-12", "08:00 AM")
	if err != nil {
		t.Log("err : ", err)
		return
	}

}

//var DisneyProjectNameDic = map[string]string{
//	"Happy Circle":               "快乐圈",
//	"Shipwreck Shore":            "沉船岸边",
//	"Slinky Dog Spin":            "弹簧狗团团转",
//	"Explorer Canoes":            "探索者独木舟",
//	"Avatar: Explore Pandora ":   "阿凡达：探索潘多拉",
//	"Jet Packs":                  "喷气背包飞行器",
//	"Rex’s Racer":                "泡泡龙冲天赛车",
//	"TRON Lightcyc... Chevrolet": "创极速光轮",
//	"Seven Dwarfs Mine Train":    "七个小矮人矿工",
//	"Storybook Court":            "晶彩奇航",
//	"Alice in Wonderland Maze":   "爱丽丝仙境梦游",
//	"Fantasia Carousel":          "幻想曲旋转木马",
//	"Challenge Tra... Discovery": "挑战之旅... 发现",
//	"Soaring Over the Horizon":   "翱翔地平线",
//	"Stitch Encounter":           "缝隙相遇",
//	"Woody’s Roundup":            "胡迪牛仔嘉年华",
//	"Meet Mickey":                "米奇俱乐部",
//	"Pirates of th...n Treasure": "加勒比海盗-沉落宝藏之战",
//	"Roaring Rapids":             "雷鸣山漂流",
//	"Dumbo the Fly...g Elephant": "小飞象",
//	"Peter Pan’s Flight":         "彼得-潘的飞行",
//	"The Many Adve...e the Pooh": "小熊维尼历险记",
//	"“Once Upon a ... Adventure": "小飞侠天空奇遇",
//	"Become Iron M...l Universe": "漫威英雄总部：变身钢铁侠",
//	"Buzz Lightyea...net Rescue": "巴斯光年星际营救",
//	"Siren's Revenge":            "探秘海妖复仇号",
//	"Voyage to the...tal Grotto": "古迹探索营的探索步行道",
//	"Hunny Pot Spin":             "旋转蜂蜜罐",
//}

func TestGetWaitTimeSortListByProjectName(t *testing.T) {
	got, err := GetWaitTimeSortListByProjectName("2023-03-12", projects, 50, nil)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	for _, data := range got {
		fmt.Println(data)
		fmt.Println()
	}
}

var (
	projects = []string{
		"Soaring Over the Horizon",
		"Seven Dwarfs Mine Train",
		"TRON Lightcyc... Chevrolet",
		"Roaring Rapids",
		"Pirates of th...n Treasure",
		"Slinky Dog Spin",
	}
	projects2 = []string{
		"Woody’s Roundup",
		"Jet Packs",
		"The Many Adve...e the Pooh",
		"Hunny Pot Spin",
		"Dumbo the Fly...g Elephant",
		"“Once Upon a ... Adventure",
		"Slinky Dog Spin",
	}

	projects3 = []string{
		"Soaring Over the Horizon",
		"Seven Dwarfs Mine Train",
		"TRON Lightcyc... Chevrolet",
		"Roaring Rapids",
		"Pirates of th...n Treasure",
		"Slinky Dog Spin",
		"Woody’s Roundup",
		"Jet Packs",
		"The Many Adve...e the Pooh",
		"Hunny Pot Spin",
		"Dumbo the Fly...g Elephant",
		"“Once Upon a ... Adventure",
		"Slinky Dog Spin",
		"Voyage to the...tal Grotto",
		"Peter Pan’s Flight",
		"Hunny Pot Spin",
	}
)

func TestGetWaitTimeSortListByProjectNameFusion(t *testing.T) {
	got, err := GetWaitTimeSortListByProjectNameFusion(util.GetLatestWeekDayList(time.Sunday, 6), projects, 50, nil)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	for _, data := range got {
		fmt.Println(data)
		fmt.Println()
	}
}

func TestGetWaitTimeSortListByProjectNameFusion2(t *testing.T) {
	got, err := GetWaitTimeSortListByProjectNameFusion(util.GetLatestWeekDayList(time.Sunday, 6), projects, 50,
		func(m map[string][]*model.WaitTimeData) {
			FilterTimeAfter(19, m)
		})
	if err != nil {
		t.Log("err : ", err)
		return
	}

	for _, data := range got {
		fmt.Println(data)
		fmt.Println()
	}
}

func TestGetAllDataOfProjectByDateByFilter(t *testing.T) {
	GetAllDataOfProjectByDateByFilter("2023-03-12", func(m map[string][]*model.WaitTimeData) {
		FilterTimeAfter(7, m)
	})
}

func TestFilterTimeAfter(t *testing.T) {
	s := []*model.WaitTimeData{
		{Name: "a"},
		{Name: "b"},
		{Name: "c"},
		{Name: "d"},
	}

	removeElements(s, []int{1, 2})
	fmt.Println(s)

}

func TestGenerateStrategy(t *testing.T) {
	got, err := GetAllDataOfProjectByDateFusion(util.GetLatestWeekDayList(time.Sunday, 6),
		func(m map[string][]*model.WaitTimeData) {
			FilterTimeAfter(19, m)
		})
	if err != nil {
		t.Log("err : ", err)
		return
	}

	sts := GenerateStrategy(got, projects3)
	for i, st := range sts {
		fmt.Println(st.Desc(i))
	}
}
