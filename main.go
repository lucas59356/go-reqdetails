package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lucas59356/go-reqdetails/v1"
)

const bind = "0.0.0.0:80"

func main() {
	println("Iniciando...")
	h := &v1.Handler{}
	r := mux.NewRouter().StrictSlash(true).SkipClean(true)
	r.HandleFunc("/{something}", h.Handle).MatcherFunc(
		func(r *http.Request, rm *mux.RouteMatch) bool {
			return true
		})
	println("Escutando em " + bind)
	err := http.ListenAndServe(bind, r)
	panic(err)
}
