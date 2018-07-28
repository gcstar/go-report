package utils

import (
	"regexp"
)

func FindAll(str string, regexpStr string) [][]byte {
	r := regexp.MustCompile(regexpStr)
	return r.FindAll([]byte(str), -1)
}

func ReplaceAll(str string, regexpStr string, replaceStr string) string {
	r := regexp.MustCompile(regexpStr)
	return string(r.ReplaceAll([]byte(str), []byte(replaceStr)))
}
