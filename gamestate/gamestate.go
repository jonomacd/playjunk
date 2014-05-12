package gamestate

import (
	//"fmt"
	"github.com/jonomacd/playjunk/character"
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	"github.com/skelterjohn/geom"
	"strconv"
)

const (
	Width      = 6
	Height     = 12
	SquareSize = 56
	FillHeight = 4
)

type Board [][]*Square

type Square struct {
	Contains []object.Object
}

func NewBoard() Board {
	var squares Board
	squares = Board(make([][]*Square, Width))

	for ii, _ := range squares {
		squares[ii] = make([]*Square, Height)
	}
	return squares
}

func (b Board) InitialFill() {
	for ii, _ := range b {
		for jj := 0; jj < FillHeight; jj++ {
			b[ii][jj] = &Square{
				Contains: make([]object.Object, 1),
			}
			mc := &character.MainCharacter{}
			mc.IdMC = "monster" + strconv.Itoa(ii) + strconv.Itoa(jj)
			mc.CoordMC = &geom.Coord{X: float64(ii * SquareSize), Y: float64(jj * SquareSize)}
			mc.ImageMC = image.Images["resources/fakewario.gif"]
			mc.SizeMC = &mc.ImageMC.Size
			mc.ZMC = 1
			mc.PreviousLoc = mc.SizeMC
			mc.DirtyMC = true
			b[ii][jj].Contains[0] = mc
		}
	}
}

func (b Board) GetObjects() []object.Object {

	objects := make([]object.Object, 0)

	for ii, _ := range b {
		for jj, _ := range b[ii] {
			if b[ii][jj] != nil {
				objects = append(objects, b[ii][jj].Contains...)
			}
		}
	}

	return objects
}
