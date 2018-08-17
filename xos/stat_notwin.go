// +build !windows

package xos

import (
	"os"
	"syscall"
)

// SysStat retunrs Stat_t struct for given file.
func SysStat(name string) (*syscall.Stat_t, error) {
	stat, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	return stat.Sys().(*syscall.Stat_t), nil
}
