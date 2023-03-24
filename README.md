# 迪士尼攻略一键跑出

## 攻略思路
在早8：00 - 9：00 时间内（最早统计数据为8点，如果是早享票则没有7：30到八点的参考理论上差异不大可以尝试冲热门项目）
为拥挤程度最低的时间，利用该段时间去玩在正常游玩时间最为拥挤的项目为性价比最高的选择。

## 具体策略
1. 获取最近几个周的同星期所有项目不同时间点的排队时间数据
2. 进行同时间点的数据融合取均值
3. 计算同一个项目早8：00-9：10的等待时间均值为早享时间，计算9：30-17：30的等待时间均值为游玩时间
4. 获取早享时间和游玩时间的差值为节省时间
5. 因为排队及执行项目耗费的时间直接占据你的早享时间比例，所以排队时间加执行时间应与分数成反比，所以用节省时间除以排队时间与执行项目时间和得到分数
6. 分数最高则为早上优先级最高的项目，可以先冲。
7. 当然如果你买了早享卡我觉得是可以直接冲最热的项目的，我和女票就准备先冲小矮人矿工



## 执行结果
第1名：泡泡龙冲天赛车 || 早享时间：6.64 | 普通时间：49.93 | 节省时间：43.29 | 分数：0.77

第2名：小熊维尼历险记 || 早享时间：4.94 | 普通时间：32.65 | 节省时间：27.70 | 分数：0.62

第3名：喷气背包飞行器 || 早享时间：6.52 | 普通时间：38.05 | 节省时间：31.53 | 分数：0.57

第4名：弹簧狗团团转 || 早享时间：5.68 | 普通时间：33.06 | 节省时间：27.38 | 分数：0.57

第5名：彼得-潘的飞行 || 早享时间：5.04 | 普通时间：28.81 | 节省时间：23.77 | 分数：0.43

第6名：胡迪牛仔嘉年华 || 早享时间：7.25 | 普通时间：33.07 | 节省时间：25.83 | 分数：0.42

第7名：小飞象 || 早享时间：8.11 | 普通时间：33.73 | 节省时间：25.62 | 分数：0.37

第8名：古迹探索营的探索步行道 || 早享时间：5.04 | 普通时间：22.21 | 节省时间：17.17 | 分数：0.34

第9名：创极速光轮 || 早享时间：9.06 | 普通时间：31.38 | 节省时间：22.32 | 分数：0.29

第10名：加勒比海盗-沉落宝藏之战 || 早享时间：5.18 | 普通时间：19.32 | 节省时间：14.14 | 分数：0.26

第11名：旋转蜂蜜罐 || 早享时间：5.00 | 普通时间：15.94 | 节省时间：10.94 | 分数：0.25

第12名：雷鸣山漂流 || 早享时间：10.09 | 普通时间：30.56 | 节省时间：20.46 | 分数：0.23

第13名：七个小矮人矿工 || 早享时间：27.94 | 普通时间：62.70 | 节省时间：34.76 | 分数：0.15

第14名：巴斯光年星际营救 || 早享时间：5.00 | 普通时间：11.14 | 节省时间：6.14 | 分数：0.14

第15名：翱翔地平线 || 早享时间：43.76 | 普通时间：67.55 | 节省时间：23.79 | 分数：0.07 

## 使用方式

直接在Idea中执行main函数，可以调整myProjects为你喜欢的项目，也可以调整GetAllDataOfProjectByDateFusion的第一个参数util.GetLatestWeekDayList(time.Sunday, 6)
为你要去的星期，我们是星期日去，所以这个用的星期日.util.GetLatestWeekDayList(time.Sunday, 6)第二个参数为取最近几个同星期的数据做融合均值，我取了最近6个周日的数据来做平均
当然建议不要取太多这个网站获取数据的速度一言难进。
这个函数前面是过滤数据用的func(m map[string][]*model.WaitTimeData) {
business.FilterTimeAfter(19, m)}
是过滤数据用的,我里面使用的过滤数据为过滤调晚上七点以后的数据，你可以根据你的需要调整

每天不同项目时间点等待时间是从[这个网站](https://www.thrill-data.com/waits/park/dis/shanghai-disneyland/)爬的。

```go
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

	got, err := business.GetAllDataOfProjectByDateFusion(util.GetLatestWeekDayList(time.Sunday, 6),
		func(m map[string][]*model.WaitTimeData) {
			business.FilterTimeAfter(19, m)
		})
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	sts := business.GenerateStrategy(got, myProjects)
	for i, st := range sts {
		fmt.Println(st.Desc(i))
	}
}



```


## 注意
由于爬取速度很慢所以取得数据越多越慢，我上面这个例子就差不多要10秒左右。