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
	num1, _ := strconv.Atoi(inputs[0])
	num2, _ := strconv.Atoi(inputs[1])

	if num1 < num2 {
		fmt.Println("<")
	} else if num1 > num2 {
		fmt.Println(">")
	} else {
		fmt.Println("==")
	}
}
