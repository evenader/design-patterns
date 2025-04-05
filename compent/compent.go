package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type compent interface {
	Display(indent string)
}

type Department struct {
	name       string
	subCompent []compent
}

func NewDepartment(name string) *Department {
	return &Department{name: name, subCompent: []compent{}}
}

func (d *Department) GetName() string {
	return d.name
}

func (d *Department) Display(indent string) {
	fmt.Println(indent + d.name)
	for _, sub := range d.subCompent {
		sub.Display(indent + "  ")
	}
}

func (d *Department) Add(com compent) {
	d.subCompent = append(d.subCompent, com)
}

type Employee struct {
	name string
}

func NewEmployee(name string) *Employee {
	return &Employee{name: name}
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) Display(indent string) {
	fmt.Println(indent + e.name)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	companyName := scanner.Text()
	scanner.Scan()
	n := 0
	fmt.Sscanf(scanner.Text(), "%d", &n)

	company := NewDepartment(companyName)
	temp := &Department{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		if len(parts) < 2 {
			continue
		}
		if parts[0] == "D" {
			depart := NewDepartment(parts[1])
			company.Add(depart)
			temp = depart
		} else if parts[0] == "E" {
			temp.Add(NewEmployee(parts[1]))
		}
	}

	fmt.Println("Company Structure:")
	company.Display("")
}
