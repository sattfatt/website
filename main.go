package main

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sattfatt/website/src/pages"
	"net/http"
)

//go:embed static/*
var embeddedFS embed.FS

func main() {
	fmt.Printf("hell world")

	e := echo.New()

	e = pages.InitializePages(pages.AllPages, e)

	// setup the static file server for css and stuff like that
	e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(embeddedFS))))

	e.Logger.Fatal(e.Start(":10000"))
}
