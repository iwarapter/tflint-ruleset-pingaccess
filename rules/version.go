package rules

import "fmt"

var (
	// RELEASE returns the release version
	RELEASE = "UNKNOWN"
	// REPO returns the git repository URL
	REPO = "UNKNOWN"
	// COMMIT returns the short sha from git
	COMMIT = "UNKNOWN"
)

// ReferenceLink returns the rule reference link
func referenceLink(name string) string {
	return fmt.Sprintf("https://github.com/iwarapter/tflint-ruleset-pingaccess/blob/%s/docs/rules/%s.md", RELEASE, name)
}