package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/mojibake", Mojibake),
		)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Query struct {
	ReqBody string
}

func Mojibake(w rest.ResponseWriter, r *rest.Request){

}