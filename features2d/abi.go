// Code generated by wit-bindgen-go. DO NOT EDIT.

package features2d

import (
	"go.bytecodealliance.org/cm"
	"unsafe"
)

// DetectorResultShape is used for storage in variant or result types.
type DetectorResultShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(DetectorResult{})]byte
}
