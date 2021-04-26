package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_EngineListenderSecureEnabledRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit false is error",
			Content: `
resource "pingaccess_engine_listener" "acc_test" {
  name   = "example"
  port   = 443
  secure = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewEngineListenderSecureEnabledRule(),
					Message: "secure is false",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 12},
						End:      hcl.Pos{Line: 5, Column: 17},
					},
				},
			},
		},
		{
			Name: "explicit true is no error",
			Content: `
resource "pingaccess_engine_listener" "acc_test" {
  name   = "example"
  port   = 443
  secure = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default true is no error",
			Content: `
resource "pingaccess_engine_listener" "acc_test" {
  name   = "example"
  port   = 443
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewEngineListenderSecureEnabledRule()

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
