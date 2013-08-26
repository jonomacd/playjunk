package panel

/*import (
	"github.com/jonomacd/playjunk/object"
	"github.com/skelterjohn/geom"
	"image"
)

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

func (this *Panel) Coord() *geom.Coord {
	return &this.Position
}

func (this *Panel) Size() *geom.Rect {
	return &this.Extent
}

func (this *Panel) Animate() bool {
	return false
}
func (this *Panel) Panel() *Panel {
	return this.Pan
}

func (this *Panel) Z() int {
	return this.Depth
}

func (this *Panel) Image() *image.Image {
	return nil
}

func (this *Panel) Alpha() int {
	return this.Alph
}

func (this *Panel) Equals(that Panel) bool {
	return this.Position.Equals(that.Coord()) &&
		this.Extent.Equals(that.Size()) &&
		this.Depth == that.Z() &&
		this.Alph == that.Alpha()

}
*/
