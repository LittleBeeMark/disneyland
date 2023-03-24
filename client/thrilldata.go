package client

import (
	"LittleBeeMark/disneyland/model"
	"LittleBeeMark/disneyland/util"
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"strconv"
)

const GetWaitTimeUrlTemplate = "https://www.thrill-data.com/waits/graph/quick/parkheat?id=%s&dateStart=%s&tag=%s"

type ThrillData struct {
	ID          string
	DateStart   string
	Tag         string
	UrlTemplate string
}

type resJson struct {
	Plot1 string
	Title string
}

func NewThrillData(id string, dateStart, tag, urlTempl string) *ThrillData {
	return &ThrillData{
		id,
		dateStart,
		tag,
		urlTempl,
	}
}

func NewEvery15Data(date string) *ThrillData {
	return NewThrillData("54", date, "min", GetWaitTimeUrlTemplate)
}
func NewEvery5Data(date string) *ThrillData {
	return NewThrillData("54", date, "five", GetWaitTimeUrlTemplate)
}
func NewEveryHourData(date string) *ThrillData {
	return NewThrillData("54", date, "hour", GetWaitTimeUrlTemplate)
}

func (td *ThrillData) getUrl() string {
	return fmt.Sprintf(td.UrlTemplate, td.ID, td.DateStart, td.Tag)
}

func GetEvery15minWaitTime(date string) (map[string][]*model.WaitTimeData, error) {
	return NewEvery15Data(date).WaitTime()
}

func GetEvery5minWaitTime(date string) (map[string][]*model.WaitTimeData, error) {
	return NewEvery5Data(date).WaitTime()
}

func GetEveryHourWaitTime(date string) (map[string][]*model.WaitTimeData, error) {
	return NewEveryHourData(date).WaitTime()
}

func (td *ThrillData) WaitTime() (map[string][]*model.WaitTimeData, error) {
	res, err := util.HttpGet(td.getUrl(), nil)
	if err != nil {
		return nil, err
	}
	resultJson := &resJson{}

	err = json.Unmarshal(res, resultJson)
	if err != nil {
		return nil, err
	}
	//fmt.Println("plot1:", resultJson.Plot1)

	reg1, err := regexp2.Compile(`(?<=,"y":).+(?=,"ygap":)`, 0)
	if err != nil {
		return nil, err
	}
	match1, err := reg1.FindStringMatch(resultJson.Plot1)
	if err != nil {
		return nil, err
	}
	//fmt.Println("match 1: ", match1.String())
	nameSlice := []string{}
	err = json.Unmarshal([]byte(match1.String()), &nameSlice)
	if err != nil {
		return nil, err
	}
	//for _, name := range nameSlice {
	//	fmt.Println("name slice", name)
	//}

	reg2, err := regexp2.Compile(`(?<=,"x":).+(?=,"xgap":)`, 0)
	if err != nil {
		return nil, err
	}
	match2, err := reg2.FindStringMatch(resultJson.Plot1)
	if err != nil {
		return nil, err
	}
	timeSlice := []string{}
	err = json.Unmarshal([]byte(match2.String()), &timeSlice)
	if err != nil {
		return nil, err
	}

	//fmt.Println("match 2: ", match2.String())

	var reg3 *regexp2.Regexp
	if td.Tag == "five" {
		reg3, err = regexp2.Compile(`(?<="ygap":0,"z":).+(?=,"zmax":)`, 0)
	} else {
		reg3, err = regexp2.Compile(`(?<=true,"text":).+(?=,"textfont":)`, 0)
	}
	if err != nil {
		return nil, err
	}
	match3, err := reg3.FindStringMatch(resultJson.Plot1)
	if err != nil {
		return nil, err
	}
	//fmt.Println("match 3: ", match3.String())

	waitSlice := [][]interface{}{}
	err = json.Unmarshal([]byte(match3.String()), &waitSlice)
	if err != nil {
		return nil, err
	}

	resMap := make(map[string][]*model.WaitTimeData, len(nameSlice))
	for i, name := range nameSlice {
		for j, d := range timeSlice {
			waitTime := waitSlice[i][j]
			resData := &model.WaitTimeData{
				Name:    name,
				TimeStr: d,
			}

			var waitTimeData float64 = -1
			if waitTime != nil {
				if v, ok := waitTime.(float64); ok {
					waitTimeData = v
				} else if v, ok := waitTime.(string); ok {
					if v != "" {
						waitTimeData, err = strconv.ParseFloat(v, 64)
					}
				}
			}

			resData.WaitTime = waitTimeData
			resMap[name] = append(resMap[name], resData)

			//fmt.Printf("名字： %s,时间 %s 等待：%s \n", name, d, waitTime)
		}
	}

	//fmt.Println("match 3: ", match3.String())
	//fmt.Println("res : ", resultJson.Plot1)
	return resMap, nil
}

func toJSON(v interface{}) (js string) {
	b, _ := json.Marshal(v)
	return string(b)
}
