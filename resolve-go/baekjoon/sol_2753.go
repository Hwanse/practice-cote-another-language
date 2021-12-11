package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	year, _ := strconv.Atoi(sc.Text())
	fmt.Println(getLeafYearResult(year))
}

func getLeafYearResult(year int) int {
	if year % 4 == 0 && year % 100 != 0 {
		return 1
	}
	if year % 400 == 0 {
		return 1
	}
	return 0
}