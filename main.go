package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/markbates/pkger"
	"github.com/webview/webview"
)

type pkgerServer struct{}

func (p *pkgerServer) Open(name string) (http.File, error) {
	return pkger.Open(name)
}

func httpServer() *httptest.Server {
	return httptest.NewServer(
		http.FileServer(
			&pkgerServer{}))
}

func main() {
	_ = pkger.Include("/frontend")

	srv := httpServer()
	defer srv.Close()

	url := fmt.Sprintf("%s/frontend/pages/main.html", srv.URL)
	fmt.Println(srv.URL)
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Ксюкулятор")
	w.SetSize(480, 320, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
