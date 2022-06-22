//go:build linux || darwin
// +build linux darwin

package core

/*
#cgo CFLAGS: -I./lwip/src/include
#include "lwip/init.h"
*/
import "C"

func lwipInit() {
	C.lwip_init() // Initialze modules.
}
