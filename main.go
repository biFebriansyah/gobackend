package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/biFebriansyah/gobackend/src/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("service run on port 8080")
	http.ListenAndServe(":8080", mainRoute)

}
