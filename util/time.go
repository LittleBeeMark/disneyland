package util

import (
	"fmt"
	"time"
)

func GetLatestWeekDayList(weekday time.Weekday, n int) []string {
	var (
		checkDay      = time.Now()
		checkNum      int
		resultDayList = make([]string, 0, n)
	)
	for checkNum < n {
		if checkDay.Weekday() == weekday {
			resultDayList = append(resultDayList, checkDay.Format("2006-01-02"))
			checkNum++
		}

		checkDay = checkDay.AddDate(0, 0, -1)
	}
	return resultDayList
}

func GetTimeStrHour(timeStr string) int {
	t, err := time.Parse("03:04 PM", timeStr)
	if err != nil {
		fmt.Println("解析时间字符串出错：", err)
		return 0
	}

	return t.Hour()
}

func GetStrInTimeZone(timeStr, start, end string) bool {
	t, err := time.Parse("03:04 PM", timeStr)
	if err != nil {
		fmt.Println("解析时间字符串出错：", err)
		return false
	}

	var s, e time.Time
	if start != "" {
		s, err = time.Parse("03:04 PM", start)
		if err != nil {
			fmt.Println("解析时间字符串出错：", err)
			return false
		}
	}

	if end != "" {
		e, err = time.Parse("03:04 PM", end)
		if err != nil {
			fmt.Println("解析时间字符串出错：", err)
			return false
		}
	}

	if start != "" && end != "" {
		return t.After(s) && t.Before(e)
	} else if start != "" {
		return t.After(s)
	} else if end != "" {
		return t.Before(e)
	}

	return false
}
