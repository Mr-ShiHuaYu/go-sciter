package cmdutil

import "os/exec"

func Exec(name string, arg ...string) (string, error) {
	c := exec.Command(name, arg...)
	output, err := c.CombinedOutput()
	return string(output), err
}

func ExecCmd(cmd string) (string, error) {
	c := exec.Command("cmd", "/C", cmd)
	output, err := c.CombinedOutput()
	return string(output), err
}
