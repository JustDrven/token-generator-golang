package models

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Addr string
}

func (srv Server) Listen() {
	println(fmt.Sprintf("Server is running at localhost%v", srv.Addr))

	err := http.ListenAndServe(srv.Addr, nil)

	if err != nil {
		log.Fatal(err)
		return
	}



}