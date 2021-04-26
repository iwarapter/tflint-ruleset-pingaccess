package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ApplicationResourceAuditLevelONRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit OFF is error",
			Content: `
resource "pingaccess_application_resource" "demo" {
  name = "example"
  methods = ["*"]
  path_patterns {
    pattern = "/as/token.oauth2"
    type    = "WILDCARD"
  }
  path_patterns {
    pattern = "%s"
    type    = "WILDCARD"
  }
  path_prefixes = [
    "/as/token.oauth2",
    "%s"
  ]
  audit_level = "OFF"
  application_id = pingaccess_application.demo.id
}`,
			Expected: helper.Issues{
				{
					Rule:    NewApplicationResourceAuditLevelONRule(),
					Message: "audit_level is OFF",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 17, Column: 17},
						End:      hcl.Pos{Line: 17, Column: 22},
					},
				},
			},
		},
		{
			Name: "explicit ON is no error",
			Content: `
resource "pingaccess_application_resource" "demo" {
  name = "example"
  methods = ["*"]
  path_patterns {
    pattern = "/as/token.oauth2"
    type    = "WILDCARD"
  }
  path_patterns {
    pattern = "%s"
    type    = "WILDCARD"
  }
  path_prefixes = [
    "/as/token.oauth2",
    "%s"
  ]
  audit_level = "ON"
  application_id = pingaccess_application.demo.id
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default ON is no error",
			Content: `
resource "pingaccess_application_resource" "demo" {
  name = "example"
  methods = ["*"]
  path_patterns {
    pattern = "/as/token.oauth2"
    type    = "WILDCARD"
  }
  path_patterns {
    pattern = "%s"
    type    = "WILDCARD"
  }
  path_prefixes = [
    "/as/token.oauth2",
    "%s"
  ]
  application_id = pingaccess_application.demo.id
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewApplicationResourceAuditLevelONRule()

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
