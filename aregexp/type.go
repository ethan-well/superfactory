package aregexp

import "regexp"

const (
	number = `^\d+$`
)

func IsNumber(v string) bool {
	regexp := regexp.MustCompile(number)
	return regexp.MatchString(v)
}
