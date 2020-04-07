package utils

import (
	"bytes"
	"os"
	"os/exec"
)

const ShellToUse = "bash"

func ExecGetOutput(command string, directory *string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	if directory != nil {
		cmd.Dir = *directory
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func Exec(command string, directory *string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	if directory != nil {
		cmd.Dir = *directory
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// required for things like ssh, kubectl-exec, and vim
func ExecNotCapturingOutput(command string, args []string, directory *string) error {
	cmd := exec.Command(command, args...)
	if directory != nil {
		cmd.Dir = *directory
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
