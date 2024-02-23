package util

import (
	"strings"
)

func Error(err error) (int, string) {
	e := ""
	s := 200

	if strings.TrimSpace(err.Error()) == "record not found" {
		e = "{}"
		s = 404
	} else if strings.Contains(err.Error(), "vl: ") {
		e = strings.Replace(err.Error(), "vl: ", "", 1)
		s = 400
	} else {
		e = string(err.Error())
		s = 500
	}
	return s, e
}
