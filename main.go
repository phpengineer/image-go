package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func init() {
	initConfig()
}


func main() {
	server := conf.GetString("listen.server")
	router := httprouter.New()
	handler := NewHandler()
	router.GET("/", handler.Index)
	router.POST("image/upload", handler.Upload)
	router.GET("/image/:name",handler.getImage)

	log.Println("start server: ", server)
	err := http.ListenAndServe(server, router)
	if err != nil {
		log.Println(err.Error())
	}
}



