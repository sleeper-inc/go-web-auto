package engine

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"strings"
	"time"
)

// --- CORE & NAVIGATION ---
// This section contains the core structs and fundamental actions for element retrieval and page navigation.
// GetElement is a crucial helper for finding elements, while Visit navigates to a new page.

type LocatorProvider interface {
	Get(name string) (string, error)
}

type Action struct {
	Page    *rod.Page
	Locator LocatorProvider
}

func (e *Action) GetElement(locator string) (*rod.Element, error) {
	sel, err := e.Locator.Get(locator)
	if err != nil {
		return nil, err
	}

	if sel == "" {
		return nil, fmt.Errorf("locator found '%s' but get empty value on YML file", locator)
	}

	var element *rod.Element
	if strings.HasPrefix(sel, "/") || strings.HasPrefix(sel, "(") {
		element, err = e.Page.ElementX(sel)
	} else {
		element, err = e.Page.Element(sel)
	}
	return element, nil
}

func (e *Action) Visit(url string) error {
	e.Page.MustNavigate(url)
	return nil
}

// --- END CORE & NAVIGATION ---

// --- CLICK ACTIONS ---
// This section provides various methods for simulating mouse clicks on web elements,
// including single left-clicks, right-clicks, and multiple clicks.

func (e *Action) clickElement(locator string, button proto.InputMouseButton, count int) error {
	element, err := e.GetElement(locator)
	if err != nil {
		return err
	}

	if err := element.ScrollIntoView(); err != nil {
		return err
	}

	return element.Click(button, count)
}

func (e *Action) Click(locator string) error {
	return e.clickElement(locator, proto.InputMouseButtonLeft, 1)
}

func (e *Action) RightClick(locator string) error {
	return e.clickElement(locator, proto.InputMouseButtonRight, 1)
}

func (e *Action) ClickMultipleTimes(locator string, times int) error {
	for i := 0; i < times; i++ {
		if err := e.Click(locator); err != nil {
			return fmt.Errorf("failed to click on attempt-%d for element '%s': %w", i+1, locator, err)
		}

		if times > 1 {
			time.Sleep(200 * time.Millisecond)
		}
	}
	return nil
}

// --- END CLICK ACTIONS ---

// --- INPUT ACTIONS ---
// This section handles actions related to user input, such as filling out form fields.

func (e *Action) Fill(locator string, text string) error {
	sel, _ := e.Locator.Get(locator)
	return e.Page.MustElement(sel).Input(text)
}

// --- END INPUT ACTIONS ---

// --- TEXT ASSERTION ACTIONS ---
// This section includes a comprehensive set of methods for asserting text content on the page.
// It can check for exact text, substrings, and regex patterns, either on the whole page or within specific elements.

func (e *Action) searchTextOnPage(pattern string) error {
	_, err := e.Page.Timeout(10*time.Second).ElementR("body", pattern)
	return err
}

func (e *Action) IsTextPresent(expectedText string) bool {
	_, err := e.Page.ElementR("body", expectedText)
	return err == nil
}

func (e *Action) IsTextEqualInElement(locator, expectedText string) error {
	element, err := e.GetElement(locator)
	if err != nil {
		return err
	}
	actualText, err := element.Text()
	if err != nil {
		return fmt.Errorf("failed get text from element '%s': %w", locator, err)
	}
	if actualText != expectedText {
		return fmt.Errorf("text element not equal, expected text: '%s', actual text: '%s'", expectedText, actualText)
	}
	return nil
}

func (e *Action) IsTextContains(substring string) error {
	err := e.searchTextOnPage(substring)
	if err != nil {
		return fmt.Errorf("element contains text '%s' not found: %w", substring, err)
	}
	return nil
}

func (e *Action) IsTextContainsInElement(locator, substring string) error {
	container, err := e.GetElement(locator)
	if err != nil {
		return err
	}

	_, err = container.ElementR("*", substring)
	if err != nil {
		return fmt.Errorf("text '%s' not found inside element '%s': %w", substring, locator, err)
	}
	return nil
}

func (e *Action) AssertTextMatchingRegex(pattern string) error {
	err := e.searchTextOnPage(pattern)
	if err != nil {
		return fmt.Errorf("failed to found text match with regex '%s': %w", pattern, err)
	}

	return nil
}

func (e *Action) AssertTextMatchingRegexInElement(locator, pattern string) error {
	container, err := e.GetElement(locator)
	if err != nil {
		return err
	}

	_, err = container.ElementR("*", pattern)
	if err != nil {
		return fmt.Errorf("text matching regex '%s' not found inside element '%s': %w", pattern, locator, err)
	}
	return nil
}

// --- END TEXT ASSERTION ACTIONS ---

// --- MOUSE & SCROLL ACTIONS ---
// This section contains actions related to mouse movement and page scrolling,
// such as hovering over elements, scrolling to a specific element, or scrolling in a general direction.

func (e *Action) Hover(locator string) error {
	sel, err := e.Locator.Get(locator)
	if err != nil {
		return fmt.Errorf("could not find locator '%s': %w", locator, err)
	}
	element, err := e.Page.Element(sel)
	if err != nil {
		return fmt.Errorf("element not found with selector '%s': %w", sel, err)
	}
	return element.Hover()
}

func (e *Action) ScrollToElement(locator string) error {
	element, err := e.GetElement(locator)
	if err != nil {
		return err
	}
	return element.ScrollIntoView()
}

func (e *Action) ScrollInDirection(direction string, times int) error {
	const scrollAmount = 100

	var pixelX, pixelY float64

	switch direction {
	case "down":
		pixelY = float64(scrollAmount)
	case "up":
		pixelY = -float64(scrollAmount)
	case "right":
		pixelX = float64(scrollAmount)
	case "left":
		pixelX = -float64(scrollAmount)
	default:
		return fmt.Errorf("invalid direction '%s'", direction)
	}

	for i := 0; i < times; i++ {
		if err := e.Page.Mouse.Scroll(pixelX, pixelY, 1); err != nil {
			return fmt.Errorf("failed to scroll to '%s' for attempt %d: %w", direction, i+1, err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

// --- END MOUSE & SCROLL ACTIONS ---

func (e *Action) ClickDropdown(locator string, option string) error {

	sel, _ := e.Locator.Get(locator)
	options := e.Page.MustElement(sel).MustElements("option")

	found := false
	for _, opt := range options {
		text := opt.MustText()
		if text == option {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("option with text %q not found in dropdown %q", option, locator)
	}

	e.Page.MustElement(sel).MustVisible()
	e.Page.MustElement(sel).MustSelect(option)
	return nil
}

func (e *Action) DragAndDrop(locatorFrom string, locatorTo string) error {
	selectorFrom, _ := e.Locator.Get(locatorFrom)
	selectorTo, _ := e.Locator.Get(locatorTo)

	e.Page.MustElement(selectorFrom).MustWaitVisible()
	e.Page.MustElement(selectorTo).MustWaitVisible()

	source := e.Page.MustElement(selectorFrom)
	target := e.Page.MustElement(selectorTo)

	from := source.MustShape().OnePointInside()
	to := target.MustShape().OnePointInside()

	e.Page.Mouse.MustMoveTo(from.X, from.Y)
	e.Page.Mouse.MustDown("left")
	e.Page.Mouse.MustMoveTo(to.X, to.Y)
	e.Page.Mouse.MustUp("left")

	return nil
}
