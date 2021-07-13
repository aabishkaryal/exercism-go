// Package bob lets you talk to bob
package bob

import (
	"strings"
)

var alphabets string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Function Hey return what bob would say for the remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if strings.ToUpper(remark) == remark &&
		strings.ContainsAny(remark, alphabets) &&
		strings.HasSuffix(remark, "?") {
		return "Calm down, I know what I'm doing!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if strings.ToUpper(remark) == remark &&
		strings.ContainsAny(remark, alphabets) {
		return "Whoa, chill out!"
	} else if remark == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
