package models

import (
	"strconv"
	"strings"
)

func lessIP(a, b string) bool {
	a_split := strings.Split(a, ".")
	b_split := strings.Split(b, ".")
	if len(a_split) != len(b_split) {
		return len(a_split) < len(b_split)
	}

	for i, v := range a_split {
		aInt, _ := strconv.Atoi(v)
		bInt, _ := strconv.Atoi(b_split[i])
		if aInt < bInt {
			return true
		}
		if aInt > bInt {
			return false
		}
	}
	return false
}
