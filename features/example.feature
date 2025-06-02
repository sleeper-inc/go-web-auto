Feature: Example Feature

  @Example
  Scenario: Visit web and validate text
    Given visit "https://example.com/"
    Then assert text "Example Domain"

  @Example
  Scenario: Use I - Visit web and validate text
    Given I visit "https://example.com/"
    Then I assert text "Example Domain"

  @Example
  Scenario: Use i - Visit web and validate text
    Given i visit "https://example.com/"
    Then i assert text "Example Domain"

  @Example
  Scenario: Use User - Visit web and validate text
    Given User visit "https://example.com/"
    Then User assert text "Example Domain"

  @Example
  Scenario: Use user - Visit web and validate text
    Given user visit "https://example.com/"
    Then user assert text "Example Domain"

  @Example
  Scenario: Use Client - Visit web and validate text
    Given Client visit "https://example.com/"
    Then Client assert text "Example Domain"

  @Example
  Scenario: Use client - Visit web and validate text
    Given client visit "https://example.com/"
    Then client assert text "Example Domain"
