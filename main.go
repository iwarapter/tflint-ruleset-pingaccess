package main

import (
	"github.com/iwarapter/tflint-ruleset-pingaccess/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: tflint.RuleSet{
			Name:    "pingaccess",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewPingFederateRunetimeSingletonTypeRule(),
			},
		},
	})
}
