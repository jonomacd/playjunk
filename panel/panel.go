package panel

import (
	"github.com/skelterjohn/geom"
)

type Panel struct {
	Position geom.Coord
	Z        int
}

func NewPanel() *Panel {
	p := &Panel{
		Position: geom.Coord{X: 0, Y: 0},
		Z:        0,
	}
	return p
}
