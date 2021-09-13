package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// SiteSendPaCookieRule checks whether ...
type SiteSendPaCookieRule struct {
	resourceType string
}

// NewSiteSendPaCookieRule returns a new rule
func NewSiteSendPaCookieRule() *SiteSendPaCookieRule {
	return &SiteSendPaCookieRule{
		resourceType: "pingaccess_site",
	}
}

// Name returns the rule name
func (r *SiteSendPaCookieRule) Name() string {
	return "pingaccess_site_send_pa_cookie_check"
}

// Enabled returns whether the rule is enabled by default
func (r *SiteSendPaCookieRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *SiteSendPaCookieRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *SiteSendPaCookieRule) Link() string {
	return referenceLink(r.Name())
}

// Check checks whether ...
func (r *SiteSendPaCookieRule) Check(runner tflint.Runner) error {
	return runner.WalkResources(r.resourceType, func(resource *configs.Resource) error {
		attrs, diags := resource.Config.JustAttributes()
		if diags.HasErrors() && len(attrs) == 0 {
			return diags
		}
		for _, attribute := range attrs {
			if attribute.Name == "send_pa_cookie" {
				var value string
				err := runner.EvaluateExpr(attribute.Expr, &value, nil)
				if value == "true" {
					return runner.EnsureNoError(err, func() error {
						return runner.EmitIssueOnExpr(
							r,
							fmt.Sprintf("send_pa_cookie is %s", value),
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
			"send_pa_cookie is unset for this resource, default is true",
			resource.DeclRange,
		)
	})
}
