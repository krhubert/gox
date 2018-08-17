package xexec

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Execute the specified command and waits for it to complete.
// It retruns standard output and standard error as error.
// The returned error is nil if the command runs, has no problems.
// If the command fails to run or doesn't complete successfully,
// standard error is stored in error.
func Execute(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("%s %s %s", name, arg, err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("%s %s %s", name, arg, err)
	}
	if err = cmd.Start(); err != nil {
		return "", fmt.Errorf("%s %s %s", name, arg, err)
	}

	be, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", fmt.Errorf("%s %s %s", name, arg, err)
	}

	bo, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("%s %s %s", name, arg, err)
	}
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("%s %s", name, string(be))
	}
	return string(bo), nil
}

// Run the specified command and waits for it to complete.
// The returned error is nil if the command runs, has no problems.
// If the command fails to run or doesn't complete successfully,
// standard error is stored in error.
func Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("%s %s %s", name, arg, err)
	}
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("%s %s %s", name, arg, err)
	}
	be, err := ioutil.ReadAll(stderr)
	if err != nil {
		return fmt.Errorf("%s %s %s", name, arg, err)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("%s %s", name, string(be))
	}
	return nil
}
