package usermanagement

import (
	"fmt"
	"github.com/jonomacd/playjunk/character"
	"github.com/jonomacd/playjunk/draw"
	"github.com/jonomacd/playjunk/image"
	"github.com/jonomacd/playjunk/object"
	pq "github.com/jonomacd/playjunk/priorityqueue"
	rs "github.com/jonomacd/playjunk/rectanglestore"
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
	Id              string
	Name            string
	Dataqueue       *pq.PriorityQueue
	BasePanel       *object.Panel
	ds              []object.Object
	DirtyRectangles *rs.RectangleStore
}

func (self *User) Read(p []byte) (n int, err error) {
	if self == nil {
		return 0, fmt.Errorf("User does not exist")
	}
	if self.Dataqueue != nil {
		n = copy(p, []byte(self.Dataqueue.Pop().(*pq.Item).Value))
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
	u := &User{Id: id, Name: name}
	if _, ok := Usermap[id]; !ok {
		u.ds = make([]object.Object, 2)
		u.DirtyRectangles = rs.NewRectangleStore(&geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 500.0, Y: 500.0}}, &geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 10, Y: 10}})

		mc := &character.MainCharacter{}
		mc.CoordMC = &geom.Coord{X: 20, Y: 20}
		mc.ImageMC = image.Images["/home/jonomacd/go/src/github.com/jonomacd/playjunk/resources/ness.jpg"]
		mc.SizeMC = &mc.ImageMC.Size

		u.ds[1] = mc

		u.DirtyRectangles.Add(mc.Size(), mc.Coord(), mc)
		fmt.Println("added")
		mc2 := &character.MainCharacter{}
		mc2.CoordMC = &geom.Coord{}
		mc2.ImageMC = image.Images["/home/jonomacd/go/src/github.com/jonomacd/playjunk/resources/paula.jpg"]
		mc2.SizeMC = &mc2.ImageMC.Size
		mc.ZMC = 5
		u.ds[0] = mc2
		u.DirtyRectangles.Add(mc2.Size(), mc2.Coord(), mc2)
		fmt.Println("added")
		mc3 := character.MainCharacter{}
		u.DirtyRectangles.Add(&geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 100.0, Y: 50.0}}, &geom.Coord{X: 57, Y: 77}, mc3)
		u.DirtyRectangles.Remove(&geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 100.0, Y: 50.0}}, &geom.Coord{X: 57, Y: 77}, mc3)

		fmt.Printf("%+v\n", u.DirtyRectangles.Inside(&geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 50.0, Y: 50.0}}, &geom.Coord{X: 10, Y: 10}))
		u.DirtyRectangles.Remove(mc.Size(), mc.Coord(), mc)
		fmt.Printf("%+v\n", u.DirtyRectangles.Inside(&geom.Rect{Min: geom.Coord{}, Max: geom.Coord{X: 50.0, Y: 50.0}}, &geom.Coord{X: 10, Y: 10}))
		Usermap[id] = u
		log.Println("User Added!", id)

		go CheckForData(id)
		return nil
	} else {
		return fmt.Errorf("User Already Exists: %s: %s", id, name)
	}
	return nil
}

func DeleteUser(id string) error {
	delete(Usermap, id)
	log.Println("User Removed:", id)
	return nil
}

func CheckForData(id string) {
	var p = make([]byte, 40)
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
				Usermap[id].BasePanel = object.NewPanel()
				Usermap[id].BasePanel.Extent = geom.Rect{Min: geom.Coord{X: 0, Y: 0}, Max: geom.Coord{X: float64(x), Y: float64(y)}}

				fmt.Printf("Updated Screen Size %+v\n", Usermap[id].BasePanel)
			}

			u := Usermap[id]

			if dataArr[0] == "c" {
				fmt.Println("Move ", dataArr[1])

				if dataArr[1] == "down" {
					crd := u.ds[0].Coord()
					crd.Y = crd.Y + 3
					u.ds[0].SetCoord(crd)
				}
				if dataArr[1] == "right" {
					crd := u.ds[0].Coord()
					crd.X = crd.X + 3
					u.ds[0].SetCoord(crd)
				}
				if dataArr[1] == "left" {
					crd := u.ds[0].Coord()
					crd.X = crd.X - 3
					u.ds[0].SetCoord(crd)
				}
				if dataArr[1] == "up" {
					crd := u.ds[0].Coord()
					crd.Y = crd.Y - 3
					u.ds[0].SetCoord(crd)
				}

				Usermap[id].Write(draw.MarshalToWire(u.ds))
			}
			if dataArr[0] == "im" {
				Usermap[id].Write([]byte(`[{"Image":"` + u.ds[0].Image().Url + `", "Id":"M"}]`))
			}
		}
	}
}
