package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	caseCount := convertStringToInt(sc.Text())
	results := make([]int, 0, caseCount)

	for i := 0; i < caseCount; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")

		sum := convertStringToInt(inputs[0]) + convertStringToInt(inputs[1])
		results = append(results, sum)
	}

	for _, value := range results {
		fmt.Println(value)
	}
}

func convertStringToInt(numberStr string) int {
	number, _ := strconv.Atoi(numberStr)
	return number
}