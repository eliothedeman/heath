package web

import (
	"github.com/eliothedeman/heath/block"
	"github.com/eliothedeman/heath/util"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type user struct {
	vecty.Core
	publicKey block.PublicKey
}

type settings struct {
	vecty.Core
	user *user `vecty:"prop"`
}

func (s *settings) Render() vecty.ComponentOrHTML {
	var udata vecty.ComponentOrHTML
	if s.user == nil {

		k := util.GenerateKey()
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("settings"),
		),
		elem.Heading1(
			vecty.Text("Settings"),
		),
		udata,
	)
}
