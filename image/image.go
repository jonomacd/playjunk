package image

import (
	"fmt"
	"github.com/skelterjohn/geom"
	"image"
	"os"
)

type Image struct {
	Path            string
	Size            geom.Rect
	CellSize        geom.Rect
	CellNumber      int
	AnimationGroups map[string]AnimationGroup
}

type AnimationGroup struct {
	Start int
	End   int
	Speed int
}

func (self *Image) ConstructCells(number int, max geom.Coord) (err error) {

	size := geom.Rect{}
	size.Min = geom.Coord{X: 0, Y: 0}
	size.Max = max
	self.CellNumber = number
	self.CellSize = size
	return
}

func (self *Image) AddAnimationGroup(name string, start int, end int, speed int) (err error) {
	ag := AnimationGroup{Start: start, End: end, Speed: speed}
	if end >= self.CellNumber {
		return fmt.Errorf("Range greater than Cell number. %v - %v", start, end)
	}
	self.AnimationGroups[name] = ag
	return
}

func NewImage(path string) (*Image, error) {
	f, ferr := os.Open(path)
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", ferr)
		return nil, fmt.Errorf("Could not open image", ferr)
	}
	imConf, _, err := image.DecodeConfig(f)
	if err != nil {
		return nil, fmt.Errorf("Could not decode image", err)
	}

	im := Image{}
	im.Path = path
	size := geom.Rect{}
	size.Min = geom.Coord{X: 0, Y: 0}
	size.Max = geom.Coord{X: float64(imConf.Width), Y: float64(imConf.Height)}
	im.Size = size
	f.Close()
	return &im, nil
}
