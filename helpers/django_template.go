package helpers

import (
	"regexp"
	"time"
)

func GetFirstImageUrlFromString(str string) string {
	pattern := regexp.MustCompile("<img src=\"(.*?)\">")
	matches := pattern.FindAllStringSubmatch(str, -1)

	if matches == nil {
		return ""
	}
	return matches[0][0]
}

func FormatDate(date time.Time) string {
	return date.Format("02/01/2006 - 15:04:05")
}
