package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_WebsessionAuthTypeRule(t *testing.T) {
	cases := []struct {
		Name       string
		Content    string
		RuleConfig string
		Expected   helper.Issues
	}{
		{
			Name: "explicit SECRET is error",
			Content: `
resource "pingaccess_websession" "example" {
  name = "example"
  audience = "aud"
  client_credentials {
	client_id = "websession"
	credentials_type = "SECRET"
	client_secret {
		value = "top_secret"
	}
  }
  scopes = ["profile","address","email","phone"]
  pkce_challenge_type = "OFF"
}`,
			RuleConfig: `rule "pingaccess_websession_auth_type_check" {
  enabled = true
  allowed_credential_types = ["PRIVATE_KEY_JWT"]
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewWebsessionAuthTypeRule(),
					Message: "credentials_type is SECRET, allowed types are PRIVATE_KEY_JWT",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 7, Column: 21},
						End:      hcl.Pos{Line: 7, Column: 29},
					},
				},
			},
		},
		{
			Name: "explicit PRIVATE_KEY_JWT is no error",
			Content: `
resource "pingaccess_websession" "example" {
  name = "example"
  audience = "aud"
  client_credentials {
	client_id = "websession"
	credentials_type = "PRIVATE_KEY_JWT"
	client_secret {
		value = "top_secret"
	}
  }
  scopes = ["profile","address","email","phone"]
  pkce_challenge_type = "SHA256"
}`,
			RuleConfig: `rule "pingaccess_websession_auth_type_check" {
  enabled = true
  allowed_credential_types = ["PRIVATE_KEY_JWT"]
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "default SECRET is error",
			Content: `
resource "pingaccess_websession" "example" {
  name = "example"
  audience = "aud"
  client_credentials {
	client_id = "websession"
	client_secret {
		value = "top_secret"
	}
  }
  scopes = ["profile","address","email","phone"]
}`,
			RuleConfig: `rule "pingaccess_websession_auth_type_check" {
  enabled = true
  allowed_credential_types = ["PRIVATE_KEY_JWT"]
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewWebsessionAuthTypeRule(),
					Message: "credentials_type is unset for this resource, default is SECRET",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 43},
					},
				},
			},
		},
		{
			Name: "default SECRET is allowed",
			Content: `
resource "pingaccess_websession" "example" {
  name = "example"
  audience = "aud"
  client_credentials {
	client_id = "websession"
	client_secret {
		value = "top_secret"
	}
  }
  scopes = ["profile","address","email","phone"]
}`,
			RuleConfig: `rule "pingaccess_websession_auth_type_check" {
  enabled = true
  allowed_credential_types = ["SECRET","PRIVATE_KEY_JWT"]
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewWebsessionAuthTypeRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content, ".tflint.hcl": tc.RuleConfig})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}
			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
