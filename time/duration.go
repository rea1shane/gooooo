package time

import (
	"regexp"
	"strconv"
	"time"
)

// FormatDuration 格式化 time.Duration，保留正负号，精度到秒
// 样例：-2d 15h 4m 5s
func FormatDuration(duration time.Duration) string {
	var symbol, day, hour, min, sec string

	durationString := duration.String()
	durationRegexp := regexp.MustCompile(`(-)?(\d*h)?(\d*m)?(\d*\.?\d*s)?`)
	durationParams := durationRegexp.FindStringSubmatch(durationString)

	symbol = durationParams[1]
	min = durationParams[3]

	h := durationParams[2]
	s := durationParams[4]

	hRegexp := regexp.MustCompile(`(\d*)h?`)
	hParams := hRegexp.FindStringSubmatch(h)
	hCount, _ := strconv.Atoi(hParams[1])
	if hCount >= 24 {
		day = strconv.Itoa(hCount/24) + "d"
	}
	hour = strconv.Itoa(hCount%24) + "h"

	sRegexp := regexp.MustCompile(`(\d*)\.?\d*s?`)
	sParams := sRegexp.FindStringSubmatch(s)
	sCount, _ := strconv.Atoi(sParams[1])
	sec = strconv.Itoa(sCount) + "s"

	if day != "" {
		day += " "
	}
	if hour != "" {
		hour += " "
	}
	if min != "" {
		min += " "
	}

	return symbol + day + hour + min + sec
}
