package say

import "github.com/loadimpact/k6/js/modules"

func init() {
	modules.Register("k6/x/say", new(SAY))
}

// SAY is the k6 SQL plugin.
type SAY struct{}

// Hello says hello to someone
func (*SAY) Hello(name string) string {
	return "Hello " + name
}
