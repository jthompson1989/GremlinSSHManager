package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ExeSSHPwd(host string, port int, username string) error {
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", username, host), "-p", fmt.Sprintf("%d", port))

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
