package sliceutil

import (
	"github.com/ItsWewin/superfactory/aerror"
	"regexp"
	"strconv"
	"strings"
)

const (
	int64ArrStringFormat = `^\[((\d,)*\d*)\]$`
)

func ArrStringToInt64Arr(str string) ([]int64, aerror.Error) {
	var arr []int64
	if str == "" {
		return arr, nil
	}

	matched, err := regexp.MatchString(int64ArrStringFormat, str)
	if err != nil {
		return arr, nil
	}

	if !matched {
		return arr, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "expected format: %s, get: %s", int64ArrStringFormat, str)
	}

	str = strings.TrimLeft(str, "[")
	str = strings.TrimRight(str, "]")

	if len(str) == 0 {
		return arr, nil
	}

	arrStr := strings.Split(str, ",")
	if len(arrStr) == 0 {
		return arr, nil
	}

	for _, s := range arrStr {
		a, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "sting parse int failed")
		}
		arr = append(arr, a)
	}

	return arr, nil
}
