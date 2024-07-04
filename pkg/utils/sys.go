package utils

import (
	"os"
)

const nameMaxLen = 128

func GetHostname() string {
	hostname, _ := os.Hostname()
	hostRune := []rune("[PAM] - " + hostname)

	if len(hostRune) <= nameMaxLen {
		return string(hostRune)
	}
	name := make([]rune, nameMaxLen)
	index := nameMaxLen / 2
	copy(name[:index], hostRune[:index])
	start := len(hostRune) - index
	copy(name[index:], hostRune[start:])
	return string(name)
}
