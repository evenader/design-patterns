package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rectangle struct {
	color, length, width string
}

func (r rectangle) Clone() rectangle {
	return r
}

func (r *rectangle) Set(color, length, width string) {
	r.length = length
	r.color = color
	r.width = width
}

func (r rectangle) String() {
	fmt.Printf("Color: %s, Width: %s, Height: %s\n", r.color, r.length, r.width)
}

func NewRectangle(color, length, width string) rectangle {
	return rectangle{color, length, width}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	originObj := NewRectangle(" ", " ", " ")

	scanner.Scan()

	inputs := strings.Fields(scanner.Text())

	scanner.Scan()
	numStr := strings.Fields(scanner.Text())[0]
	num, _ := strconv.Atoi(numStr)

	for i := 0; i < num; i++ {
		newObj := originObj.Clone()
		newObj.Set(inputs[0], inputs[1], inputs[2])
		newObj.String()
	}

}

/*
看完人家的答案发现理解错了，顺带说一句这个设计模式好傻
package main

import (
	"fmt"
)

// 抽象原型类
type Prototype interface {
	clone() Prototype
	getDetails() string
}

// 具体矩形原型类
type RectanglePrototype struct {
	color  string
	width  int
	height int
}

// 构造方法
func NewRectanglePrototype(color string, width, height int) *RectanglePrototype {
	return &RectanglePrototype{
		color:  color,
		width:  width,
		height: height,
	}
}

// 实现 Prototype 接口的 clone 方法
func (r *RectanglePrototype) clone() Prototype {
	return &RectanglePrototype{
		color:  r.color,
		width:  r.width,
		height: r.height,
	}
}

// 获取矩形的详细信息
func (r *RectanglePrototype) getDetails() string {
	return fmt.Sprintf("Color: %s, Width: %d, Height: %d", r.color, r.width, r.height)
}

// 客户端程序
func main() {
	// 读取需要创建的矩形数量
	var N int
	fmt.Scan(&N)

	// 读取每个矩形的属性信息并创建矩形对象
	for i := 0; i < N; i++ {
		var color string
		var width, height int
		fmt.Scan(&color, &width, &height)

		// 创建原型对象
		originalRectangle := NewRectanglePrototype(color, width, height)

		// 克隆对象并输出详细信息
		clonedRectangle := originalRectangle.clone()
		fmt.Println(clonedRectangle.getDetails())
	}
}
*/
