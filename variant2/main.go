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
		i         uint64 = 3
		n, res    uint64
		cur, prev uint64 = 1, 1
		reader           = bufio.NewReader(os.Stdin)
	)

	fmt.Println("====== Fibonacci ======")
	fmt.Print("n -> ")

	text, _ := reader.ReadString('\n')
	n, err := strconv.ParseUint(strings.TrimSuffix(text, "\n"), 10, 64)
	if err != nil {
		fmt.Println(fmt.Errorf("can`t use this number"))
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

Loop:
	if i > n {
		goto Result
	}

	res = cur + prev
	cur, prev = res, cur
	i++
	goto Loop

Result:
	fmt.Println("Result ->", res)
}
