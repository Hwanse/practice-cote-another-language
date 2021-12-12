package programmers

import (
	"regexp"
	"strconv"
	"strings"
)

var numberStrings = map[string]string {
	"zero":"0",
	"one":"1",
	"two":"2",
	"three":"3",
	"four":"4",
	"five":"5",
	"six":"6",
	"seven":"7",
	"eight":"8",
	"nine":"9",
}

func solution(s string) int {
	tempString := s

	for ; isContainStrings(tempString); {
		for key, value := range numberStrings {
			tempString = strings.ReplaceAll(tempString, key, value)
		}
	}

	result, _ := strconv.Atoi(tempString)
	return result
}

/*
 풀이2 strings.NewReplacer를 활용한 방법
 */
func solution2(s string) int {
	r := strings.NewReplacer(
		"zero", "0",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	result, _ := strconv.Atoi(r.Replace(s))
	return result
}

func isContainStrings(numberString string) bool{
	contain, _ := regexp.MatchString("^[0-9]+$", numberString)
	return !contain
}