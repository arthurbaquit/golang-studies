package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElements struct {
	name, text string
	elements   []HtmlElements
}

func (e *HtmlElements) String() string {
	return e.StringImpl(0)
}

func (e *HtmlElements) StringImpl(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, v := range e.elements {
		sb.WriteString(v.StringImpl(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElements
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName: rootName, root: HtmlElements{name: rootName}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElements{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
}
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElements{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	someWord := "hello world"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(someWord)
	sb.WriteString("</p>")
	println(sb.String())

	someWords := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range someWords {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	println(sb.String())

	// using HtmlBuilder
	hb := NewHtmlBuilder("ul")
	hb.AddChildFluent("li", "hello").AddChildFluent("li", "world")
	println(hb.String())
}
