package helpers

import (
	"regexp"
	"time"
	"unicode"
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

func CamelCaseToCapitalizeFirstWorldAndAddSpaces(str string) string {
	var newString []rune

	for index, ch := range str {
		if index == 0 {
			newString = append(newString, unicode.ToUpper(ch))
			continue
		}
		if unicode.IsUpper(ch) {
			newString = append(newString, ' ')
			newString = append(newString, unicode.ToLower(ch))
			continue
		}
		newString = append(newString, ch)
	}

	return string(newString)
}
