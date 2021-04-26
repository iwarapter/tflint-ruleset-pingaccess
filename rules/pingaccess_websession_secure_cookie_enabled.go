package rules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// WebsessionSecureCookieEnabledRule checks whether ...
type WebsessionSecureCookieEnabledRule struct {
	resourceType string
}

// NewWebsessionSecureCookieEnabledRule returns a new rule
func NewWebsessionSecureCookieEnabledRule() *WebsessionSecureCookieEnabledRule {
	return &WebsessionSecureCookieEnabledRule{
		resourceType: "pingaccess_websession",
	}
}

// Name returns the rule name
func (r *WebsessionSecureCookieEnabledRule) Name() string {
	return "pingaccess_websession_secure_cookie_check"
}

// Enabled returns whether the rule is enabled by default
func (r *WebsessionSecureCookieEnabledRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *WebsessionSecureCookieEnabledRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *WebsessionSecureCookieEnabledRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *WebsessionSecureCookieEnabledRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, "secure_cookie", func(attribute *hcl.Attribute) error {
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)
		if value == "false" {
			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf("secure_cookie is %s", value),
					attribute.Expr,
				)
			})
		}
		return nil
	})
}
