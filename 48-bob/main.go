package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(Hey("WHAT'S GOING ON?"))
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	match, _ := regexp.MatchString("[a-zA-Z]", remark)

	noSpace := strings.TrimSpace(remark)
	capRemark := strings.ToUpper(remark)

	if noSpace == "" {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(noSpace, "?") {
		if remark == capRemark && match {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}

	if remark == capRemark && match {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
