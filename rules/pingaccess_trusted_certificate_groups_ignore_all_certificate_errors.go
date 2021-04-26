package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TrustedCertificateGroupsIgnoreAllCertificateErrorsRule checks whether ...
type TrustedCertificateGroupsIgnoreAllCertificateErrorsRule struct {
	resourceType string
}

// NewTrustedCertificateGroupsIgnoreAllCertificateErrorsRule returns a new rule
func NewTrustedCertificateGroupsIgnoreAllCertificateErrorsRule() *TrustedCertificateGroupsIgnoreAllCertificateErrorsRule {
	return &TrustedCertificateGroupsIgnoreAllCertificateErrorsRule{
		resourceType: "pingaccess_trusted_certificate_group",
	}
}

// Name returns the rule name
func (r *TrustedCertificateGroupsIgnoreAllCertificateErrorsRule) Name() string {
	return "pingaccess_trusted_certificate_group_ignore_all_certificate_errors_check"
}

// Enabled returns whether the rule is enabled by default
func (r *TrustedCertificateGroupsIgnoreAllCertificateErrorsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TrustedCertificateGroupsIgnoreAllCertificateErrorsRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TrustedCertificateGroupsIgnoreAllCertificateErrorsRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *TrustedCertificateGroupsIgnoreAllCertificateErrorsRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, "ignore_all_certificate_errors", func(attribute *hcl.Attribute) error {
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)
		if value == "true" {
			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf("ignore_all_certificate_errors is %s", value),
					attribute.Expr,
				)
			})
		}
		return nil
	})
}
