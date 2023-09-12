package util

import (
	"regexp"
	"runtime"
	"strings"
)

func ValidatePath(path string) bool {
	var incorrectDelimiter string
	if runtime.GOOS != "windows" {
		incorrectDelimiter = "\\"
	} else {
		incorrectDelimiter = "/"
	}
	if path != "" && !strings.Contains(path, incorrectDelimiter) {
		invalidCharsPattern := regexp.MustCompile(`[<>:"|?*]`)
		return !invalidCharsPattern.MatchString(path)
	}
	return false
}
