package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_TrustedCertificateGroupsSkipCertificateDateCheckRule(t *testing.T) {
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
  skip_certificate_date_check = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTrustedCertificateGroupsSkipCertificateDateCheckRule(),
					Message: "skip_certificate_date_check is true",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 33},
						End:      hcl.Pos{Line: 5, Column: 37},
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
  skip_certificate_date_check = false
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

	rule := NewTrustedCertificateGroupsSkipCertificateDateCheckRule()

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
