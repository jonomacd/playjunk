package usermanagement

import (
	"fmt"
)

var Usermap map[string]*User

type User struct {
	Id   string
	Name string
}

func (self User) Read(p []byte) (n int, err error) {
	return
}

func (self User) Write(p []byte) (n int, err error) {
	_, err = http.PostForm("http://localhost:8080/forward",
		url.Values{"id": {self.Id}, "data": {p}})
	if err != nil {
		log.Println("error forwarding ::", err)
		return 0, err
	}

	return len(p), nil
}

func InsertUser(id string, name string) error {
	u := &User{Id: id, Name: name}
	if _, ok := Usermap[id]; !ok {
		Usermap[id] = u
		return nil
	} else {
		return fmt.Errorf("User Already Exists: %s: %s", id, name)
	}
}

func DeleteUser(id string) error {
	delete(Usermap, id)
	return nil
}
