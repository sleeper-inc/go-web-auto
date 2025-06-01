package engine

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Action struct {
	Page    *rod.Page
	Locator *Locator
}

func (e *Action) Visit(url string) error {
	e.Page.MustNavigate(url)
	return nil
}

func (e *Action) Click(locator string) error {
	sel, _ := e.Locator.Get(locator)
	return e.Page.MustElement(sel).Click(proto.InputMouseButtonLeft, 1)
}

func (e *Action) Fill(locator string, text string) error {
	sel, _ := e.Locator.Get(locator)
	return e.Page.MustElement(sel).Input(text)
}

func (e *Action) IsTextPresent(expectedText string) bool {
	_, err := e.Page.ElementR("body", expectedText)
	return err == nil
}
