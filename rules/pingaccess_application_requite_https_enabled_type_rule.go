package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ApplicationRequireHttpsEnabledTypeCheck checks whether ...
type ApplicationRequireHttpsEnabledTypeCheck struct {
	resourceType string
}

// NewApplicationRequireHttpsEnabledTypeCheck returns a new rule
func NewApplicationRequireHttpsEnabledTypeCheck() *ApplicationRequireHttpsEnabledTypeCheck {
	return &ApplicationRequireHttpsEnabledTypeCheck{
		resourceType: "pingaccess_application",
	}
}

// Name returns the rule name
func (r *ApplicationRequireHttpsEnabledTypeCheck) Name() string {
	return "pingaccess_application_requite_https_enabled_check"
}

// Enabled returns whether the rule is enabled by default
func (r *ApplicationRequireHttpsEnabledTypeCheck) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ApplicationRequireHttpsEnabledTypeCheck) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ApplicationRequireHttpsEnabledTypeCheck) Link() string {
	return ""
}

// Check checks whether ...
func (r *ApplicationRequireHttpsEnabledTypeCheck) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("pingaccess_application", "require_https", func(attribute *hcl.Attribute) error {
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
