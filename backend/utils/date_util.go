package utils

import (
	"time"
)

const (
	FORMAT      = "2006-01-02"
	BASE_FORMAT = "2006-01-02 15:04:05"
)

func GetNow() string {
	str := time.Now().Format(FORMAT)
	return str
}

func GetNowOfFormat(format string) string {
	format_time := time.Now().Format(format)
	return format_time
}

func GetDate(date string, format string) string {
	time, err := time.Parse(FORMAT, date)
	if err != nil {
		panic(err)
	}
	return time.Format(format)
}

func AddDays(date string, days int) string {
	time, err := time.Parse(FORMAT, date)
	if err != nil {
		panic(err)
	}
	return time.AddDate(0, 0, days).Format(FORMAT)
}

func AddDaysFromNow(days int) string {
	time := time.Now().AddDate(0, 0, days).Format(FORMAT)
	return time
}

func AddDaysOfPattern(date string, days int, format string) string {
	time, err := time.Parse(FORMAT, date)
	if err != nil {
		panic(err)
	}
	return time.AddDate(0, 0, days).Format(format)
}
