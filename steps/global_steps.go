package steps

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"strconv"
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

func rightClick(locator string) error {
	return e.RightClick(locator)
}

func clickMultipleTimes(locator string, timesStr string) error {
	times, err := strconv.Atoi(timesStr)
	if err != nil {
		return fmt.Errorf("click time not valid: '%s'. value must be number", timesStr)
	}

	if times <= 0 {
		return fmt.Errorf("click must me more than 0")
	}

	return e.ClickMultipleTimes(locator, times)
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

	_ = el
	return nil
}

func seeTextEqualInElement(locator, expectedText string) error {
	return e.IsTextEqualInElement(locator, expectedText)
}

func isTextContains(expectedText string) error {
	return e.IsTextContains(expectedText)
}

func isTextContainsInElement(locator, expectedText string) error {
	return e.IsTextContainsInElement(locator, expectedText)
}

func isTextMatchingRegexInElement(locator, pattern string) error {
	return e.AssertTextMatchingRegexInElement(locator, pattern)
}

func isTextMatchingRegex(pattern string) error {
	return e.AssertTextMatchingRegex(pattern)
}

func hover(locator string) error {
	return e.Hover(locator)
}

func scrollToElement(locator string) error {
	return e.ScrollToElement(locator)
}

func scrollInDirection(direction, timeStr string) error {
	times := 1
	var err error

	if timeStr != "" {
		times, err = strconv.Atoi(timeStr)
		if err != nil {
			return fmt.Errorf("invalid time value '%s', must be a number", timeStr)
		}
	}
	return e.ScrollInDirection(direction, times)
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
	ctx.Step(`^(?:I |i |User |user |Client |client )?right click "([^"]*)"$`, rightClick)
	ctx.Step(`^(?:I |i |User |user |Client |client )?click "([^"]*)" "([^"]*)" times$`, clickMultipleTimes)
	ctx.Step(`^(?:I |i |User |user |Client |client )?fill "([^"]*)" with "([^"]*)"$`, fill)
	ctx.Step(`^(?:I |i |User |user |Client |client )?assert text "([^"]*)"$`, assertTextPresent)
	ctx.Step(`^(?:I |i |User |user |Client |client )?see text on "([^"]*)" equal to "([^"]*)"$`, seeTextEqualInElement)
	ctx.Step(`^(?:I |i |User |user |Client |client )?see text contains "([^"]*)"$`, isTextContains)
	ctx.Step(`^(?:I |i |User |user |Client |client )?see text contains "([^"]*)" inside "([^"]*)"$`, isTextContainsInElement)
	ctx.Step(`^(?:I |i |User |user |Client |client )?see text matching regex "([^"]*)"$`, isTextMatchingRegex)
	ctx.Step(`^(?:I |i |User |user |Client |client )?see text matching regex "([^"]*)" inside "([^"]*)"$`, isTextMatchingRegexInElement)
	ctx.Step(`^(?:I |i |User |user |Client |client )?hover "([^"]*)"$`, hover)
	ctx.Step(`^(?:I |i |User |user |Client |client )?scroll to element "([^"]*)"$`, scrollToElement)
	ctx.Step(`^(?:I |i |User |user |Client |client )?scroll (up|down|left|right)(?: "([^"]*)" times)?$`, scrollInDirection)
}
