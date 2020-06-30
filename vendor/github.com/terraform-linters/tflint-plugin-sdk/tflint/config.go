package tflint

// Config is a TFLint configuration applied to the plugin.
// Currently, it is not expected that each plugin will reference this directly.
type Config struct {
	Rules map[string]*RuleConfig
}

// RuleConfig is a TFLint's rule configuration.
type RuleConfig struct {
	Name    string
	Enabled bool
}
