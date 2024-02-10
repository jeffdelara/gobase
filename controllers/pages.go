package pagescontroller

import (
	"github.com/a-h/templ"
	pagesview "github.com/jeffdelara/gobase/views/pages"
)

func Home() *templ.ComponentHandler {
	return templ.Handler(pagesview.Home())
}

func About() *templ.ComponentHandler {
	return templ.Handler(pagesview.About())
}
