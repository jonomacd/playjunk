package object

import (
	"github.com/skelterjohn/geom"
	"image"
)

type Object interface {
	Coord() *geom.Coord
	Size() *geom.Rect
	Animate() bool
	Panel() *Panel
	Z() int
	Image() *image.Image
	Alpha() int
	Equals(o Object) bool
	Dirty() bool
}

type BlankObject struct {
	BlankCoord *geom.Coord
	BlankSize  *geom.Rect
	BlankPanel *Panel
}

func (self BlankObject) Coord() *geom.Coord {
	return self.BlankCoord
}

func (self BlankObject) Size() *geom.Rect {
	return self.BlankSize
}

func (self BlankObject) Panel() *Panel {
	return self.BlankPanel
}

func (self BlankObject) Animate() bool {
	return false
}

func (self BlankObject) Z() int {
	return 0
}

func (self BlankObject) Image() *image.Image {
	return nil
}

func (self BlankObject) Alpha() int {
	return 0
}

func (self BlankObject) Dirty() bool {
	return false
}

func (self BlankObject) Equals(o Object) bool {
	return o.Coord().Equals(self.Coord()) &&
		o.Size().Equals(self.Size()) &&
		o.Panel().Equals(o.Panel())
}

type Panel struct {
	Position geom.Coord
	Extent   geom.Rect
	Depth    int
	Alph     int
	Pan      *Panel
}

func NewPanel() *Panel {
	p := &Panel{
		Position: geom.Coord{X: 0, Y: 0},
		Depth:    0,
	}
	return p
}

func (this Panel) Coord() *geom.Coord {
	return &this.Position
}

func (this Panel) Size() *geom.Rect {
	return &this.Extent
}

func (this Panel) Animate() bool {
	return false
}
func (this Panel) Panel() *Panel {
	return this.Pan
}

func (this Panel) Z() int {
	return this.Depth
}

func (this Panel) Image() *image.Image {
	return nil
}

func (this Panel) Alpha() int {
	return this.Alph
}

func (this Panel) Dirty() bool {
	return false
}

func (this Panel) Equals(that Object) bool {
	return this.Position.Equals(that.Coord()) &&
		this.Extent.Equals(that.Size()) &&
		this.Depth == that.Z() &&
		this.Alph == that.Alpha()

}
