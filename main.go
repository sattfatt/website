package main

import (
	"fmt"

	"github.com/sattfatt/website/site/pages"
)

func main() {
	fmt.Printf("hello world")

	server := pages.InitializePages(pages.AllPages)

	server.Logger.Fatal(server.Start(":10000"))
}
