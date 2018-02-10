package main

import (
	"github.com/eliothedeman/heath/web"
	"github.com/gopherjs/vecty"
)

func main() {
	vecty.SetTitle("Heath")
	vecty.RenderBody(web.Make())
}
