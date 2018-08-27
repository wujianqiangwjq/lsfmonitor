package common

import (
	"os/exec"
	"strings"
)

func ExcuteWithoutpw(cmd string, args ...string) string {
	cmdinstance := exec.Command(cmd, args...)
	out, err := cmdinstance.CombinedOutput()
	if err != nil {
		return ""
	} else {
		return strings.TrimSpace(string(out))
	}
}
