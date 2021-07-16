package rules

import "fmt"

var (
	version = ""
)

// ReferenceLink returns the rule reference link
func referenceLink(name string) string {
	return fmt.Sprintf("https://github.com/iwarapter/tflint-ruleset-pingaccess/blob/%s/docs/rules/%s.md", version , name)
}