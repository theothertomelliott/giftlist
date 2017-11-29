package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func RenderMoney(value int64) string {
	quotient := value / 100 // integer division, decimals are truncated
	remainder := value % 100
	sign := ""
	if quotient < 0 {
		quotient *= -1
		sign = "-"
	}
	if remainder < 0 {
		remainder *= -1
		sign = "-"
	}
	return fmt.Sprintf("%s%d.%02d", sign, quotient, remainder)
}

func ParseMoney(value string) (int64, error) {

	sep := strings.Split(value, ".")
	if len(sep) == 0 || len(sep) > 2 {
		return 0, errors.New("price is not a decimal")
	}
	high, err := strconv.Atoi(sep[0])
	if err != nil {
		return 0, errors.WithStack(err)
	}
	low := 0
	if len(sep) == 2 {
		low, err = strconv.Atoi(sep[1])
		if err != nil {
			return 0, errors.WithStack(err)
		}
		if len(sep[1]) == 1 {
			low *= 10
		}
	}
	if low >= 100 {
		return 0, errors.New("price is not correctly formatted")
	}
	total := high*100 + low
	if value[0] == '-' {
		total = high*100 - low
	}
	return int64(total), nil
}
