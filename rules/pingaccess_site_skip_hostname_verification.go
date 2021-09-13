package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// SiteSkipHostnameVerificationRule checks whether ...
type SiteSkipHostnameVerificationRule struct {
	resourceType string
}

// NewSiteSkipHostnameVerificationRule returns a new rule
func NewSiteSkipHostnameVerificationRule() *SiteSkipHostnameVerificationRule {
	return &SiteSkipHostnameVerificationRule{
		resourceType: "pingaccess_site",
	}
}

// Name returns the rule name
func (r *SiteSkipHostnameVerificationRule) Name() string {
	return "pingaccess_site_skip_hostname_verification_check"
}

// Enabled returns whether the rule is enabled by default
func (r *SiteSkipHostnameVerificationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *SiteSkipHostnameVerificationRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *SiteSkipHostnameVerificationRule) Link() string {
	return referenceLink(r.Name())
}

// Check checks whether ...
func (r *SiteSkipHostnameVerificationRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, "skip_hostname_verification", func(attribute *hcl.Attribute) error {
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
