package usermanagement

import (
	"fmt"
	"github.com/jonomacd/playjunk/character"
	"github.com/jonomacd/playjunk/draw"
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	pq "github.com/jonomacd/playjunk/priorityqueue"
	"github.com/skelterjohn/geom"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var Usermap map[string]*User = make(map[string]*User)

type User struct {
	Id          string
	Name        string
	UserState   *State
	UserContext *Context
}

type State struct {
	DrawState *draw.DrawState
	GameState string // TODO build the game state
}

type Context struct {
	Dataqueue *pq.PriorityQueue
}

func (self *User) Read(p []byte) (n int, err error) {
	if self == nil {
		return 0, fmt.Errorf("User does not exist")
	}
	if self.UserContext.Dataqueue != nil {
		n = copy(p, []byte(self.UserContext.Dataqueue.Pop().(*pq.Item).Value))
		p = p[:n]
		if n == 0 {
			return n, nil
		}
	} else {
		return 0, fmt.Errorf("No Queue")
	}
	return
}

func (self *User) Write(p []byte) (n int, err error) {
	_, err = http.PostForm("http://localhost:8080/forward",
		url.Values{"id": {self.Id}, "body": {string(p)}})
	if err != nil {
		log.Println("error forwarding ::", err)
		return 0, err
	}

	return len(p), nil
}

func (self *User) Draw() {
	var something object.Object
	something = object.NewPanel()
	somethingarr := make([]object.Object, 1)
	somethingarr[1] = something
	draw.SortObjects(somethingarr)
}

func InsertUser(id string, name string) error {

	if _, ok := Usermap[id]; !ok {

		u := &User{Id: id, Name: name}

		u.UserContext = &Context{}
		// Set up the data queue
		u.UserContext.Dataqueue = pq.NewPriorityQueue()

		u.UserState = &State{}
		// Set up the draw state
		u.UserState.DrawState = draw.NewDrawState()

		bg := &character.MainCharacter{}
		bg.IdMC = "background"
		bg.CoordMC = &geom.Coord{X: 0, Y: 0}
		bg.ImageMC = image.Images["resources/Nice-blue-background-Desktop-Wallpaper.jpg"]
		bg.SizeMC = &bg.ImageMC.Size
		bg.ZMC = 0
		bg.PreviousLoc = bg.SizeMC
		bg.DirtyMC = true
		fmt.Printf("%+v\n", bg.Size())
		u.UserState.DrawState.Add(bg)

		mc := &character.MainCharacter{}
		mc.IdMC = "Kirby"
		mc.CoordMC = &geom.Coord{X: 20, Y: 20}
		mc.ImageMC = image.Images["resources/iceKing.png"]
		mc.SizeMC = &mc.ImageMC.Size
		mc.ZMC = 1
		mc.PreviousLoc = mc.SizeMC
		mc.DirtyMC = false
		u.UserState.DrawState.Add(mc)

		mc2 := &character.MainCharacter{}
		mc2.IdMC = "Paula"
		mc2.CoordMC = &geom.Coord{}
		mc2.ImageMC = image.Images["resources/paula.jpg"]
		mc2.SizeMC = &mc2.ImageMC.Size
		mc2.ZMC = 5
		mc2.PreviousLoc = mc2.SizeMC
		mc2.DirtyMC = true
		u.UserState.DrawState.Add(mc2)

		// Initialize the game state (TODO)
		u.UserState.GameState = "Not Implemented"

		Usermap[id] = u
		log.Println("User Added!", id)
		time.Sleep(1 * time.Second)
		go CheckForData(id)
		return nil
	} else {
		return fmt.Errorf("User Already Exists: %s: %s", id, name)
	}
	return nil
}

func DeleteUser(id string) error {

	// Todo Probaly need clean up a lot of junk when users leave
	delete(Usermap, id)
	log.Println("User Removed:", id)
	return nil
}

func CheckForData(id string) {
	var p = make([]byte, 40)
	u := Usermap[id]
	Usermap[id].Write(u.UserState.DrawState.MarshalToWire())
	for {
		time.Sleep(10 * time.Millisecond)
		n, err := Usermap[id].Read(p)
		if err != nil {
			if Usermap[id] == nil {
				return
			}
			fmt.Println("Not sure what to do with this error so we ignore it for now", err)
		}
		data := string(p[:n])
		dataArr := strings.Split(data, ":")

		if len(dataArr) != 0 {
			if dataArr[0] == "i" {
				screensizeArr := strings.Split(dataArr[1], ",")

				x, _ := strconv.Atoi(screensizeArr[0])
				y, _ := strconv.Atoi(screensizeArr[1])
				Usermap[id].UserState.DrawState.BasePanel = object.NewPanel()
				Usermap[id].UserState.DrawState.BasePanel.Extent = geom.Rect{Min: geom.Coord{X: 0, Y: 0}, Max: geom.Coord{X: float64(x), Y: float64(y)}}

				fmt.Printf("Updated Screen Size %+v\n", Usermap[id].UserState.DrawState.BasePanel)
			}

			if dataArr[0] == "c" {
				fmt.Println("Move ", dataArr[1])
				ob := u.UserState.DrawState.Objects["Paula"]

				if dataArr[1] == "down" {
					crd := ob.Coord().Y
					crd += 3
					ob.SetCoord(&geom.Coord{X: ob.Coord().X, Y: crd})
					u.UserState.DrawState.Objects["Paula"].(*character.MainCharacter).ZMC = 1
					ob.(*character.MainCharacter).ZMC = 5
				}
				if dataArr[1] == "right" {
					crd := ob.Coord().X
					crd += 3
					ob.SetCoord(&geom.Coord{X: crd, Y: ob.Coord().Y})
					u.UserState.DrawState.Objects["Paula"].(*character.MainCharacter).ZMC = 5
					ob.(*character.MainCharacter).ZMC = 1
				}
				if dataArr[1] == "left" {
					crd := ob.Coord().X
					crd += -3
					ob.SetCoord(&geom.Coord{X: crd, Y: ob.Coord().Y})
				}
				if dataArr[1] == "up" {
					crd := ob.Coord().Y
					crd += -3
					ob.SetCoord(&geom.Coord{X: ob.Coord().X, Y: crd})
				}

				Usermap[id].Write(u.UserState.DrawState.MarshalToWire())

			}
			if dataArr[0] == "im" {
				Usermap[id].Write([]byte(`[{"Image":"` + u.UserState.DrawState.Objects["Ness"].Image().Url + `", "Id":"M"}]`))
			}
		}
	}
}
