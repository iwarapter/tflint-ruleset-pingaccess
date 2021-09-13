package main

import (
	"github.com/iwarapter/tflint-ruleset-pingaccess/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var (
	version = ""
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "pingaccess",
			Version: version,
			Rules: []tflint.Rule{
				rules.NewKeyPairKeySizeRule(),
				rules.NewPingFederateRunetimeSingletonRule(),
				rules.NewSiteSkipHostnameVerificationRule(),
				rules.NewSiteSecureEnabledRule(),
				rules.NewSiteSendPaCookieRule(),
				rules.NewApplicationRequireHttpsEnabledRule(),
				rules.NewApplicationResourceAuditLevelONRule(),
				rules.NewThirdPartyServiceSecureEnabledRule(),
				rules.NewThirdPartyServiceSkipHostnameVerificationRule(),
				rules.NewTrustedCertificateGroupsSkipCertificateDateCheckRule(),
				rules.NewWebsessionAuthTypeRule(),
				rules.NewWebsessionSecureCookieRule(),
				rules.NewWebsessionPkceChallengeTypeRule(),
			},
		},
	})
}
