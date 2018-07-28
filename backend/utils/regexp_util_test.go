package utils

import (
	"testing"
)

func Test_FindAll(t *testing.T) {
	regexpStr := "\\$\\{.*?\\}"
	str := `select * from '${dt}'`

	strs := FindAll(str, regexpStr)
	rst := string(strs[0])

	if rst != "${dt}" {
		t.Error("find all error!")
	}
}

func Test_RelaceAll(t *testing.T) {
	result := "${dt}"
	regexpStr := "[\\$\\{\\}]"
	rst := string(ReplaceAll(result, regexpStr, ""))
	if rst != "dt" {
		t.Error("find all error!")
	}
}
