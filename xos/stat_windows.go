// +build windows

package xos

import (
	"os"
	"syscall"
)

// SysStat retunrs Wind32FileAttributeData struct for given file.
func SysStat(name string) (*syscall.Win32FileAttributeData, error) {
	stat, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	return stat.Sys().(*syscall.Win32FileAttributeData), nil
}
