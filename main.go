package main

import (
	"LittleBeeMark/disneyland/business"
	"LittleBeeMark/disneyland/model"
	"LittleBeeMark/disneyland/util"
	"fmt"
	"sync"
	"time"
)

// "Happy Circle":               "快乐圈",
// "Shipwreck Shore":            "沉船岸边",
// "Slinky Dog Spin":            "弹簧狗团团转",
// "Explorer Canoes":            "探索者独木舟",
// "Avatar: Explore Pandora ":   "阿凡达：探索潘多拉",
// "Jet Packs":                  "喷气背包飞行器",
// "Rex’s Racer":                "泡泡龙冲天赛车",
// "TRON Lightcyc... Chevrolet": "创极速光轮",
// "Seven Dwarfs Mine Train":    "七个小矮人矿工",
// "Storybook Court":            "晶彩奇航",
// "Alice in Wonderland Maze":   "爱丽丝仙境梦游",
// "Fantasia Carousel":          "幻想曲旋转木马",
// "Challenge Tra... Discovery": "挑战之旅... 发现",
// "Soaring Over the Horizon":   "翱翔地平线",
// "Stitch Encounter":           "缝隙相遇",
// "Woody’s Roundup":            "胡迪牛仔嘉年华",
// "Meet Mickey":                "米奇俱乐部",
// "Pirates of th...n Treasure": "加勒比海盗-沉落宝藏之战",
// "Roaring Rapids":             "雷鸣山漂流",
// "Dumbo the Fly...g Elephant": "小飞象",
// "Peter Pan’s Flight":         "彼得-潘的飞行",
// "The Many Adve...e the Pooh": "小熊维尼历险记",
// "“Once Upon a ... Adventure": "小飞侠天空奇遇",
// "Become Iron M...l Universe": "漫威英雄总部：变身钢铁侠",
// "Buzz Lightyea...net Rescue": "巴斯光年星际营救",
// "Siren's Revenge":            "探秘海妖复仇号",
// "Voyage to the...tal Grotto": "古迹探索营的探索步行道",
// "Hunny Pot Spin":             "旋转蜂蜜罐",
var myProjects = []string{
	"Soaring Over the Horizon",
	"Seven Dwarfs Mine Train",
	"TRON Lightcyc... Chevrolet",
	"Roaring Rapids",
	"Pirates of th...n Treasure",
	"Slinky Dog Spin",
	"Woody’s Roundup",
	"Jet Packs",
	"Rex’s Racer",
	"The Many Adve...e the Pooh",
	"Hunny Pot Spin",
	"Dumbo the Fly...g Elephant",
	"Slinky Dog Spin",
	"Voyage to the...tal Grotto",
	"Peter Pan’s Flight",
	"Hunny Pot Spin",
	"Buzz Lightyea...net Rescue",
}

func main() {
	stopChan := make(chan struct{}, 1)
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		defer group.Done()
		for {
			select {
			case <-stopChan:
				return
			default:
				fmt.Println("迪士尼项目策略生成中，有点慢，请稍等片刻。。。。")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	group.Add(1)
	go func() {
		defer group.Done()
		dateList := util.GetLatestWeekDayList(time.Sunday, 6)
		got, err := business.GetAllDataOfProjectByDateFusion(dateList,
			func(m map[string][]*model.WaitTimeData) {
				business.FilterTimeAfter(19, m)
			})
		if err != nil {
			fmt.Println("GetAllDataOfProjectByDateFusion err : ", err)
			return
		}

		sts := business.GenerateStrategy(got, myProjects)
		fmt.Println("========================= 迪士尼策略优先级排行榜 =========================")
		for i, st := range sts {
			fmt.Println(st.Desc(i))
		}

		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("========================= 单个项目不同时间点等待时间排名 =========================")
		sortListGot, err := business.DescSortListData(got, dateList, myProjects, 50)
		if err != nil {
			fmt.Println("DescSortListData err : ", err)
			return
		}
		for _, data := range sortListGot {
			fmt.Println(data)
			fmt.Println()
		}

		stopChan <- struct{}{}
	}()

	group.Wait()
	fmt.Println("迪士尼项目策略生成完毕。。。")
}
