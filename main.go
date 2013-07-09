package main

import (
	"fmt"
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	"github.com/jonomacd/playjunk/panel"
	um "github.com/jonomacd/playjunk/usermanagement"
)

var _ object.Object
var _ = image.Image{}
var _ = um.User{}
var _ = panel.Panel{}

func main() {
	fmt.Println("Hello World!")
}
