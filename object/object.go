package object

import (
	"github.com/jonomacd/playjunk/panel"
	"github.com/skelterjohn/geom"
	"image"
)

type Object interface {
	Coord() *geom.Coord
	Size() *geom.Rect
	Animate() bool
	Panel() *panel.Panel
	Z() int
	Image() *image.Image
	Alpha() int
}
