package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/datumforge/datum/pkg/utils/totp/testing/views"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":3321"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	Handlers()
	http.ListenAndServe(Port, nil)
}

func Handlers() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", views.GenerateTOTP)
}
