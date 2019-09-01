package components

import (
	"imagepredict/classify"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// PageView is a vecty.Component which represents the entire page.
type PageView struct {
	vecty.Core
	Labels []classify.Label
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Section(
			vecty.Markup(
				vecty.Class("ImgClassifierapp"),
			),

			p.renderHeader(),
			vecty.If(len(p.Labels) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		),

		p.renderInfo(),
	)
}

func (p *PageView) renderHeader() *vecty.HTML {
	return elem.Header(
		vecty.Markup(vecty.Class("header")),
		elem.Heading1(vecty.Text("todos")),
		elem.Input(vecty.Markup(vecty.Class("new-todo"), prop.Type("file"))),
		elem.Image(
			vecty.Markup(
				vecty.Class("image"),
				prop.Src("https://boygeniusreport.files.wordpress.com/2016/11/puppy-dog.jpg"),
				prop.Alt("Puppy"),
			),
		),
	)
}

func (p *PageView) renderFooter() *vecty.HTML {
	return elem.Footer(
		vecty.Markup(
			vecty.Class("footer"),
		),

		elem.Span(
			vecty.Markup(
				vecty.Class("todo-count"),
			),

			elem.Strong(
				vecty.Text("Strong"),
			),
			vecty.Text("Remain Text"),
		),
	)
}

func (p *PageView) renderInfo() *vecty.HTML {
	return elem.Footer(
		vecty.Markup(
			vecty.Class("info"),
		),

		elem.Paragraph(
			vecty.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			vecty.Text("Created by "),
			elem.Anchor(
				vecty.Markup(
					prop.Href("http://github.com/neelance"),
				),
				vecty.Text("Richard Musiol"),
			),
		),
		elem.Paragraph(
			vecty.Text("Part of "),
			elem.Anchor(
				vecty.Markup(
					prop.Href("http://todomvc.com"),
				),
				vecty.Text("TodoMVC"),
			),
		),
	)
}

func (p *PageView) renderItemList() *vecty.HTML {
	var items vecty.List
	for _, l := range p.Labels {
		items = append(items, vecty.Text(l.Label))
	}

	return elem.Section(
		vecty.Markup(
			vecty.Class("main"),
		),
		elem.UnorderedList(
			vecty.Markup(
				vecty.Class("todo-list"),
			),
			items,
		),
	)
}
