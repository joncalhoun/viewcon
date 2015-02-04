package main

import (
	"github.com/joncalhoun/viewcon/examples/controllers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	catsC := controllers.Cats
	router := httprouter.New()
	router.GET("/cats", catsC.Perform(catsC.Index))
	router.GET("/cats/feed", catsC.ReqKey(catsC.Feed))

	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
