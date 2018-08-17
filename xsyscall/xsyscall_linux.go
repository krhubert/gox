// +build linux

package xsyscall

import (
	"fmt"
	"os"
	"runtime"
	"syscall"

	"github.com/krhubert/gox/xstrings"
)

// We need different setns values for the different platforms and arch
// We are declaring the macro here because the SETNS syscall does not exist in th stdlib
var setNsMap = map[string]uintptr{
	"linux/386":   346,
	"linux/amd64": 308,
	"linux/arm":   374,
}

// Setns syscall implementation
func Setns(fd uintptr, flags uintptr) error {
	ns, exists := setNsMap[fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)]
	if !exists {
		return fmt.Errorf("unsupported platform %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	if _, _, err := syscall.RawSyscall(ns, fd, flags, 0); err != 0 {
		return err
	}

	return nil
}

// SetnsWithFileName open file and call Setns(fd, 0)
func SetnsWithFileName(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}

	return Setns(f.Fd(), 0)
}

// Utsname is fraindly warapper for syscall.Utsname with string as filed instead of cstring.
type Utsname struct {
	Sysname    string
	Nodename   string
	Release    string
	Version    string
	Machine    string
	Domainname string
}

// Uname is warper for syscall.Uname that convert syscall.Utsname to xsyscall.Utsname
func Uname() (*Utsname, error) {
	var u syscall.Utsname

	if err := syscall.Uname(&u); err != nil {
		return nil, err
	}

	return &Utsname{
		xstrings.CharsToString(u.Sysname[:]),
		xstrings.CharsToString(u.Nodename[:]),
		xstrings.CharsToString(u.Release[:]),
		xstrings.CharsToString(u.Version[:]),
		xstrings.CharsToString(u.Machine[:]),
		xstrings.CharsToString(u.Domainname[:]),
	}, nil

}
