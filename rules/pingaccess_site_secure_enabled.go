package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// SiteSecureEnabledRule checks whether ...
type SiteSecureEnabledRule struct {
	resourceType string
}

// NewSiteSecureEnabledRule returns a new rule
func NewSiteSecureEnabledRule() *SiteSecureEnabledRule {
	return &SiteSecureEnabledRule{
		resourceType: "pingaccess_site",
	}
}

// Name returns the rule name
func (r *SiteSecureEnabledRule) Name() string {
	return "pingaccess_site_secure_check"
}

// Enabled returns whether the rule is enabled by default
func (r *SiteSecureEnabledRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *SiteSecureEnabledRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *SiteSecureEnabledRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *SiteSecureEnabledRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("pingaccess_site", "secure", func(attribute *hcl.Attribute) error {
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)
		if value == "false" {
			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf("secure is %s", value),
					attribute.Expr,
				)
			})
		}
		return nil
	})
}
