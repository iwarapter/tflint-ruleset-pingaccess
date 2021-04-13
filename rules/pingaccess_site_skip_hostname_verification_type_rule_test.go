package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_SiteSkipHostnameVerificationTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit true is error",
			Content: `
resource "pingaccess_site" "example1" {
	name                         = "example1"
	targets                      = ["www.google.com"]
	max_connections              = -1
	max_web_socket_connections   = -1
	availability_profile_id      = 1
	skip_hostname_verification   = true
	use_target_host_header     	 = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewSiteSkipHostnameVerificationTypeRule(),
					Message: "skip_hostname_verification is true",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 8, Column: 33},
						End:      hcl.Pos{Line: 8, Column: 37},
					},
				},
			},
		},
		{
			Name: "explicit false is no error",
			Content: `
resource "pingaccess_site" "example1" {
	name                         = "example1"
	targets                      = ["www.google.com"]
	max_connections              = -1
	max_web_socket_connections   = -1
	availability_profile_id      = 1
	skip_hostname_verification   = false
	use_target_host_header     	 = false
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default false is no error",
			Content: `
resource "pingaccess_site" "example1" {
	name                         = "example1"
	targets                      = ["www.google.com"]
	max_connections              = -1
	max_web_socket_connections   = -1
	availability_profile_id      = 1
	use_target_host_header     	 = false
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewSiteSkipHostnameVerificationTypeRule()

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
