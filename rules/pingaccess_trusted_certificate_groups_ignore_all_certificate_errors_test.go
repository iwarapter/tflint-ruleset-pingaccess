package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_TrustedCertificateGroupsIgnoreAllCertificateErrorsRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit true is error",
			Content: `
resource "pingaccess_trusted_certificate_group" "example" {
  name = "example"
  use_java_trust_store = true
  ignore_all_certificate_errors = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTrustedCertificateGroupsIgnoreAllCertificateErrorsRule(),
					Message: "ignore_all_certificate_errors is true",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 35},
						End:      hcl.Pos{Line: 5, Column: 39},
					},
				},
			},
		},
		{
			Name: "explicit false is no error",
			Content: `
resource "pingaccess_trusted_certificate_group" "example" {
  name = "example"
  use_java_trust_store = true
  ignore_all_certificate_errors = false
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default false is no error",
			Content: `
resource "pingaccess_trusted_certificate_group" "example" {
  name = "example"
  use_java_trust_store = true
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewTrustedCertificateGroupsIgnoreAllCertificateErrorsRule()

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
