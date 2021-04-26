package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ApplicationRequireHttpsEnabledRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit false is error",
			Content: `
resource "pingaccess_application" "demo" {
  application_type  = "API"
  name              = "api-demo"
  context_root      = "/"
  default_auth_type = "API"
  destination       = "Site"
  site_id           = pingaccess_site.example1.id
  virtual_host_ids  = [pingaccess_virtualhost.demo.id]
  require_https     = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewApplicationRequireHttpsEnabledRule(),
					Message: "require_https is false",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 10, Column: 23},
						End:      hcl.Pos{Line: 10, Column: 28},
					},
				},
			},
		},
		{
			Name: "explicit true is no error",
			Content: `
resource "pingaccess_application" "demo" {
  application_type  = "API"
  name              = "api-demo"
  context_root      = "/"
  default_auth_type = "API"
  destination       = "Site"
  site_id           = pingaccess_site.example1.id
  virtual_host_ids  = [pingaccess_virtualhost.demo.id]
  require_https     = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default false is error",
			Content: `
resource "pingaccess_application" "demo" {
  application_type  = "API"
  name              = "api-demo"
  context_root      = "/"
  default_auth_type = "API"
  destination       = "Site"
  site_id           = pingaccess_site.example1.id
  virtual_host_ids  = [pingaccess_virtualhost.demo.id]
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewApplicationRequireHttpsEnabledRule()

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
