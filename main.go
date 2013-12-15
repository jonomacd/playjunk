package main

import (
	"fmt"
	cc "github.com/jonomacd/playjunk/clientconnection"
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	um "github.com/jonomacd/playjunk/usermanagement"
	"github.com/skelterjohn/geom"
	"github.com/jonomacd/playjunk/draw"
)

//TEMP FOR COMPILATION
var _ object.Object
var _ = image.Image{}
var _ = um.User{}
var _ = object.Panel{}
var thing object.Object = &object.Panel{Alph: 6}

//END TEMP

func main() {
	fmt.Println("Hello World!", thing.Alpha())

	p := &object.Panel{}
	p.Position = geom.Coord{}
	p.Extent = geom.Rect{}
	p.Extent.Min = p.Position
	p.Extent.Max = geom.Coord{X:20.0, Y:20.0}
	p.Depth = 0
	fmt.Println(p)

	poo := &object.Panel{}
	poo.Position = geom.Coord{X:2.0, Y:2.0}
	poo.Extent = geom.Rect{}
	poo.Extent.Min = p.Position
	poo.Extent.Max = geom.Coord{X:20.0, Y:20.0}
	poo.Depth = 0
	poo.Pan = p
	fmt.Println(poo)

	po := &object.Panel{}
	po.Position = geom.Coord{X:10.0, Y:10.0}
	po.Extent = geom.Rect{}
	po.Extent.Min = p.Position
	po.Extent.Max = geom.Coord{X:20.0, Y:20.0}
	po.Depth = 0
	po.Pan = poo
	fmt.Println(po)

	bo:=&object.BlankObject{}
	bo.BlankCoord = &geom.Coord{X:5,Y:5}
	bo.BlankPanel = po
	fmt.Println(bo)

	boo:=&object.BlankObject{}
	boo.BlankCoord = &geom.Coord{X:1,Y:7}
	boo.BlankPanel = poo
	fmt.Println(boo)

	os := make([]object.Object, 2)
	os[0] = bo
	os[1] = boo

	draw.FlattenObjects(os)
	fmt.Printf("%+v, \n",os[0].Coord())
	fmt.Printf("%+v, \n",os[1].Coord())

	image.AddImage("/home/jono/code/go/home/src/github.com/jonomacd/playjunk/resources/ness.jpg")

	conErr := make(chan error)
	go cc.Connect(conErr)
	fmt.Println(<-conErr)
}
