package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const START_HOUR = 0
const END_HOUR = 24
const END_MINUTE = 60

const EARLY_MINUTE = 45

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	hour, _ := strconv.Atoi(inputs[0])
	minute, _ := strconv.Atoi(inputs[1])

	resultHour := hour
	resultMinute := minute

	if minute < EARLY_MINUTE {
		resultMinute = minute + END_MINUTE - EARLY_MINUTE

		if hour == START_HOUR {
			resultHour = END_HOUR - 1
		} else {
			resultHour = hour - 1
		}
	} else {
		resultMinute = minute - EARLY_MINUTE
	}

	fmt.Println(resultHour, resultMinute)
}
