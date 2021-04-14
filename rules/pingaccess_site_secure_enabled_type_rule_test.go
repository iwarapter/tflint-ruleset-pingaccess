package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_SiteSecureEnabledTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit false is error",
			Content: `
resource "pingaccess_site" "example1" {
	name                         = "example1"
	targets                      = ["www.google.com"]
	max_connections              = -1
	max_web_socket_connections   = -1
	availability_profile_id      = 1
	use_target_host_header     	 = false
	secure                       = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewSiteSecureEnabledTypeRule(),
					Message: "secure is false",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 9, Column: 33},
						End:      hcl.Pos{Line: 9, Column: 38},
					},
				},
			},
		},
		{
			Name: "explicit true is no error",
			Content: `
resource "pingaccess_site" "example1" {
	name                         = "example1"
	targets                      = ["www.google.com"]
	max_connections              = -1
	max_web_socket_connections   = -1
	availability_profile_id      = 1
	use_target_host_header     	 = false
	secure						 = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default false is error",
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

	rule := NewSiteSecureEnabledTypeRule()

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
