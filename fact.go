package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type blockFactory interface {
	produceBlock() blockProduct
}

type circleFactory struct{}

func (c *circleFactory) produceBlock() blockProduct {
	fmt.Println("Circle Block")
	return &circleBlock{}
}

type squareFactory struct{}

func (c *squareFactory) produceBlock() blockProduct {
	fmt.Println("Square Block")
	return &squareBlock{}
}

type blockProduct interface {
}

type circleBlock struct{}
type squareBlock struct{}

type blockFactorySystem struct {
}

func (bfs *blockFactorySystem) produceBlocks(bFactory blockFactory, num int) {
	for i := 0; i < num; i++ {
		bFactory.produceBlock()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bfs := blockFactorySystem{}

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		str, numStr := parts[0], parts[1]
		num, _ := strconv.Atoi(numStr)

		if str == "Square" {
			bfs.produceBlocks(&squareFactory{}, num)
		} else {
			bfs.produceBlocks(&circleFactory{}, num)
		}

	}

}
