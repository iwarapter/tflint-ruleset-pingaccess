package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TrustedCertificateGroupsSkipCertificateDateCheckRule checks whether ...
type TrustedCertificateGroupsSkipCertificateDateCheckRule struct {
	resourceType string
}

// NewTrustedCertificateGroupsSkipCertificateDateCheckRule returns a new rule
func NewTrustedCertificateGroupsSkipCertificateDateCheckRule() *TrustedCertificateGroupsSkipCertificateDateCheckRule {
	return &TrustedCertificateGroupsSkipCertificateDateCheckRule{
		resourceType: "pingaccess_trusted_certificate_group",
	}
}

// Name returns the rule name
func (r *TrustedCertificateGroupsSkipCertificateDateCheckRule) Name() string {
	return "pingaccess_trusted_certificate_group_skip_certificate_date_check"
}

// Enabled returns whether the rule is enabled by default
func (r *TrustedCertificateGroupsSkipCertificateDateCheckRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TrustedCertificateGroupsSkipCertificateDateCheckRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TrustedCertificateGroupsSkipCertificateDateCheckRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *TrustedCertificateGroupsSkipCertificateDateCheckRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("pingaccess_trusted_certificate_group", "skip_certificate_date_check", func(attribute *hcl.Attribute) error {
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)
		if value == "true" {
			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf("skip_certificate_date_check is %s", value),
					attribute.Expr,
				)
			})
		}
		return nil
	})
}
