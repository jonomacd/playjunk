package clientconnection

import (
	"fmt"
	pq "github.com/jonomacd/playjunk/priorityqueue"
	um "github.com/jonomacd/playjunk/usermanagement"
	"net/http"
)

var defaultpriority int = 1

func Connect(errChan chan error) {
	http.HandleFunc("/data", data)
	http.HandleFunc("/delete", removeUser)
	http.HandleFunc("/new", newUser)
	if err := http.ListenAndServe(":8099", nil); err != nil {
		errChan <- err
	}
}

func data(w http.ResponseWriter, r *http.Request) {
	if um.Usermap[r.FormValue("id")] != nil {
		if um.Usermap[r.FormValue("id")].UserContext != nil {
			if um.Usermap[r.FormValue("id")].UserContext.Dataqueue == nil {
				um.Usermap[r.FormValue("id")].UserContext.Dataqueue = pq.NewPriorityQueue()
			}
		} else {
			return
		}
	} else {
		return
	}
	body := r.FormValue("body")
	fmt.Println("got message: ", body)
	um.Usermap[r.FormValue("id")].UserContext.Dataqueue.Push(
		&pq.Item{
			Value:    body,
			Priority: defaultpriority,
		})

	fmt.Println("handler returning")
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	um.DeleteUser(r.FormValue("id"))
}

func newUser(w http.ResponseWriter, r *http.Request) {

	um.InsertUser(r.FormValue("id"), "")
}
