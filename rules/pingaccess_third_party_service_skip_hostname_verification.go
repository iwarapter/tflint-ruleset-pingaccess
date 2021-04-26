package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ThirdPartyServiceSkipHostnameVerificationRule checks whether ...
type ThirdPartyServiceSkipHostnameVerificationRule struct {
	resourceType string
}

// NewThirdPartyServiceSkipHostnameVerificationRule returns a new rule
func NewThirdPartyServiceSkipHostnameVerificationRule() *ThirdPartyServiceSkipHostnameVerificationRule {
	return &ThirdPartyServiceSkipHostnameVerificationRule{
		resourceType: "pingaccess_third_party_service",
	}
}

// Name returns the rule name
func (r *ThirdPartyServiceSkipHostnameVerificationRule) Name() string {
	return "pingaccess_third_party_service_skip_hostname_verification_check"
}

// Enabled returns whether the rule is enabled by default
func (r *ThirdPartyServiceSkipHostnameVerificationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ThirdPartyServiceSkipHostnameVerificationRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ThirdPartyServiceSkipHostnameVerificationRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *ThirdPartyServiceSkipHostnameVerificationRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("pingaccess_third_party_service", "skip_hostname_verification", func(attribute *hcl.Attribute) error {
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)
		if value == "true" {
			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf("skip_hostname_verification is %s", value),
					attribute.Expr,
				)
			})
		}
		return nil
	})
}
