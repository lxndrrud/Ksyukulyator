package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/webviewKsyukulyator/storage"
	window_controllers "github.com/lxndrrud/webviewKsyukulyator/windowControllers"
	"github.com/markbates/pkger"
	_ "github.com/mattn/go-sqlite3"
)

/*
type pkgerServer struct{}

func (p *pkgerServer) Open(name string) (http.File, error) {
	return pkger.Open(name)
}
*/

func httpServer() *httptest.Server {
	return httptest.NewServer(
		http.FileServer(
			pkger.Dir("/frontend")))
}

func main() {
	exec, err := os.Executable()
	if err != nil {
		return
	}
	_ = pkger.Include(filepath.Dir(exec) + filepath.FromSlash("/frontend"))
	srv := httpServer()
	defer srv.Close()
	db, err := sqlx.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	err = storage.InitStorage(db)
	if err == nil {
		window_controllers.NewWindowController(db).SetupWindow(srv)
	} else {
		fmt.Println(err)
	}
}
