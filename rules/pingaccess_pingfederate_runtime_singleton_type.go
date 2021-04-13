package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// PingFederateRunetimeSingletonTypeRule checks whether ...
type PingFederateRunetimeSingletonTypeRule struct {
	resourceType string
}

// NewPingFederateRunetimeSingletonTypeRule returns a new rule
func NewPingFederateRunetimeSingletonTypeRule() *PingFederateRunetimeSingletonTypeRule {
	return &PingFederateRunetimeSingletonTypeRule{
		resourceType: "pingaccess_pingfederate_runtime",
	}
}

// Name returns the rule name
func (r *PingFederateRunetimeSingletonTypeRule) Name() string {
	return "pingaccess_pingfederate_runtime_duplicate"
}

// Enabled returns whether the rule is enabled by default
func (r *PingFederateRunetimeSingletonTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *PingFederateRunetimeSingletonTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *PingFederateRunetimeSingletonTypeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *PingFederateRunetimeSingletonTypeRule) Check(runner tflint.Runner) error {
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
