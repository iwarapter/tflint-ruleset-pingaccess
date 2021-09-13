package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_SiteSendPaCookieRule(t *testing.T) {
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
	use_target_host_header     	 = false
	send_pa_cookie               = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewSiteSendPaCookieRule(),
					Message: "send_pa_cookie is true",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 9, Column: 33},
						End:      hcl.Pos{Line: 9, Column: 37},
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
	use_target_host_header     	 = false
	send_pa_cookie				 = false
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default true is error",
			Content: `
resource "pingaccess_site" "example1" {
	name                         = "example1"
	targets                      = ["www.google.com"]
	max_connections              = -1
	max_web_socket_connections   = -1
	availability_profile_id      = 1
	use_target_host_header     	 = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewSiteSendPaCookieRule(),
					Message: "send_pa_cookie is unset for this resource, default is true",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 38},
					},
				},
			},
		},
	}

	rule := NewSiteSendPaCookieRule()

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
