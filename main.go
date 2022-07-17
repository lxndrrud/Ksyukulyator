package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/webviewKsyukulyator/storage"
	window_controllers "github.com/lxndrrud/webviewKsyukulyator/windowControllers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	go func() {
		http.Handle("/", http.FileServer(http.Dir("./frontend")))
		http.ListenAndServe(":8083", nil)
	}()
	db, err := sqlx.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	err = storage.InitStorage(db)
	if err == nil {
		window_controllers.NewWindowController(db).SetupWindow()
	} else {
		fmt.Println(err)
	}
}
