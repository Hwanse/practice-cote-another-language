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
	inputs := sc.Text()

	split := strings.Split(inputs, " ")
	num1, _ := strconv.ParseFloat(split[0], 32)
	num2, _ := strconv.ParseFloat(split[1], 32)

	fmt.Println(num1 / num2)
}
