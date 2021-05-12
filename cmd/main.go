package main

import (
	"config-server/pkg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)
func getproperties(w http.ResponseWriter, r *http.Request) {
	name:= viper.Get("myname");
	log.Println(name)
}
func main() {
	pkg.NewConfig()

	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/props", getproperties)
    log.Fatal(http.ListenAndServe(":8081", muxRouter))
}