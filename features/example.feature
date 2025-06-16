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

  @Example2
  Scenario: Test Sauce Demo
    Given I visit "https://www.saucedemo.com/"
    When I fill "saucedemo.username_field" with "standard_user"
    And I fill "saucedemo.password_field" with "secret_sauce"
    And I click "saucedemo.login_button"
    Then I assert text "Swag Labs"
    And I see text contains "Swag"
    And I click "saucedemo.product_sorter"
    And I click "saucedemo.button_4"
    And I click "saucedemo.button_3" "3" times
    And I see text on "saucedemo.button_3" equal to "Remove"
    And I scroll up
    And I scroll down
    And I scroll left
    And I scroll right
    And I scroll up "3" times
    And I scroll down "3" times
    And I see text matching regex "9.\d{2}"
    And I scroll to element "saucedemo.facebook_icon"
    And I click "saucedemo.catalog_name_2"
