package main

import (
	"fmt"
	cc "github.com/jonomacd/playjunk/clientconnection"
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	"github.com/jonomacd/playjunk/panel"
	um "github.com/jonomacd/playjunk/usermanagement"
)

//TEMP FOR COMPILATION
var _ object.Object
var _ = image.Image{}
var _ = um.User{}
var _ = panel.Panel{}
var thing object.Object = &panel.Panel{Alph: 6}

//END TEMP

func main() {
	fmt.Println("Hello World!", thing.Alpha())
	conErr := make(chan error)
	cc.Connect(conErr)
}
