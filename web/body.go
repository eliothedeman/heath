package web

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// TextBox lets you enter text
type TextBox struct {
	vecty.Core
	text string
}

func (t *TextBox) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Style("float", "right"),
			),
			elem.TextArea(
				vecty.Markup(
					vecty.Style("font-family", "monspace"),
					vecty.Property("rows", 14),
					vecty.Property("cols", 70),
				),
				vecty.Text(t.text),
			),
		),
	)
}

func Make() vecty.Component {
	return &TextBox{
		text: "hello",
	}
}
