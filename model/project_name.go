package model

var DisneyProjectNameDic = map[string]string{
	"Happy Circle":               "快乐圈",
	"Shipwreck Shore":            "沉船岸边",
	"Slinky Dog Spin":            "弹簧狗团团转",
	"Explorer Canoes":            "探索者独木舟",
	"Avatar: Explore Pandora ":   "阿凡达：探索潘多拉",
	"Jet Packs":                  "喷气背包飞行器",
	"Rex’s Racer":                "泡泡龙冲天赛车",
	"TRON Lightcyc... Chevrolet": "创极速光轮",
	"Seven Dwarfs Mine Train":    "七个小矮人矿工",
	"Storybook Court":            "晶彩奇航",
	"Alice in Wonderland Maze":   "爱丽丝仙境梦游",
	"Fantasia Carousel":          "幻想曲旋转木马",
	"Challenge Tra... Discovery": "挑战之旅... 发现",
	"Soaring Over the Horizon":   "翱翔地平线",
	"Stitch Encounter":           "缝隙相遇",
	"Woody’s Roundup":            "胡迪牛仔嘉年华",
	"Meet Mickey":                "米奇俱乐部",
	"Pirates of th...n Treasure": "加勒比海盗-沉落宝藏之战",
	"Roaring Rapids":             "雷鸣山漂流",
	"Dumbo the Fly...g Elephant": "小飞象",
	"Peter Pan’s Flight":         "彼得-潘的飞行",
	"The Many Adve...e the Pooh": "小熊维尼历险记",
	"“Once Upon a ... Adventure": "小飞侠天空奇遇",
	"Become Iron M...l Universe": "漫威英雄总部：变身钢铁侠",
	"Buzz Lightyea...net Rescue": "巴斯光年星际营救",
	"Siren's Revenge":            "探秘海妖复仇号",
	"Voyage to the...tal Grotto": "古迹探索营的探索步行道",
	"Hunny Pot Spin":             "旋转蜂蜜罐",
}

var DisneyProjectSpendTime = map[string]float64{
	"Happy Circle":               5,
	"Shipwreck Shore":            5,
	"Slinky Dog Spin":            3,
	"Explorer Canoes":            5,
	"Avatar: Explore Pandora ":   5,
	"Jet Packs":                  3,
	"Rex’s Racer":                3,
	"TRON Lightcyc... Chevrolet": 5,
	"Seven Dwarfs Mine Train":    6,
	"Storybook Court":            13,
	"Alice in Wonderland Maze":   15,
	"Fantasia Carousel":          7,
	"Challenge Tra... Discovery": 5,
	"Soaring Over the Horizon":   7,
	"Stitch Encounter":           5,
	"Woody’s Roundup":            4,
	"Meet Mickey":                10,
	"Pirates of th...n Treasure": 12,
	"Roaring Rapids":             10,
	"Dumbo the Fly...g Elephant": 5,
	"Peter Pan’s Flight":         15,
	"The Many Adve...e the Pooh": 5,
	"“Once Upon a ... Adventure": 5,
	"Become Iron M...l Universe": 5,
	"Buzz Lightyea...net Rescue": 5,
	"Siren's Revenge":            5,
	"Voyage to the...tal Grotto": 10,
	"Hunny Pot Spin":             3,
}

func GetProjectName(name string) string {
	if v, ok := DisneyProjectNameDic[name]; ok {
		return v
	}

	//fmt.Printf(">>>>>>>>>>>> %+v", name)
	return "没有统计。。。"
}
func GetProjectSpeendTime(name string) float64 {
	if v, ok := DisneyProjectSpendTime[name]; ok {
		return v
	}

	//fmt.Printf(">>>>>>>>>>>> %+v", name)
	return 5
}
