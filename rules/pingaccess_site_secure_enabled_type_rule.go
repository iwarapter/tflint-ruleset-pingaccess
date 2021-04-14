package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// SiteSecureEnabledTypeRule checks whether ...
type SiteSecureEnabledTypeRule struct {
	resourceType string
}

// NewSiteSecureEnabledTypeRule returns a new rule
func NewSiteSecureEnabledTypeRule() *SiteSecureEnabledTypeRule {
	return &SiteSecureEnabledTypeRule{
		resourceType: "pingaccess_site",
	}
}

// Name returns the rule name
func (r *SiteSecureEnabledTypeRule) Name() string {
	return "pingaccess_site_secure_check"
}

// Enabled returns whether the rule is enabled by default
func (r *SiteSecureEnabledTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *SiteSecureEnabledTypeRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *SiteSecureEnabledTypeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *SiteSecureEnabledTypeRule) Check(runner tflint.Runner) error {
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
