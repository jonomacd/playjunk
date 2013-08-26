package clientconnection

import (
	pq "github.com/jonomacd/playjunk/priorityqueue"
	um "github.com/jonomacd/playjunk/usermanagement"
	//"log"
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
	if um.Usermap[r.FormValue("id")].Dataqueue == nil {
		um.Usermap[r.FormValue("id")].Dataqueue = pq.NewPriorityQueue()
	}
	um.Usermap[r.FormValue("id")].Dataqueue.Push(
		&pq.Item{
			Value:    r.FormValue("body"),
			Priority: defaultpriority,
		})
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	um.DeleteUser(r.FormValue("id"))
}

func newUser(w http.ResponseWriter, r *http.Request) {
	um.InsertUser(r.FormValue("id"), "")
}
