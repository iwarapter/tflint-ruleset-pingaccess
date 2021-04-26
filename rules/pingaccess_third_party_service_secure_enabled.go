package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ThirdPartyServiceSecureEnabledRule checks whether ...
type ThirdPartyServiceSecureEnabledRule struct {
	resourceType string
}

// NewThirdPartyServiceSecureEnabledRule returns a new rule
func NewThirdPartyServiceSecureEnabledRule() *ThirdPartyServiceSecureEnabledRule {
	return &ThirdPartyServiceSecureEnabledRule{
		resourceType: "pingaccess_third_party_service",
	}
}

// Name returns the rule name
func (r *ThirdPartyServiceSecureEnabledRule) Name() string {
	return "pingaccess_third_party_service_secure_enabled_check"
}

// Enabled returns whether the rule is enabled by default
func (r *ThirdPartyServiceSecureEnabledRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ThirdPartyServiceSecureEnabledRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ThirdPartyServiceSecureEnabledRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *ThirdPartyServiceSecureEnabledRule) Check(runner tflint.Runner) error {
	return runner.WalkResources(r.resourceType, func(resource *configs.Resource) error {
		attrs, diags := resource.Config.JustAttributes()
		if diags.HasErrors() {
			return diags
		}
		for _, attribute := range attrs {
			if attribute.Name == "secure" {
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
				} else {
					return nil
				}
			}
		}
		return runner.EmitIssue(
			r,
			"secure is unset for this resource, default is false",
			resource.DeclRange,
		)
	})
	//return runner.WalkResourceAttributes("pingaccess_third_party_service", "secure", func(attribute *hcl.Attribute) error {
	//	var value string
	//	err := runner.EvaluateExpr(attribute.Expr, &value, nil)
	//	if value == "false" {
	//		return runner.EnsureNoError(err, func() error {
	//			return runner.EmitIssueOnExpr(
	//				r,
	//				fmt.Sprintf("secure is %s", value),
	//				attribute.Expr,
	//			)
	//		})
	//	}
	//	return nil
	//})
}
