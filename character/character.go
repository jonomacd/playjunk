package character

import(
	"github.com/jonomacd/playjunk/object"
	"github.com/jonomacd/playjunk/image"
	"github.com/skelterjohn/geom"
)

type MainCharacter struct {
	CoordMC *geom.Coord
	SizeMC *geom.Rect
	AnimateMC bool
	PanelMC *object.Panel
	ZMC int
	ImageMC *image.Image
	AlphaMC int
	VisibleMC bool
}

func (self *MainCharacter) Coord() *geom.Coord {
	return self.CoordMC
}

func (self *MainCharacter) SetCoord(coord *geom.Coord)  {
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
	return false
}

func (self *MainCharacter) Equals(o object.Object) bool {
	return o.Coord().Equals(self.Coord()) &&
		o.Size().Equals(self.Size()) &&
		o.Panel().Equals(o.Panel())
}

func (self *MainCharacter) Visible() bool {
	return self.VisibleMC
}
