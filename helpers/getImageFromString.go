package helpers

import "regexp"

func GetImageFromString(str string) string {
	pattern := regexp.MustCompile("<img src=\"(?s)(.*)\">")
	firstMatchSubstring := pattern.FindString(str)
	return firstMatchSubstring
}
