package draw

import (
	"errors"
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

func TranslateCoords(p *object.Panel, o object.Object) (*geom.Coord, error) {
	if p.Panel() != nil {
		return nil, errors.New("The given panel is not a base panel")
	}
	absCoord := *o.Coord()
	pan := *o.Panel()
	for !p.Equals(pan) {
		absCoord = o.Panel().Coord().Plus(*o.Coord())
		pan = *pan.Panel()
	}

	return &absCoord, nil
}

func FilterDirty(os []object.Object) {
	for _, o := range os {
		if o.Dirty() {
			for _, do := range os {
				if !do.Dirty() {
					if Intersect(o, do) {
						//set do to dirty

					}
				}
			}
		}
	}
}

func Intersect(o1 object.Object, o2 object.Object) bool {
	return true
}
