package lib

import (
	"fmt"
	"regexp"

	"github.com/olekukonko/ts"
)

func Substring(aString string, limit int) string {
	var length = len(aString)

	if length > limit {
		return aString[0:limit-3] + "..."
	}

	return aString
}

func FmtMessage(
	repoName string,
	branch string,
	hash string,
	status string,
	commitMessage string,
	authorName string,
	authorEmail string,
) string {
	size, _ := ts.GetSize()
	var message = ""
	if size.Col() < 150 {
		message = fmt.Sprintf(
			"%-25s %-10s %-10s %-10s",
			Substring(repoName, 25),
			Substring(status, 10),
			Substring(branch, 10),
			hash,
		)

	} else {
		re := regexp.MustCompile(`\r?\n`)
		message = fmt.Sprintf(
			"%-25s %-10s %-10s %-10s %-25s %15s <%s>",
			Substring(repoName, 25),
			Substring(status, 10),
			Substring(branch, 10),
			hash,
			Substring(re.ReplaceAllString(commitMessage, " "), 25),
			Substring(authorName, 15),
			Substring(authorEmail, 15),
		)
	}

	return message
}
