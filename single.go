package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cart *shopCart

type shopCart struct {
	goods map[string]int
}

func (s *shopCart) add(name string, num int) {
	s.goods[name] += num
}

func (s *shopCart) print() {
	for k, v := range s.goods {
		print(k, " ", v)
	}
}

func getCart() *shopCart {
	if cart == nil {
		cart = &shopCart{
			make(map[string]int),
		}
	}
	return cart
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		parts := strings.Split(strings.TrimSpace(line), " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input format. Expected: <string> <number>", parts)
			continue
		}

		str := parts[0]
		numStr := parts[1]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", numStr)
			continue
		}

		getCart().add(str, num)
	}

	getCart().print()

	/*
		Apple 3
		Banana 2
		Orange 5
	*/
}
