// +build darwin

package term

import (
	"syscall"
	"unsafe"
)

const IOSSIOSPEED = 0x80045402 // in ioss.h

func (t Term) SetHighSpeed(baud int) error {
	_, _, e := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(t.fd),
		IOSSIOSPEED,
		uintptr(unsafe.Pointer(&baud)))
	if e != 0 {
		return e
	}
	return nil
}
