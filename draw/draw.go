package draw

import (
	"encoding/json"
	"fmt"
	"github.com/jonomacd/playjunk/object"
	rs "github.com/jonomacd/playjunk/rectanglestore"
	"github.com/skelterjohn/geom"
	"sort"
)

/*const (
	InitalObjectSize = 20
)*/

var DirtyResolution *geom.Rect = &geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 10, Y: 10}}

type DrawState struct {
	objectArr []object.Object
	Dirty     *rs.RectangleStore
	BasePanel *object.Panel
	Objects   map[string]object.Object
}

func NewDrawState() *DrawState {
	ds := &DrawState{
		objectArr: make([]object.Object, 0),
		Dirty:     rs.NewRectangleStore(&geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 2000.0, Y: 3000.0}}, &geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 10, Y: 10}}),
		BasePanel: object.NewPanel(),
		Objects:   make(map[string]object.Object),
	}
	return ds
}

func (ds *DrawState) Add(o object.Object) error {

	// Must have an Id to add an object
	if len(o.Id()) == 0 {
		return fmt.Errorf("Cannot add object: ID not set")
	}

	// Add to users map store
	ds.Objects[o.Id()] = o

	// Add to users array object store
	ds.objectArr = append(ds.objectArr, o)

	// Add to objects Dirty array
	return ds.Dirty.Add(o.Size(), o.Coord(), o)
}

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

type DrawObject struct {
	object.Object
	sx   int
	sy   int
	size *geom.Rect
}

func (ds *DrawState) GetDrawSet() []*DrawObject {
	objects := make([]*DrawObject, 0)
	dupmap := make(map[string]bool)
	for _, obj := range ds.objectArr {
		if obj.Dirty() {

			do := &DrawObject{}
			do.Object = obj
			do.size = obj.Size()
			objects = append(objects, do)
			dupmap[obj.Id()] = true
			todraw := ds.Dirty.Inside(obj.Size(), obj.Coord())
			for _, interObj := range todraw {
				objCast := interObj.(object.Object)
				if !dupmap[objCast.Id()] && !objCast.Dirty() {

					x1 := obj.Previous().Min.X
					y1 := obj.Previous().Min.Y

					x2 := objCast.Coord().X
					y2 := objCast.Coord().Y

					x := int(x1 - x2)
					y := int(y1 - y2)
					do := &DrawObject{}
					do.Object = objCast
					do.sx = x
					do.sy = y
					do.size = obj.Previous()

					objects = append(objects, do)
					dupmap[objCast.Id()] = true
				}
			}
			obj.ClearDirty()
		}
	}
	return objects
}

//func

func (ds *DrawState) MarshalToWire() []byte {

	os := ds.GetDrawSet()
	wireArr := make([]map[string]interface{}, len(os))
	for ii, o := range os {
		fmt.Println("Drawing: ", o.Id())
		m := make(map[string]interface{})
		m["Image"] = o.Image().Url
		fmt.Println(m)
		m["Id"] = o.Image().Url
		m["SX"] = o.sx
		m["SY"] = o.sy
		m["SW"] = int(o.size.Width())
		m["SH"] = int(o.size.Height())
		m["DX"] = int(o.Coord().X) + o.sx
		m["DY"] = int(o.Coord().Y) + o.sy
		m["DW"] = int(o.size.Width())
		m["DH"] = int(o.size.Height())
		m["DO"] = o.Z()

		wireArr[ii] = m

	}
	b, _ := json.Marshal(wireArr)
	return b
}
