package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	if err != nil {
		fmt.Println(fmt.Errorf("Regex could not be compiled"))
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[~*=-]*>`)
	if err != nil {
		fmt.Println(fmt.Errorf("Regex could not be compiled"))
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`".*(?i)password.*"`)
	counter := 0
	if err != nil {
		fmt.Println(fmt.Errorf("Regex could not be compiled"))
	}

	for _, v := range lines {
		if re.MatchString(v) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(` end-of-line[0-9]+`)
	return re.ReplaceAllString(text, " ")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User\s+([A-Za-z0-9]+)\s`)
	jojo := make([]string, 0)
	if err != nil {
		fmt.Println(fmt.Errorf("Regex could not be compiled %s", err))
	}
	for _, v := range lines {
		match := re.FindStringSubmatch(v)
		if match != nil {
			jojo = append(jojo, fmt.Sprintf("[USR] %s %s", match[1], v))
		} else {
			jojo = append(jojo, v)
		}
	}
	return jojo
}
