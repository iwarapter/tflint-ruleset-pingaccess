package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ThirdPartyServiceSecureEnabledRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit false is error",
			Content: `
resource "pingaccess_third_party_service" "example" {
	name = "example"
	targets = ["localhost:1234"]
	secure = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewThirdPartyServiceSecureEnabledRule(),
					Message: "secure is false",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 11},
						End:      hcl.Pos{Line: 5, Column: 16},
					},
				},
			},
		},
		{
			Name: "explicit true is no error",
			Content: `
resource "pingaccess_third_party_service" "example" {
	name = "example"
	targets = ["localhost:1234"]
	secure = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default false is error",
			Content: `
resource "pingaccess_third_party_service" "example" {
	name = "example"
	targets = ["localhost:1234"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewThirdPartyServiceSecureEnabledRule(),
					Message: "secure is unset for this resource, default is false",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 52},
					},
				},
			},
		},
	}

	rule := NewThirdPartyServiceSecureEnabledRule()

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
