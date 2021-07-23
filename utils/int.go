package utils

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

func UInt64To100StringFixed2(e uint64) string {
	return decimal.NewFromInt(int64(e)).Div(decimal.NewFromInt(100)).StringFixed(2)
}


func UInt64ToString(e uint64) string {
	return strconv.FormatUint(e, 10)
}

func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}