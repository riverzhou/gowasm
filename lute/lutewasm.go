package main

import (
	"github.com/riverzhou/lute"
)

func main() {
	luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
	html:= luteEngine.MarkdownStr("demo", "**Lute** - A structured markdown engine.")
	//html := "<p><strong>Lute</strong> - A structured Markdown engine.</p>"
	println(html)
	// <p><strong>Lute</strong> - A structured Markdown engine.</p>
}

