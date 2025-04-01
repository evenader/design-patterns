package main

import (
	"bufio"
	"os"
	"strings"
)

type director struct{}

func (d *director) direct(b bikeBuilder) {
	b.buildFrame()
	b.buildTires()
}

type bikeBuilder interface {
	frameBuilder
	tiresBuilder
}

type frameBuilder interface {
	buildFrame()
}

type tiresBuilder interface {
	buildTires()
}

type roadBikeBuilder struct{}

func (r *roadBikeBuilder) buildFrame() {
	print("Carbon Frame ")
}
func (r *roadBikeBuilder) buildTires() {
	println("Slim Tries")
}

type moutainBikeBuilder struct{}

func (r *moutainBikeBuilder) buildFrame() {
	print("Aluminum Frame ")
}
func (r *moutainBikeBuilder) buildTires() {
	println("Knobby Tires")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	d := director{}

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		bikeName := strings.Fields(line)[0]

		if bikeName == "road" {
			d.direct(&roadBikeBuilder{})
		} else {
			d.direct(&moutainBikeBuilder{})
		}

	}

}
