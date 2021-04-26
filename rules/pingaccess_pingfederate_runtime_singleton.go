package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// PingFederateRunetimeSingletonRule checks whether ...
type PingFederateRunetimeSingletonRule struct {
	resourceType string
}

// NewPingFederateRunetimeSingletonRule returns a new rule
func NewPingFederateRunetimeSingletonRule() *PingFederateRunetimeSingletonRule {
	return &PingFederateRunetimeSingletonRule{
		resourceType: "pingaccess_pingfederate_runtime",
	}
}

// Name returns the rule name
func (r *PingFederateRunetimeSingletonRule) Name() string {
	return "pingaccess_pingfederate_runtime_duplicate"
}

// Enabled returns whether the rule is enabled by default
func (r *PingFederateRunetimeSingletonRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *PingFederateRunetimeSingletonRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *PingFederateRunetimeSingletonRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *PingFederateRunetimeSingletonRule) Check(runner tflint.Runner) error {
	count := 0

	return runner.WalkResources(r.resourceType, func(res *configs.Resource) error {
		count++

		if count > 1 {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("duplicate instance of %s", res.Type),
				res.DeclRange,
			)
		}
		return nil
	})
}
