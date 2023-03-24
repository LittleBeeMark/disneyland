package client

import (
	"LittleBeeMark/disneyland/model"
	"fmt"
	"testing"
	"time"
)

func TestGetEvery15minWaitTime(t *testing.T) {

	resMap, err := GetEvery15minWaitTime("2023-03-12")
	if err != nil {
		t.Log("err : ", err)
		return
	}
	for name, datas := range resMap {
		fmt.Printf("=========== name : %s ==============\n", model.GetProjectName(name))
		for _, data := range datas {
			fmt.Printf("timeStr : %s, waittime : %v \n", data.TimeStr, data.WaitTime)
		}
	}
}

func TestGetEvery5minWaitTime(t *testing.T) {

	resMap, err := GetEvery5minWaitTime("2023-03-12")
	if err != nil {
		t.Log("err : ", err)
		return
	}
	for name, datas := range resMap {
		fmt.Printf("=========== name : %s ==============\n", model.GetProjectName(name))
		for _, data := range datas {
			fmt.Printf("timeStr : %s, waittime : %v \n", data.TimeStr, data.WaitTime)
		}
	}
}
func TestGetEveryHourWaitTime(t *testing.T) {

	resMap, err := GetEveryHourWaitTime("2023-03-12")
	if err != nil {
		t.Log("err : ", err)
		return
	}
	for name, datas := range resMap {
		fmt.Printf("=========== name : %s ==============\n", model.GetProjectName(name))
		for _, data := range datas {
			fmt.Printf("timeStr : %s, waittime : %v \n", data.TimeStr, data.WaitTime)
		}
	}
}

func TestNewEvery5Data(t *testing.T) {
	t.Log("today is ", time.Now().Weekday() == time.Thursday)
}
