package usermanagement

import (
	"fmt"
)

var Usermap map[string]*User

type User struct {
	Id   string
	Name string
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
