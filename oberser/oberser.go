package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Subject interface {
	AddObserver(ob Observers)
	NotifyObserver(nowTime int)
}

type timeSubject struct {
	curretnTime int
	observers   []Observers
}

func Init() Subject {
	return &timeSubject{
		curretnTime: 0,
		observers:   make([]Observers, 0, 0),
	}
}

func (t *timeSubject) AddObserver(ob Observers) {
	if t != nil {
		t.observers = append(t.observers, ob)
	}
}

func (t *timeSubject) NotifyObserver(nowTime int) {
	if t != nil {
		for _, ob := range t.observers {
			ob.Update(nowTime)
		}
	}
}

type Observers interface {
	Update(nowTime int)
}

type stuObserver struct {
	name string
}

func (s *stuObserver) Update(nowTime int) {
	fmt.Println(s.name + " " + fmt.Sprintf("%d", nowTime))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sub := Init()

	scanner.Scan()
	line := scanner.Text()
	times, _ := strconv.Atoi(strings.Fields(line)[0])

	for ; times >= 0; times-- {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		sub.AddObserver(&stuObserver{parts[0]})
	}

	scanner.Scan()
	line = scanner.Text()
	times, _ = strconv.Atoi(strings.Fields(line)[0])
	for time := 1; time < times; time++ {
		sub.NotifyObserver(time)
	}

}
