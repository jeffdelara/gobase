package main

import (
	"fmt"
	pagescontroller "github.com/jeffdelara/gobase/controllers"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets"))

	server.Handle("/assets/", http.StripPrefix("/assets/", fs))
	server.Handle("/", pagescontroller.Home())
	server.Handle("/about", pagescontroller.About())

	fmt.Println("Listening on :3001")
	_ = http.ListenAndServe(":3001", server)
}
