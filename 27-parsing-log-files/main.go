package main

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(ERR|BDG|INF|WRN|ERR|FTL)\].*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<[~*=-]*\>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`\".*(?i)password.*\"`)
	total := 0

	for _, v := range lines {
		if re.MatchString(v) {
			total += 1
		}
	}

	return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	result := make([]string, len(lines))
	for i, v := range lines {
		userName := re.FindStringSubmatch(v)
		if userName != nil {
			result[i] = fmt.Sprintf("[USR] %s %s", userName[1], v)
		} else {
			result[i] = v
		}
	}

	return result
}

func main() {
	fmt.Println(IsValidLine("[ERR] A good error here"))
	// => true
	fmt.Println(IsValidLine("Any old [ERR] text"))
	// => false
	fmt.Println(IsValidLine("[BOB] Any old text"))
	// => false
	fmt.Println(SplitLogLine("section 1<*>section 2<~~~>section 3"))
	// => []string{"section 1", "section 2", "section 3"},
	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	// => 2
	fmt.Println(lines)
	fmt.Println(RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27"))
	// => "[INF]  Network Failure "
	result := TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})
	// => []string {
	//  "[USR] James123 [WRN] User James123 has exceeded storage space.",
	//  "[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
	//  "[INF] Users can login again after 23:00.",
	//  "[DBG] We need to check that user names are at least 6 chars long."
	// }
	fmt.Println(result)
}
