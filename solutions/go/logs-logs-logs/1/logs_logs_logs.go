package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var logCategory string
Loop:
	for _, v := range log {
		switch v {
		case '\u2757':
			logCategory = "recommendation"
			break Loop
		case '\U0001f50d':
			logCategory = "search"
			break Loop
		case '\u2600':
			logCategory = "weather"
			break Loop
		default:
			logCategory = "default"
		}
	}
	return logCategory

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeLog := []rune(log)
	for i, v := range runeLog {
		if v == oldRune {
			runeLog[i] = newRune
		}
	}
	return string(runeLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.

func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
		return true
	} else {
		return false
	}
}
