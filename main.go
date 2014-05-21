package main

import (
	//"fmt"
	"net/http"
	"github.com/blooberr/netrunner-draft/draft"
)

func main() {
	server := draft.NewServer()
	go server.Launch()

	// can split this off into a separate service (rails, node.js, whatever.)
	http.Handle("/", http.FileServer(http.Dir("public/")))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic(err)
	}
}

