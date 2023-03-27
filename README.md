
# 程序员如何为女朋友做迪士尼攻略

> 身为一个程序员，我们每天大多时间对着计算机没有交际应酬，更不懂得浪漫与体贴。我能为女朋友做点什么呢？我是不是能用我的专业来帮女票做点事情。刚好女票生日计划去迪士尼玩，做攻略的重任就落到了我的身上。我是不是可以利用数据和逻辑科学的帮助女友获取一份攻略呢？
>



## 前言

我在苦思冥想该怎么做攻略的时候看了很多小红书，视频，还有直播，我发现他们都会给出游玩攻略。但很可惜他们清一色的就是告诉你该以什么顺序去玩却不告诉你形成攻略的具体逻辑是什么，他们形成攻略的说辞基本就是靠经验，这对于我一个逻辑至上的程序员来说是不能忍受的。所有的up主都千篇一律的就是第一个冲翱翔地平线，第二个小矮人矿车。。。

我知道他们通过经验形成的逻辑点就是在最早的时间冲最火的项目，这样在主观感受上确实是可以节省更多时间。但最火的项目在早上不也一样是最火的？不也一样等的时间比其他的久，在所有项目都接近排队时长最短的黄金早晨时间项目排队耗费时间也是能否玩更多项目的关键。

我觉得可以根据多日的单项目不同时间点平均等待时间辅以一些算法来获取一个项目游玩优先级权重。



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
7. 接下来我还准备了**单个项目不同时间点等待时间排名 **帮助你安排除了早享黄金时间外的游玩攻略
8. 当然如果你买了早享卡我觉得是可以直接冲最热的项目的，我和女票就准备先冲小矮人矿工



## 执行结果

```shell

========================= 迪士尼策略优先级排行榜 =========================
第1名：泡泡龙冲天赛车 || 早享时间：8.09 | 普通时间：64.89 | 节省时间：56.80 | 分数：0.84 
第2名：小熊维尼历险记 || 早享时间：5.23 | 普通时间：41.26 | 节省时间：36.03 | 分数：0.77 
第3名：弹簧狗团团转 || 早享时间：6.14 | 普通时间：44.13 | 节省时间：37.99 | 分数：0.73 
第4名：彼得-潘的飞行 || 早享时间：5.04 | 普通时间：38.97 | 节省时间：33.93 | 分数：0.61 
第5名：喷气背包飞行器 || 早享时间：7.99 | 普通时间：48.09 | 节省时间：40.10 | 分数：0.60 
第6名：古迹探索营的探索步行道 || 早享时间：5.08 | 普通时间：30.77 | 节省时间：25.69 | 分数：0.51 
第7名：胡迪牛仔嘉年华 || 早享时间：9.32 | 普通时间：45.86 | 节省时间：36.54 | 分数：0.47 
第8名：加勒比海盗-沉落宝藏之战 || 早享时间：5.27 | 普通时间：28.79 | 节省时间：23.52 | 分数：0.43 
第9名：旋转蜂蜜罐 || 早享时间：5.00 | 普通时间：22.95 | 节省时间：17.95 | 分数：0.42 
第10名：小飞象 || 早享时间：11.04 | 普通时间：48.12 | 节省时间：37.09 | 分数：0.40 
第11名：创极速光轮 || 早享时间：12.77 | 普通时间：44.78 | 节省时间：32.01 | 分数：0.30 
第12名：雷鸣山漂流 || 早享时间：14.88 | 普通时间：47.39 | 节省时间：32.51 | 分数：0.25 
第13名：巴斯光年星际营救 || 早享时间：5.00 | 普通时间：15.70 | 节省时间：10.70 | 分数：0.24 
第14名：七个小矮人矿工 || 早享时间：47.74 | 普通时间：78.91 | 节省时间：31.17 | 分数：0.08 
第15名：翱翔地平线 || 早享时间：65.22 | 普通时间：88.34 | 节省时间：23.12 | 分数：0.04 



========================= 单个项目不同时间点等待时间排名 =========================
2023-03-26,2023-03-19,2023-03-12,2023-03-05,2023-02-26,2023-02-19 项目等待时间点排名: 

翱翔地平线: 
 1. 08:35 AM : 55 
 2. 06:55 PM : 55.9375 
 3. 06:35 PM : 56.875 
 4. 06:40 PM : 56.875 
 5. 06:45 PM : 56.875 
 6. 06:50 PM : 56.875 
 7. 06:30 PM : 57.34375 
 8. 08:30 AM : 58 
 9. 08:40 AM : 61.25 
 10. 06:25 PM : 62.34375 
 11. 06:20 PM : 63.28125 
 12. 08:00 AM : 64.375 
 13. 08:45 AM : 65.46875 
 14. 06:15 PM : 67.03125 
 15. 06:10 PM : 68.4375 
 16. 06:05 PM : 68.90625 
 17. 08:50 AM : 69.21875 
 18. 08:55 AM : 70.9375 
 19. 09:00 AM : 70.9375 
 20. 09:05 AM : 70.9375 
 21. 01:00 PM : 74.53125 
 22. 01:05 PM : 74.53125 
 23. 01:10 PM : 75.46875 
 24. 02:25 PM : 75.9375 
 25. Average : 76.21875 
 26. 12:45 PM : 76.40625 
 27. 12:50 PM : 76.40625 
 28. 12:55 PM : 76.40625 
 29. 06:00 PM : 76.40625 
 30. 12:15 PM : 76.875 
 31. 12:20 PM : 76.875 
 32. 12:25 PM : 76.875 
 33. 12:30 PM : 76.875 
 34. 04:25 PM : 76.875 
 35. 04:30 PM : 76.875 
 36. 12:35 PM : 77.34375 
 37. 12:40 PM : 77.34375 
 38. 04:20 PM : 77.34375 
 39. 01:20 PM : 78.75 
 40. 02:30 PM : 78.75 
 41. 04:10 PM : 79.21875 
 42. 01:25 PM : 79.6875 
 43. 02:15 PM : 79.6875 
 44. 02:20 PM : 79.6875 
 45. 04:35 PM : 79.6875 
 46. 02:35 PM : 80.15625 
 47. 10:40 AM : 80.625 
 48. 10:45 AM : 80.625 
 49. 01:15 PM : 80.625 
 50. 02:40 PM : 80.625 

七个小矮人矿工: 
 1. 08:35 AM : 43.125 
 2. 08:30 AM : 44.59375 
 3. 08:40 AM : 45.46875 
 4. 08:00 AM : 46.96875 
 5. 08:45 AM : 48.28125 
 6. 08:55 AM : 48.75 
 7. 08:50 AM : 49.21875 
 8. 09:10 AM : 50.9375 
 9. 09:00 AM : 51.25 
 10. 09:05 AM : 51.25 
 11. 06:55 PM : 53.4375 
 12. 06:50 PM : 55.9375 
 13. 09:15 AM : 56.40625 
 14. 06:35 PM : 56.5625 
 15. 06:40 PM : 56.5625 
 16. 06:45 PM : 57.8125 
 17. 09:20 AM : 59.21875 
 18. 05:55 PM : 62.34375 
 19. 12:45 PM : 62.5 
 20. 09:25 AM : 62.96875 
 21. 12:25 PM : 62.96875 
 22. 12:40 PM : 62.96875 
 23. 12:30 PM : 63.4375 
 24. 12:35 PM : 63.4375 
 25. 09:30 AM : 63.90625 
 26. 09:35 AM : 63.90625 
 27. 09:40 AM : 64.21875 
 28. 09:45 AM : 64.21875 
 29. 09:50 AM : 64.21875 
 30. 06:20 PM : 65.46875 
 31. 06:25 PM : 65.46875 
 32. 06:30 PM : 65.46875 
 33. 06:05 PM : 65.78125 
 34. 06:10 PM : 65.78125 
 35. 06:15 PM : 65.78125 
 36. 05:40 PM : 65.9375 
 37. 05:45 PM : 65.9375 
 38. 05:50 PM : 66.09375 
 39. 06:00 PM : 66.09375 
 40. Average : 66.625 
 41. 05:35 PM : 66.875 
 42. 12:15 PM : 67.96875 
 43. 12:20 PM : 67.96875 
 44. 12:50 PM : 70 
 45. 12:55 PM : 70 
 46. 01:00 PM : 70.9375 
 47. 12:05 PM : 72.96875 
 48. 12:10 PM : 72.96875 
 49. 11:05 AM : 73.59375 
 50. 11:10 AM : 73.59375 

创极速光轮: 
 1. 08:30 AM : 9.6875 
 2. 08:35 AM : 9.6875 
 3. 08:40 AM : 9.6875 
 4. 08:00 AM : 10.0625 
 5. 08:45 AM : 10.3125 
 6. 08:50 AM : 10.3125 
 7. 08:55 AM : 10.3125 
 8. 09:00 AM : 18.4375 
 9. 06:35 PM : 20.3125 
 10. 06:30 PM : 22.8125 
 11. 09:05 AM : 23.75 
 12. 06:55 PM : 24.0625 
 13. 09:10 AM : 24.6875 
 14. 06:40 PM : 25.3125 
 15. 06:45 PM : 25.3125 
 16. 06:50 PM : 25.3125 
 17. 12:30 PM : 26.25 
 18. 12:35 PM : 26.25 
 19. 12:40 PM : 26.25 
 20. 06:00 PM : 29.0625 
 21. 06:05 PM : 29.0625 
 22. 06:10 PM : 29.0625 
 23. 05:55 PM : 29.6875 
 24. 12:25 PM : 31.5625 
 25. 12:20 PM : 31.875 
 26. 05:50 PM : 31.875 
 27. 06:15 PM : 32.1875 
 28. 12:15 PM : 32.5 
 29. 05:25 PM : 33.125 
 30. 05:35 PM : 33.125 
 31. 05:40 PM : 33.4375 
 32. 05:45 PM : 33.4375 
 33. 05:30 PM : 33.75 
 34. 12:10 PM : 34.0625 
 35. 06:25 PM : 34.0625 
 36. 02:15 PM : 34.375 
 37. 05:20 PM : 34.375 
 38. 02:10 PM : 34.6875 
 39. 06:20 PM : 34.6875 
 40. 02:05 PM : 35 
 41. 04:05 PM : 35 
 42. 04:10 PM : 35 
 43. 12:05 PM : 35.3125 
 44. 03:55 PM : 35.3125 
 45. 04:00 PM : 35.3125 
 46. Average : 35.3125 
 47. 12:45 PM : 35.625 
 48. 02:00 PM : 38.125 
 49. 05:10 PM : 39.375 
 50. 05:15 PM : 39.375 

雷鸣山漂流: 
 1. 08:35 AM : 7.5 
 2. 08:30 AM : 9 
 3. 08:40 AM : 10 
 4. 08:45 AM : 11.25 
 5. 08:00 AM : 12.25 
 6. 08:50 AM : 13.75 
 7. 06:55 PM : 13.75 
 8. 08:55 AM : 18.75 
 9. 09:00 AM : 18.75 
 10. 06:50 PM : 18.75 
 11. 05:55 PM : 22.5 
 12. 06:00 PM : 22.5 
 13. 06:35 PM : 22.5 
 14. 06:40 PM : 22.5 
 15. 06:45 PM : 23.75 
 16. 05:50 PM : 25 
 17. 06:10 PM : 25 
 18. 06:15 PM : 25 
 19. 06:20 PM : 25 
 20. 06:25 PM : 25 
 21. 06:30 PM : 25 
 22. 05:45 PM : 26.25 
 23. 06:05 PM : 27.5 
 24. 09:05 AM : 30 
 25. 09:10 AM : 35 
 26. 09:15 AM : 35 
 27. 09:20 AM : 35 
 28. 09:25 AM : 35 
 29. 09:30 AM : 37.5 
 30. 09:40 AM : 37.5 
 31. 09:45 AM : 37.5 
 32. Average : 38.25 
 33. 09:35 AM : 38.75 
 34. 05:35 PM : 38.75 
 35. 05:40 PM : 38.75 
 36. 09:50 AM : 40 
 37. 01:05 PM : 40 
 38. 11:50 AM : 41.25 
 39. 11:55 AM : 41.25 
 40. 12:00 PM : 41.25 
 41. 12:05 PM : 41.25 
 42. 01:00 PM : 41.25 
 43. 02:15 PM : 41.25 
 44. 02:20 PM : 41.25 
 45. 05:20 PM : 41.25 
 46. 05:25 PM : 41.25 
 47. 05:30 PM : 41.25 
 48. 09:55 AM : 42.5 
 49. 12:50 PM : 42.5 
 50. 12:55 PM : 42.5 

.....
```



## 使用方式

直接在Idea中执行main函数。
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

```

### 第一步：获取数据源

```go
got, err := business.GetAllDataOfProjectByDateFusion(dateList,
			func(m map[string][]*model.WaitTimeData) {
				business.FilterTimeAfter(19, m)
			})
```

也可以调整GetAllDataOfProjectByDateFusion的第一个参数dateList，dateList 可以使用我封装的一个函数 util.GetLatestWeekDayList 来获取：

```go
dateList := util.GetLatestWeekDayList(time.Sunday, 6)
```

该函数的两个参数分别为你要去的星期，和你要获取多少个同星期的日子进行数据融合平均值。

我们是星期日去，所以第一个参数用的星期日 time.Sunday，第二个参数我取了最近6个周日的数据来做平均值。当然建议不要取太多这个网站获取数据的速度一言难进。

数据源我们取到后就开始获取策略数据了。



### 第二步：迪士尼策略优先级排行榜获取

```go
sts := business.GenerateStrategy(got, myProjects)
```

myProjuects 是你想要去的玩的项目。可以自行调整，我已经吧中文名称对应注释写上了（有几个确实不知道是啥，我就直译了，你可以再改改）。

这个函数签名参数是过滤数据用的func(m map[string][]*model.WaitTimeData) {business.FilterTimeAfter(19, m)}，这里我放入了一个封装 business.FilterTimeAfter 这个函数的逻辑是过滤掉晚上七点后的数据。你可以根据你的需要调整

**注意：**这里必须注意的是我生成的这个策略最早时间只有八点，而买了早享票的你7点15估计就可以入园了。你要是牛逼七点之前就到了早享票一开门就冲进去的话直接冲小矮人或是地平线，到的很早就先冲地平线再冲小矮人，因为可以想象你如果是第一个入园的那么所有项目对你来说都不需要排队，那一定是冲最火的，而且你可以额基本5-10分钟完成这个这个最火的项目丝毫不影响下一个项目的进行逻辑。（我觉得那些up一定是假设你是第一个入园的然后才生成了他们的策略，基本不适用于你八点以后甚至七点半以后入园的你）



### 第三步：生成单个项目不同时间点等待时间排名

接下来是生成**单个项目不同时间点等待时间排名**

```go
sortListGot, err := business.DescSortListData(got, dateList, myProjects, 50)
```

这里的datelist，和 myProjects 参数和刚刚介绍的两个函数一致，都介绍过了，最后一个参数50，就是获取多少个时间点，我取了50就是获取50个时间点。

这个排行榜也是我认为最有用的，因为第一步生成的策略只是早享时间的游玩顺序可以根据他来策划其他时间的游玩顺序，你只要找这个项目排队时间最短的时刻去玩就好了，当然你要结合当前的等待数据，不要完全依赖这个排名榜单。



每天不同项目时间点等待时间是从[这个网站](https://www.thrill-data.com/waits/park/dis/shanghai-disneyland/)爬的。


## 注意
由于爬取速度很慢所以取得数据越多越慢，我上面这个例子取了六天的数据均值就差不多要30秒左右。需要耐心一些。

