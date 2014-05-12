package character

import (
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	"github.com/skelterjohn/geom"

	"fmt"
)

type MainCharacter struct {
	CoordMC     *geom.Coord
	SizeMC      *geom.Rect
	AnimateMC   bool
	PanelMC     *object.Panel
	ZMC         int
	ImageMC     *image.Image
	AlphaMC     int
	DirtyMC     bool
	VisibleMC   bool
	IdMC        string
	PreviousLoc *geom.Rect
}

func (self *MainCharacter) Id() string {
	return self.IdMC
}

func (self *MainCharacter) Coord() *geom.Coord {
	return self.CoordMC
}

func (self *MainCharacter) SetCoord(coord *geom.Coord) {
	fmt.Printf("before: %+v\n", self.CoordMC)
	self.DirtyMC = true
	self.PreviousLoc = &geom.Rect{}
	self.PreviousLoc.Min = geom.Coord{X: self.CoordMC.X, Y: self.CoordMC.Y}
	self.PreviousLoc.Max = geom.Coord{X: self.CoordMC.X + self.Size().Width(), Y: self.CoordMC.Y + self.Size().Height()}

	fmt.Printf("%+v\n %+v\n", self.PreviousLoc, coord)
	self.CoordMC = coord
}

func (self *MainCharacter) Size() *geom.Rect {
	return self.SizeMC
}

func (self *MainCharacter) Panel() *object.Panel {
	return self.PanelMC
}

func (self *MainCharacter) Animate() bool {
	return self.AnimateMC
}

func (self *MainCharacter) Z() int {
	return self.ZMC
}

func (self *MainCharacter) Image() *image.Image {
	return self.ImageMC
}

func (self *MainCharacter) Alpha() int {
	return self.AlphaMC
}

func (self *MainCharacter) Dirty() bool {

	return self.DirtyMC
}

func (self *MainCharacter) Equals(o object.Object) bool {
	return o.Coord().Equals(self.Coord()) &&
		o.Size().Equals(self.Size()) &&
		o.Panel().Equals(o.Panel())
}

func (self *MainCharacter) Visible() bool {
	return self.VisibleMC
}

func (self *MainCharacter) Previous() *geom.Rect {
	return self.PreviousLoc
}

func (self *MainCharacter) ClearDirty() {
	self.DirtyMC = false
}

func (self *MainCharacter) AddToPanel(p *object.Panel) {
	self.PanelMC = p
}
