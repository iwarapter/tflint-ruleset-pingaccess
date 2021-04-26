package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_WebsessionSecureCookieEnabledRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "explicit false is error",
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
  secure_cookie = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewWebsessionSecureCookieEnabledRule(),
					Message: "secure_cookie is false",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 12, Column: 19},
						End:      hcl.Pos{Line: 12, Column: 24},
					},
				},
			},
		},
		{
			Name: "explicit true is no error",
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
  secure_cookie = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default true is no error",
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
			Expected: helper.Issues{},
		},
	}

	rule := NewWebsessionSecureCookieEnabledRule()

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
