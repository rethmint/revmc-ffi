//go:build linux && !muslc && arm64

package vm

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lrevmapi.aarch64
import "C"