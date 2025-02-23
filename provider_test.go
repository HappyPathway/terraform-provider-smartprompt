package main

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func TestAccDataSourceRefinedPrompt_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRefinedPromptConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.smartprompt_refined.test", "lazy_prompt", "what is terraform",
					),
					resource.TestCheckResourceAttrSet(
						"data.smartprompt_refined.test", "refined_prompt",
					),
				),
			},
		},
	})
}

func TestAccDataSourceRefinedPrompt_full(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRefinedPromptConfig_full,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.smartprompt_refined.test", "lazy_prompt", "explain kubernetes architecture",
					),
					resource.TestCheckResourceAttr(
						"data.smartprompt_refined.test", "domain", "architecture",
					),
					resource.TestCheckResourceAttr(
						"data.smartprompt_refined.test", "expertise_level", "expert",
					),
					resource.TestCheckResourceAttr(
						"data.smartprompt_refined.test", "output_format", "tutorial",
					),
					resource.TestCheckResourceAttrSet(
						"data.smartprompt_refined.test", "refined_prompt",
					),
					resource.TestCheckResourceAttrSet(
						"data.smartprompt_refined.test", "detected_topics.#",
					),
					resource.TestCheckResourceAttrSet(
						"data.smartprompt_refined.test", "recommended_references.#",
					),
				),
			},
		},
	})
}

func testAccPreCheck(t *testing.T) {
	// Add any required environment variables here
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"smartprompt": testAccProvider,
	}
}

const testAccDataSourceRefinedPromptConfig_basic = `
provider "smartprompt" {
  api_url = "http://localhost:8000"
  timeout = 30
}

data "smartprompt_refined" "test" {
  lazy_prompt = "what is terraform"
}
`

const testAccDataSourceRefinedPromptConfig_full = `
provider "smartprompt" {
  api_url = "http://localhost:8000"
  timeout = 30
}

data "smartprompt_refined" "test" {
  lazy_prompt            = "explain kubernetes architecture"
  domain                = "architecture"
  expertise_level       = "expert"
  output_format         = "tutorial"
  include_best_practices = true
  include_examples      = true
}
`