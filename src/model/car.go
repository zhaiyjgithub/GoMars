package model

import "GoMars/src/constant"

var MyCar *Car

type Position struct {
	X int
	Y int
}

type Car struct {
	Name string
	Position
	Direction constant.Direction
}
