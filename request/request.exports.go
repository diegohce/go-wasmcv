// Code generated by wit-bindgen-go. DO NOT EDIT.

package request

import (
	"github.com/hybridgroup/wasmcv/components/tinygo/cv/wasm/cv/mat"
)

// Exports represents the caller-defined exports from "wasm:cv/request".
var Exports struct {
	// Process represents the caller-defined, exported function "process".
	//
	//	process: func(image: mat) -> mat
	Process func(image mat.Mat) (result mat.Mat)
}
