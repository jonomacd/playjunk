package usermanagement

import (
	"fmt"
	"github.com/jonomacd/playjunk/draw"
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
	Id        string
	Name      string
	Dataqueue *pq.PriorityQueue
	BasePanel *object.Panel
	ds        []*object.Object
}

func (self *User) Read(p []byte) (n int, err error) {
	if self.Dataqueue != nil {
		n = copy(p, []byte(self.Dataqueue.Pop().(*pq.Item).Value))
		p = p[:n]
		if n == 0 {
			return n, fmt.Errorf("No Data")
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
		Usermap[id] = u
		log.Println("User Added!", id)
		u.Write([]byte("Hello There!"))
		go CheckForData(id)
		return nil
	} else {
		return fmt.Errorf("User Already Exists: %s: %s", id, name)
	}
}

func DeleteUser(id string) error {
	delete(Usermap, id)
	log.Println("User Removed:", id)
	return nil
}

func CheckForData(id string) {
	var p = make([]byte, 40)
	for {
		time.Sleep(100 * time.Millisecond)
		n, err := Usermap[id].Read(p)
		if err != nil {
			//fmt.Println("error")
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
		}
	}
}
