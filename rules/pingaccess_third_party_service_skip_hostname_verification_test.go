package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ThirdPartyServiceSkipHostnameVerificationRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit true is error",
			Content: `
resource "pingaccess_third_party_service" "example" {
	name = "example"
	targets = ["localhost:1234"]
	skip_hostname_verification = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewThirdPartyServiceSkipHostnameVerificationRule(),
					Message: "skip_hostname_verification is true",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 31},
						End:      hcl.Pos{Line: 5, Column: 35},
					},
				},
			},
		},
		{
			Name: "explicit false is no error",
			Content: `
resource "pingaccess_third_party_service" "example" {
	name = "example"
	targets = ["localhost:1234"]
	skip_hostname_verification = false
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default false is no error",
			Content: `
resource "pingaccess_third_party_service" "example" {
	name = "example"
	targets = ["localhost:1234"]
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewThirdPartyServiceSkipHostnameVerificationRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}
			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
