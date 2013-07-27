package clientconnection

import (
	um "github.com/jonomacd/playjunk/usermanagement"
	"net/http"
)

func Connect(errChan chan error) {
	http.HandleFunc("/data", data)
	http.HandleFunc("/delete", removeUser)
	http.HandleFunc("/new", newUser)
	if err := http.ListenAndServe("8099", nil); err != nil {
		errChan <- err
	}
}

func data(w http.ResponseWriter, r *http.Request) {

}

func removeUser(w http.ResponseWriter, r *http.Request) {
	um.DeleteUser(r.FormValue("id"))
}

func newUser(w http.ResponseWriter, r *http.Request) {
	um.InsertUser(r.FormValue("id"), "")
}
