package golang

import (
	"fmt"
	"regexp"
)

// embed regexp.Regexp in a new type so we can extend it
type myRegexp struct {
	*regexp.Regexp
}

// add a new method to our new regular expression type
func (r *myRegexp) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		// Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}

		captures[name] = match[i]

	}
	return captures
}

// an example regular expression
var myExp = myRegexp{regexp.MustCompile(`(?P<fst>\d+)\.(?P<sec>\d+).(?P<trd>\d+)`)}


func main() {
	var ret =myExp.FindStringSubmatchMap("1234.5678.9");
	fmt.Print(ret[`fst`]);
	fmt.Print(ret[`sec`]);
	fmt.Print(ret[`trd`]);
	fmt.Printf("%+v",ret )
}
