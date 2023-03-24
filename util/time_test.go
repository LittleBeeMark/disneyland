package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLatestWeekDayList(t *testing.T) {

	got := GetLatestWeekDayList(time.Sunday, 6)
	fmt.Println("got :", got)

}
