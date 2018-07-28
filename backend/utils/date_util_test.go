package utils

import (
	"fmt"
	"testing"
)

const format = "2006/01/02"

func Test_GetNow(t *testing.T) {
	fmt.Println(GetNow())
}

func Test_GetNowOfFormat(t *testing.T) {

	fmt.Println(GetNowOfFormat(format))
}

func Test_GetDate(t *testing.T) {
	date := GetDate("2018-01-01", format)
	if date != "2018/01/01" {
		t.Error("test parse date format error")
	}
}

func Test_AddDays(t *testing.T) {
	tomorrow := AddDays("2018-01-01", 1)
	yesterday := AddDays("2018-01-01", -1)

	if tomorrow != "2018-01-02" || yesterday != "2017-12-31" {
		t.Error("test add date error")
	}
}

func Test_AddDaysOfPattern(t *testing.T) {
	tomorrow := AddDaysOfPattern("2018-01-01", 1, format)
	yesterday := AddDaysOfPattern("2018-01-01", -1, format)

	if tomorrow != "2018/01/02" || yesterday != "2017/12/31" {
		t.Error("test add date error")
	}
}
