package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ApplicationRequireHttpsEnabledRule checks whether ...
type ApplicationRequireHttpsEnabledRule struct {
	resourceType string
}

// NewApplicationRequireHttpsEnabledRule returns a new rule
func NewApplicationRequireHttpsEnabledRule() *ApplicationRequireHttpsEnabledRule {
	return &ApplicationRequireHttpsEnabledRule{
		resourceType: "pingaccess_application",
	}
}

// Name returns the rule name
func (r *ApplicationRequireHttpsEnabledRule) Name() string {
	return "pingaccess_application_requite_https_enabled_check"
}

// Enabled returns whether the rule is enabled by default
func (r *ApplicationRequireHttpsEnabledRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ApplicationRequireHttpsEnabledRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ApplicationRequireHttpsEnabledRule) Link() string {
	return referenceLink(r.Name())
}

// Check checks whether ...
func (r *ApplicationRequireHttpsEnabledRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, "require_https", func(attribute *hcl.Attribute) error {
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)
		if value == "false" {
			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf("require_https is %s", value),
					attribute.Expr,
				)
			})
		}
		return nil
	})
}
