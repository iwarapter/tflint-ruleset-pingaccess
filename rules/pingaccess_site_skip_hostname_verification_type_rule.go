package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// SiteSkipHostnameVerificationTypeRule checks whether ...
type SiteSkipHostnameVerificationTypeRule struct {
	resourceType string
}

// NewSiteSkipHostnameVerificationTypeRule returns a new rule
func NewSiteSkipHostnameVerificationTypeRule() *SiteSkipHostnameVerificationTypeRule {
	return &SiteSkipHostnameVerificationTypeRule{
		resourceType: "pingaccess_site",
	}
}

// Name returns the rule name
func (r *SiteSkipHostnameVerificationTypeRule) Name() string {
	return "pingaccess_site_skip_hostname_verification_check"
}

// Enabled returns whether the rule is enabled by default
func (r *SiteSkipHostnameVerificationTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *SiteSkipHostnameVerificationTypeRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *SiteSkipHostnameVerificationTypeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *SiteSkipHostnameVerificationTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("pingaccess_site", "skip_hostname_verification", func(attribute *hcl.Attribute) error {
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
