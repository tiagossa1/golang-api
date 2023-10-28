package main

import (
	"log"
	"net/http"
	"os"
	"peopleapi/app"
	"peopleapi/app/database"
)

func main() {
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	checkError(err)

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("API is running on port 9000.")

	err = http.ListenAndServe(":9000", nil)
	checkError(err)
}

func checkError(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
