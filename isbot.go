package isbot

import (
	_ "embed"
	"encoding/json"
	"strings"

	regexp "github.com/dlclark/regexp2"
)

//go:generate go run ./generate/generate.go

var matchers []*regexp.Regexp

type definition struct {
	Pattern string `json:"pattern"`
}

//go:embed crawler-user-agents.json
var crawlerUserAgents []byte

//go:embed custom.json
var customUserAgents []byte

//go:embed user-agents-bots.txt
var userAgentsRaw string
var userAgents = strings.Split(userAgentsRaw, "\n")

func init() {
	var definitions []definition
	err := json.Unmarshal(crawlerUserAgents, &definitions)
	if err != nil {
		panic(err)
	}

	var customDefinitions []definition
	err = json.Unmarshal(customUserAgents, &customDefinitions)
	if err != nil {
		panic(err)
	}

	matchers = make([]*regexp.Regexp, len(customDefinitions)+len(definitions))
	for i, d := range customDefinitions {
		matcher, err := regexp.Compile(d.Pattern, regexp.IgnoreCase)
		if err != nil {
			panic(err)
		}
		matchers[i] = matcher
	}

	for i, d := range definitions {
		matcher, err := regexp.Compile(d.Pattern, regexp.IgnoreCase)
		if err != nil {
			panic(err)
		}
		matchers[i+len(customDefinitions)] = matcher
	}
}

// Check using only the regexes
func CheckRegex(userAgent string) bool {
	for _, m := range matchers {
		match, _ := m.MatchString(userAgent)
		if match {
			return true
		}
	}

	return false
}

// Check using the list of known user agent strings
func CheckList(userAgent string) bool {
	for _, ua := range userAgents {
		if userAgent == ua {
			return true
		}
	}

	return false
}

// Check using both methods
func Check(userAgent string) bool {
	for _, ua := range userAgents {
		if userAgent == ua {
			return true
		}
	}

	for _, m := range matchers {
		match, _ := m.MatchString(userAgent)
		if match {
			return true
		}
	}

	return false
}
