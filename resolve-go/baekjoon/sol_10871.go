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

	inputs := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(inputs[0])
	x, _ := strconv.Atoi(inputs[1])
	result := make([]int, 0, n)

	sc.Scan()
	numbersStr := sc.Text()

	for _, str := range strings.Split(numbersStr, " ") {
		number, _ := strconv.Atoi(str)
		if number < x {
			result = append(result, number)
		}
	}

	for _, value := range result {
		fmt.Print(value, " ")
	}
}
