package service

import (
	"encoding/json"
	"fmt"
	"regexp"
	"report/goreport/backend/utils"
	"strconv"
)

type SearchParam struct {
}

func Generate(uid string) {
	report := GetReportByUid(uid)
	if report == nil {
		panic("未找到该报表")
	}
	var reportOption ReportOptions
	json.Unmarshal([]byte(report.Options), &reportOption)

}

func GetFromParameters(params map[string]string, dateRange int) *ReportParameter {
	c, ok := params["startTime"]
	if ok {
		fmt.Println(c)
	}
	return nil
}

func GetQueryColumn(uid string) []ReportQueryParameter {
	report := GetReportByUid(uid)
	sqlText := report.SqlText
	fmt.Println(sqlText)
	params := utils.FindAll(sqlText, `\$\{.*?\}`)
	size := len(params)
	var queryParams map[string]string = make(map[string]string, size)
	for _, param := range params {
		key := utils.ReplaceAll(string(param), `[\\$\\{\\}]`, "")
		queryParams[key] = ""
	}
	var reprtQueryParameter []ReportQueryParameter
	if _, ok := queryParams["startTime"]; ok {
		options := parseOptions(report.Options)
		dategap, _ := strconv.Atoi(options.DataRange)
		buildinParams := setBuildInParams(queryParams, dategap)
		reprtQueryParameter =
			append(reprtQueryParameter, ReportQueryParameter{DefaultValue: buildinParams["startTime"], FormElement: "date", Name: "startTime", Text: "开始时间", DataType: "date"})
		reprtQueryParameter =
			append(reprtQueryParameter, ReportQueryParameter{DefaultValue: buildinParams["endTime"], FormElement: "date", Name: "endTime", Text: "结束时间", DataType: "date"})
	}

	customQueryParams := parseQueryParams(report.QueryParams)

	if len(customQueryParams) != 0 {
		for _, p := range customQueryParams {
			reprtQueryParameter = append(reprtQueryParameter, p)
		}
	}
	return reprtQueryParameter
}

func getSqlParams(sqlText string) []string {
	regex := regexp.MustCompile(`\$\{.*?\}`)
	findStr := regex.FindAll([]byte(sqlText), -1)
	var strs []string
	for i, str := range findStr {
		strs[i] = string(str)
	}
	return strs
}

func setBuildInParams(params map[string]string, dateRange int) map[string]string {
	if dateRange-1 < 0 {
		dateRange = 0
	} else {
		dateRange = dateRange - 1
	}
	params["startTime"] = utils.AddDaysFromNow(-dateRange)
	params["endTime"] = utils.GetNow()
	return params
}
