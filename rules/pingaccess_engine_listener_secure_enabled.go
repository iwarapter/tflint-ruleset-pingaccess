package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// EngineListenderSecureEnabledRule checks whether ...
type EngineListenderSecureEnabledRule struct {
	resourceType string
}

// NewEngineListenderSecureEnabledRule returns a new rule
func NewEngineListenderSecureEnabledRule() *EngineListenderSecureEnabledRule {
	return &EngineListenderSecureEnabledRule{
		resourceType: "pingaccess_engine_listener",
	}
}

// Name returns the rule name
func (r *EngineListenderSecureEnabledRule) Name() string {
	return "pingaccess_engine_listener_secure_check"
}

// Enabled returns whether the rule is enabled by default
func (r *EngineListenderSecureEnabledRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *EngineListenderSecureEnabledRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *EngineListenderSecureEnabledRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *EngineListenderSecureEnabledRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, "secure", func(attribute *hcl.Attribute) error {
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
