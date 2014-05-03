package draw

import (
	"encoding/json"
	"fmt"
	"github.com/jonomacd/playjunk/object"
	"github.com/skelterjohn/geom"
	"sort"
)

type By func(o1, o2 object.Object) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(o []object.Object) {
	ps := &objectSorter{
		objects: o,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type objectSorter struct {
	objects []object.Object
	by      func(o1, o2 object.Object) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *objectSorter) Len() int {
	return len(s.objects)
}

// Swap is part of sort.Interface.
func (s *objectSorter) Swap(i, j int) {
	s.objects[i], s.objects[j] = s.objects[j], s.objects[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *objectSorter) Less(i, j int) bool {
	return s.by(s.objects[i], s.objects[j])
}

func SortObjects(o []object.Object) {

	z := func(o1, o2 object.Object) bool {
		return o1.Z() < o2.Z()
	}

	By(z).Sort(o)
}

func TranslateCoords(o object.Object) (*geom.Coord, error) {
	pan := o.Panel()
	coord := *o.Coord()

	var err error
	fmt.Printf("%+v \n", pan)
	tries := 0
	for pan != nil {
		fmt.Printf("%+v\n", pan)
		coord = pan.Coord().Plus(coord)
		pan = pan.Panel()
		tries++
		if tries >= object.MaxPanelDepth {
			err = fmt.Errorf("Exceeded Max Panel Depth (possible circular panel path)")
			break
		}
	}
	return &coord, err
}

func FlattenObjects(os []object.Object) ([]object.Object, error) {
	for _, o := range os {
		coord, err := TranslateCoords(o)
		if err != nil {
			return nil, err
		}

		o.SetCoord(coord)
	}
	return os, nil

}

func Intersect(o1 object.Object, o2 object.Object) bool {
	return true
}

func MarshalToWire(os []object.Object) []byte {
	wireArr := make([]map[string]interface{}, len(os))
	for ii, o := range os {
		m := make(map[string]interface{})
		m["Image"] = o.Image().Url
		fmt.Println(m)
		m["Id"] = o.Image().Url
		m["SX"] = 0
		m["SY"] = 0
		m["SW"] = o.Image().Size.Max.X
		m["SH"] = o.Image().Size.Max.Y
		m["DX"] = o.Coord().X
		m["DY"] = o.Coord().Y
		m["DW"] = o.Image().Size.Max.X
		m["DH"] = o.Image().Size.Max.Y

		wireArr[ii] = m

	}
	b, _ := json.Marshal(wireArr)
	return b
}
