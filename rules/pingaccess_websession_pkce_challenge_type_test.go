package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_WebsessionPkceChallengeTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit OFF is error",
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
  pkce_challenge_type = "OFF"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewWebsessionPkceChallengeTypeRule(),
					Message: "pkce_challenge_type is OFF",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 12, Column: 25},
						End:      hcl.Pos{Line: 12, Column: 30},
					},
				},
			},
		},
		{
			Name: "explicit SHA256 is no error",
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
  pkce_challenge_type = "SHA256"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default OFF is no error",
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
			Expected: helper.Issues{
				{
					Rule:    NewWebsessionPkceChallengeTypeRule(),
					Message: "pkce_challenge_type is unset for this resource, default is OFF",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 43},
					},
				},
			},
		},
	}

	rule := NewWebsessionPkceChallengeTypeRule()

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
