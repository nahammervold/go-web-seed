package main

import (
	"log"
	"net/http"

	"github.com/nahammervold/go-web-seed/routes"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", routes.GetRouter()))
}
