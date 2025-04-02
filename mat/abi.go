// Code generated by wit-bindgen-go. DO NOT EDIT.

package mat

import (
	"go.bytecodealliance.org/cm"
	"unsafe"
	"wasmcv.org/wasm/cv/types"
)

// MixMaxLocResultShape is used for storage in variant or result types.
type MixMaxLocResultShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(MixMaxLocResult{})]byte
}

func lower_Size(v types.Size) (f0 uint32, f1 uint32) {
	f0 = (uint32)(v.X)
	f1 = (uint32)(v.Y)
	return
}

func lower_Rect(v types.Rect) (f0 uint32, f1 uint32, f2 uint32, f3 uint32) {
	f0, f1 = lower_Size(v.Min)
	f2, f3 = lower_Size(v.Max)
	return
}
