package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// WebsessionPkceChallengeTypeRule checks whether ...
type WebsessionPkceChallengeTypeRule struct {
	resourceType string
}

// NewWebsessionPkceChallengeTypeRule returns a new rule
func NewWebsessionPkceChallengeTypeRule() *WebsessionPkceChallengeTypeRule {
	return &WebsessionPkceChallengeTypeRule{
		resourceType: "pingaccess_websession",
	}
}

// Name returns the rule name
func (r *WebsessionPkceChallengeTypeRule) Name() string {
	return "pingaccess_websession_pkce_challenge_type_check"
}

// Enabled returns whether the rule is enabled by default
func (r *WebsessionPkceChallengeTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *WebsessionPkceChallengeTypeRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *WebsessionPkceChallengeTypeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *WebsessionPkceChallengeTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResources(r.resourceType, func(resource *configs.Resource) error {
		attrs, diags := resource.Config.JustAttributes()
		if diags.HasErrors() && len(attrs) == 0{
			return diags
		}
		for _, attribute := range attrs {
			if attribute.Name == "pkce_challenge_type" {
				var value string
				err := runner.EvaluateExpr(attribute.Expr, &value, nil)
				if value == "OFF" {
					return runner.EnsureNoError(err, func() error {
						return runner.EmitIssueOnExpr(
							r,
							fmt.Sprintf("pkce_challenge_type is %s", value),
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
			"pkce_challenge_type is unset for this resource, default is OFF",
			resource.DeclRange,
		)
	})
}
