package util

import "regexp"

var (
	NAME_REGEX  = regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ\s\-']{1,100}$`)
	EMAIL_REGEX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)
