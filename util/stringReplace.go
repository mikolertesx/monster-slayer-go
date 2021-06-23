package util

import (
	"runtime"
	"strings"
)

func StringReplace(text string) string {
	if runtime.GOOS == "windows" {
		text = strings.Replace(text, "\r\n", "", -1)
	} else {
		text = strings.Replace(text, "\n", "", -1)
	}
	return text
}
