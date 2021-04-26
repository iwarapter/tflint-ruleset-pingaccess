package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ApplicationResourceAuditLevelONRule checks whether ...
type ApplicationResourceAuditLevelONRule struct {
	resourceType string
}

// NewApplicationResourceAuditLevelONRule returns a new rule
func NewApplicationResourceAuditLevelONRule() *ApplicationResourceAuditLevelONRule {
	return &ApplicationResourceAuditLevelONRule{
		resourceType: "pingaccess_application_resource",
	}
}

// Name returns the rule name
func (r *ApplicationResourceAuditLevelONRule) Name() string {
	return "pingaccess_application_resource_audit_level_on_check"
}

// Enabled returns whether the rule is enabled by default
func (r *ApplicationResourceAuditLevelONRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ApplicationResourceAuditLevelONRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ApplicationResourceAuditLevelONRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *ApplicationResourceAuditLevelONRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, "audit_level", func(attribute *hcl.Attribute) error {
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)
		if value == "OFF" {
			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf("audit_level is %s", value),
					attribute.Expr,
				)
			})
		}
		return nil
	})
}
