package main

import (
	"github.com/riverzhou/lute"
)

func md2html(md string)(html string) {
	luteEngine := lute.New()
	html = luteEngine.MarkdownStr("", md)
	return
}

func main(){}

