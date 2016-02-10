package controller

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

var session *r.Session

// SetSession sets the session to be used in calls to rethinkDB
func SetSession() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "teacherMe",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
