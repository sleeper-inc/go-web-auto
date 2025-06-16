package steps

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"time"
	"web-auto-go-framework/config"
	"web-auto-go-framework/engine"
)

var e *engine.Action
var cfg *config.Config

func visitUrl(url string) error {
	return e.Visit(url)
}

func click(locator string) error {
	return e.Click(locator)
}

func fill(locator string, text string) error {
	return e.Fill(locator, text)
}

func assertTextPresent(expectedText string) error {
	if expectedText == "" {
		return fmt.Errorf("expected text is empty")
	}

	// Wait until the text appears (timeout 10s)
	el, err := e.Page.Timeout(10*time.Second).ElementR("body", expectedText)
	if err != nil {
		// Optional: dump full body text for debugging
		body, _ := e.Page.Element("body")
		text, _ := body.Text()
		fmt.Printf("Body text:\n%s\n", text)
		return fmt.Errorf("text not found: %s", expectedText)
	}

	fmt.Printf("Text found: %s\n", expectedText)
	_ = el // we don’t need to use it, just confirm it exists
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	var err error
	cfg, err = config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		l := launcher.New()
		if cfg.Browser.BrowserPath != "" {
			l.Bin(cfg.Browser.BrowserPath)
		}
		l.Headless(cfg.Rod.Headless)

		url := l.MustLaunch()

		browser := rod.New().ControlURL(url).MustConnect()

		e = &engine.Action{
			Page:    browser.Timeout(cfg.Rod.TimeoutDuration()).MustPage("about:blank"),
			Locator: engine.NewLocator(),
		}
		return ctx, nil
	})

	ctx.Step(`^(?:I |i |User |user |Client |client )?visit "([^"]*)"$`, visitUrl)
	ctx.Step(`^(?:I |i |User |user |Client |client )?click "([^"]*)"$`, click)
	ctx.Step(`^(?:I |i |User |user |Client |client )?fill "([^"]*)" with "([^"]*)"$`, fill)
	ctx.Step(`^(?:I |i |User |user |Client |client )?assert text "([^"]*)"$`, assertTextPresent)
}
