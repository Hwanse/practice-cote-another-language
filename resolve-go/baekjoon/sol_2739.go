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

	n, _ := strconv.Atoi(sc.Text())

	for i := 1; i < 10; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}
}
