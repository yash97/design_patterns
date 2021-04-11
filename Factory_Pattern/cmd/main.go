package main

import (
	"os"

	"github.com/yash97/design_patterns/Factory_Pattern/svg"
)

func main() {
	doc := &svg.Document{
		ShapeFactories: []svg.ShapeFactory{
			&svg.CircleFactory{},
			&svg.RectangleFactory{},
		},
	}
	doc.Draw(os.Stdout)
}
