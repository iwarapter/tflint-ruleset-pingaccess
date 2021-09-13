package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_KeyPairKeySizeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "RSA issue found",
			Content: `
resource "pingaccess_keypair" "test_generate" {
	alias = "acctest_test2"
	city = "Test"
	common_name = "Test"
	country = "GB"
	key_algorithm = "RSA"
	key_size = 1024
	organization = "Test"
	organization_unit = "Test"
	state = "Test"
	valid_days = 365
}`,
			Expected: helper.Issues{
				{
					Rule:    NewKeyPairKeySizeRule(),
					Message: "key_size is 1024",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 8, Column: 13},
						End:      hcl.Pos{Line: 8, Column: 17},
					},
				},
			},
		},
		{
			Name: "EC issue found",
			Content: `
resource "pingaccess_keypair" "test_generate" {
		alias = "acctest_test2"
		city = "Test"
		common_name = "Test"
		country = "GB"
		key_algorithm = "EC"
		key_size = 128
		organization = "Test"
		organization_unit = "Test"
		state = "Test"
		valid_days = 365
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewKeyPairKeySizeRule(),
					Message: "key_size is 128",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 8, Column: 14},
						End:      hcl.Pos{Line: 8, Column: 17},
					},
				},
			},
		},
		{
			Name: "imported key pairs are ignored",
			Content: `
resource "pingaccess_keypair" "test" {
	alias = "acctest_test"
	file_data = filebase64("test_cases/provider.p12")
	password = "password"
}`,
			Expected: helper.Issues{},
		},
	}
	rule := NewKeyPairKeySizeRule()
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
