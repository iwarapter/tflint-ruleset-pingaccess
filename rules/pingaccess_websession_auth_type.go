package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// WebsessionAuthTypeRule checks whether ...
type WebsessionAuthTypeRule struct {
	resourceType string
}

type websessionAuthTypeRuleConfig struct {
	AllowedCredentialTypes []string `hcl:"allowed_credential_types,optional"`
	Enabled                bool     `hcl:"enabled,optional"`
}

// NewWebsessionAuthTypeRule returns a new rule
func NewWebsessionAuthTypeRule() *WebsessionAuthTypeRule {
	return &WebsessionAuthTypeRule{
		resourceType: "pingaccess_websession",
	}
}

// Name returns the rule name
func (r *WebsessionAuthTypeRule) Name() string {
	return "pingaccess_websession_auth_type_check"
}

// Enabled returns whether the rule is enabled by default
func (r *WebsessionAuthTypeRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *WebsessionAuthTypeRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *WebsessionAuthTypeRule) Link() string {
	return referenceLink(r.Name())
}

// Check checks whether ...
func (r *WebsessionAuthTypeRule) Check(runner tflint.Runner) error {
	config := websessionAuthTypeRuleConfig{AllowedCredentialTypes: []string{}}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	return runner.WalkResources(r.resourceType, func(resource *configs.Resource) error {
		body, _, diags := resource.Config.PartialContent(
			&hcl.BodySchema{
				Blocks: []hcl.BlockHeaderSchema{
					{
						Type: "client_credentials",
					},
				},
			})
		if diags.HasErrors() {
			return diags
		}
		for _, block := range body.Blocks.OfType("client_credentials") {
			attrs, _ := block.Body.JustAttributes()
			for _, attribute := range attrs {
				if attribute.Name == "credentials_type" {
					var value string
					_ = runner.EvaluateExpr(attribute.Expr, &value, nil)
					for _, allowed := range config.AllowedCredentialTypes {
						if allowed == value {
							return nil
						}
					}
					return runner.EmitIssueOnExpr(
						r,
						fmt.Sprintf("credentials_type is %s, allowed types are %s", value, strings.Join(config.AllowedCredentialTypes, ", ")),
						attribute.Expr,
					)
				}
			}
			for _, allowed := range config.AllowedCredentialTypes {
				if allowed == "SECRET" {
					return nil
				}
			}
		}
		return runner.EmitIssue(
			r,
			"credentials_type is unset for this resource, default is SECRET",
			resource.DeclRange,
		)
	})
}
