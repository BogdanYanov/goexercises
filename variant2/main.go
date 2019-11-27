package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n, res uint64
		x, y uint64 = 1, 1
		reader = bufio.NewReader(os.Stdin)
	)

	fmt.Println("====== Fibonacci ======")
	fmt.Print("n -> ")
	text, _ := reader.ReadString('\n')
	n, err := strconv.ParseUint(strings.TrimSuffix(text, "\n"), 10, 64)
	if err != nil {
		fmt.Errorf("can`t use this number")
		return
	}

	if n == 0 {
		res = 0
		goto Result
	}

	if n <= 2 {
		res = 1
		goto Result
	}


	for i := uint64(3); i <= n; i++ {
		res = x + y
		x, y = res, x
	}

Result:
		fmt.Println("Result ->", res)
		return
}
