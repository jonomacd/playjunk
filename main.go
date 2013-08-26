package main

import (
	"fmt"
	cc "github.com/jonomacd/playjunk/clientconnection"
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	um "github.com/jonomacd/playjunk/usermanagement"
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
	conErr := make(chan error)
	go cc.Connect(conErr)
	fmt.Println(<-conErr)
}
