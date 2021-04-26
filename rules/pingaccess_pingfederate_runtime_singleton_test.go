package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_SingletonExampleType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "pingaccess_pingfederate_runtime" "foo" {
    issuer = "https://foo"
}

resource "pingaccess_pingfederate_runtime" "bar" {
    issuer = "https://bar"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewPingFederateRunetimeSingletonRule(),
					Message: "duplicate instance of pingaccess_pingfederate_runtime",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 1},
						End:      hcl.Pos{Line: 6, Column: 49},
					},
				},
			},
		},
	}

	rule := NewPingFederateRunetimeSingletonRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
