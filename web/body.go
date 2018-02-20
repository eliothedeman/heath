package web

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type header struct {
	vecty.Core
}

func (*header) Render() vecty.ComponentOrHTML {
	return nil
}

type footer struct {
	vecty.Core
}

func (*footer) Render() vecty.ComponentOrHTML {
	return nil
}

type app struct {
	vecty.Core
	header
	current renderer
	footer
}

type renderer interface {
	Render() vecty.ComponentOrHTML
}

func (a *app) Render() vecty.ComponentOrHTML {
	return elem.Body(
		a.header.Render(),
		a.current.Render(),
		a.footer.Render(),
	)
}

type empty struct{}

func (*empty) Render() vecty.ComponentOrHTML {
	return nil
}

func Make() vecty.Component {
	a := &app{
		current: &settings{},
	}
	return a
}
