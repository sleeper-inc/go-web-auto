package main

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/godogx/allure"
	"os"
	"testing"
	"web-auto-go-framework/steps"
)

func Test(t *testing.T) {
	allure.RegisterFormatter()

	opts := godog.Options{
		Output:   colors.Colored(os.Stdout),
		Format:   "allure",
		Paths:    []string{"features"},
		TestingT: t, // required `go test` will work
		Tags:     os.Getenv("TAGS"),
	}

	status := godog.TestSuite{
		Name:                "web e2e",
		ScenarioInitializer: steps.InitializeScenario,
		Options:             &opts,
	}.Run()

	if status != 0 {
		t.Fatalf("godog tests failed with status: %d", status)
	}
}
