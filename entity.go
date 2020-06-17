/* */

package main

type Entity struct {
	X     int
	Y     int
	Char  string
	Color string
}

func (e *Entity) move(dx, dy int) {
	e.X += dx
	e.Y += dy
}
