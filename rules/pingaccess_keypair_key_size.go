package rules

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// KeyPairKeySizeRule checks whether ...
type KeyPairKeySizeRule struct {
	resourceType string
}

// NewKeyPairKeySizeRule returns a new rule
func NewKeyPairKeySizeRule() *KeyPairKeySizeRule {
	return &KeyPairKeySizeRule{
		resourceType: "pingaccess_keypair",
	}
}

// Name returns the rule name
func (r *KeyPairKeySizeRule) Name() string {
	return "pingaccess_keypair_key_size_check"
}

// Enabled returns whether the rule is enabled by default
func (r *KeyPairKeySizeRule) Enabled() bool {
	return true
}

type keyPairKeySizeRuleConfig struct {
	RsaKeySizes []int `hcl:"rsa_key_sizes,optional"`
	EcKeySizes  []int `hcl:"ec_key_sizes,optional"`
	Enabled     bool  `hcl:"enabled,optional"`
}

// Severity returns the rule severity
func (r *KeyPairKeySizeRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *KeyPairKeySizeRule) Link() string {
	return referenceLink(r.Name())
}

// Check checks whether ...
func (r *KeyPairKeySizeRule) Check(runner tflint.Runner) error {
	config := keyPairKeySizeRuleConfig{
		RsaKeySizes: []int{2048, 4096},
		EcKeySizes:  []int{256, 384, 521},
	}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	return runner.WalkResources(r.resourceType, func(resource *configs.Resource) error {
		body, _, diags := resource.Config.PartialContent(
			&hcl.BodySchema{
				Attributes: []hcl.AttributeSchema{{Name: "key_size"}, {Name: "key_algorithm"}},
			})
		if diags.HasErrors() {
			return diags
		}
		if body == nil {
			return fmt.Errorf("unable to parse attributes")
		}
		if len(body.Attributes) == 0 {
			//imported keypair, no size/alg
			return nil
		}
		var size, alg string
		var attr *hcl.Attribute
		for key, value := range body.Attributes {
			switch key {
			case "key_size":
				attr = value
				if err := runner.EvaluateExpr(value.Expr, &size, nil); err != nil {
					return fmt.Errorf("unable to evaluate expression: %s", err)
				}
			case "key_algorithm":
				if err := runner.EvaluateExpr(value.Expr, &alg, nil); err != nil {
					return fmt.Errorf("unable to evaluate expression: %s", err)
				}
			}
		}
		if attr == nil {
			return fmt.Errorf("unable to parse attributes no key_size")
		}
		switch alg {
		case "RSA":
			for _, keySize := range config.RsaKeySizes {
				if size == strconv.Itoa(keySize) {
					return nil
				}
			}
			return runner.EmitIssueOnExpr(
				r,
				fmt.Sprintf("key_size is %s", size),
				attr.Expr,
			)
		case "EC":
			for _, keySize := range config.EcKeySizes {
				if size == strconv.Itoa(keySize) {
					return nil
				}
			}
			return runner.EmitIssueOnExpr(
				r,
				fmt.Sprintf("key_size is %s", size),
				attr.Expr,
			)
		}
		return nil
	})
}
