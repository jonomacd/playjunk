package rectanglestore

import (
	"fmt"
	"github.com/skelterjohn/geom"
	"math"
)

const (
	InitialSize = 5
)

type RectangleStore struct {
	Values [][][]interface{}
	Width  int
	Height int
}

func NewRectangleStore(box *geom.Rect, innerBox *geom.Rect) *RectangleStore {

	outerWidth := box.Width()
	outerHeight := box.Height()

	innerWidth := innerBox.Width()
	innerHeight := innerBox.Height()

	x := math.Floor(outerWidth / innerWidth)
	y := math.Floor(outerHeight / innerHeight)

	rs := &RectangleStore{
		Values: make([][][]interface{}, int(x)),
		Width:  int(innerWidth),
		Height: int(innerHeight),
	}

	for ii, _ := range rs.Values {
		rs.Values[ii] = make([][]interface{}, int(y))
		for jj, _ := range rs.Values[ii] {
			rs.Values[ii][jj] = make([]interface{}, InitialSize)
		}
	}

	return rs
}

func (rs *RectangleStore) Add(box *geom.Rect, coord *geom.Coord, i interface{}) error {

	x := coord.X
	y := coord.Y
	xPrime := x + box.Width()
	yPrime := y + box.Height()

	fmt.Println(x, y)
	fmt.Println(xPrime, yPrime)

	startX := int(x) / rs.Width
	endX := int(xPrime) / rs.Width

	startY := int(y) / rs.Height
	endY := int(yPrime) / rs.Height

	for ii := startX; ii <= endX; ii++ {
		for jj := startY; jj <= endY; jj++ {
			fmt.Printf("%v, %v :: %v, %v\n", ii, jj, len(rs.Values), len(rs.Values[ii]))
			rs.Values[ii][jj] = append(rs.Values[ii][jj], i)
		}
	}

	return nil
}

func (rs *RectangleStore) Remove(box *geom.Rect, coord *geom.Coord, i interface{}) error {

	x := coord.X
	y := coord.Y
	xPrime := x + box.Width()
	yPrime := y + box.Height()

	startX := int(x) / rs.Width
	endX := int(xPrime) / rs.Width

	startY := int(y) / rs.Height
	endY := int(yPrime) / rs.Height

	for ii := startX; ii < endX; ii++ {
		for jj := startY; jj < endY; jj++ {
			for kk, toDel := range rs.Values[ii][jj] {
				if toDel == i {
					fmt.Println("removed")
					copy(rs.Values[ii][jj][kk:], rs.Values[ii][jj][kk+1:])
					rs.Values[ii][jj] = rs.Values[ii][jj][:len(rs.Values[ii][jj])-1]
				}
			}
		}
	}

	return nil
}

func (rs *RectangleStore) Inside(box *geom.Rect, coord *geom.Coord) []interface{} {
	var i []interface{}

	x := coord.X
	y := coord.Y
	xPrime := x + box.Width()
	yPrime := y + box.Height()

	startX := int(x) / rs.Width
	endX := int(xPrime) / rs.Width

	startY := int(y) / rs.Height
	endY := int(yPrime) / rs.Height
	dupMap := make(map[interface{}]bool)
	for ii := startX; ii < endX; ii++ {
		for jj := startY; jj < endY; jj++ {

			for _, val := range rs.Values[ii][jj] {
				if val != nil && !dupMap[val] {
					dupMap[val] = true
					i = append(i, val)
				}
			}
		}
	}

	return i
}
