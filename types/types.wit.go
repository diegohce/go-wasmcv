// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package types represents the imported interface "wasm:cv/types".
package types

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Size represents the record "wasm:cv/types#size".
//
//	record size {
//		x: s32,
//		y: s32,
//	}
type Size struct {
	_ cm.HostLayout
	X int32
	Y int32
}

// Scalar represents the record "wasm:cv/types#scalar".
//
//	record scalar {
//		val1: f32,
//		val2: f32,
//		val3: f32,
//		val4: f32,
//	}
type Scalar struct {
	_    cm.HostLayout
	Val1 float32
	Val2 float32
	Val3 float32
	Val4 float32
}

// Rect represents the record "wasm:cv/types#rect".
//
//	record rect {
//		min: size,
//		max: size,
//	}
type Rect struct {
	_   cm.HostLayout
	Min Size
	Max Size
}

// Rgba represents the record "wasm:cv/types#rgba".
//
//	record rgba {
//		r: u8,
//		g: u8,
//		b: u8,
//		a: u8,
//	}
type Rgba struct {
	_ cm.HostLayout
	R uint8
	G uint8
	B uint8
	A uint8
}

// BorderType represents the enum "wasm:cv/types#border-type".
//
//	enum border-type {
//		border-constant,
//		border-replicate,
//		border-reflect,
//		border-wrap,
//		border-reflect101,
//		border-transparent,
//		border-default,
//		border-isolated
//	}
type BorderType uint8

const (
	BorderTypeBorderConstant BorderType = iota
	BorderTypeBorderReplicate
	BorderTypeBorderReflect
	BorderTypeBorderWrap
	BorderTypeBorderReflect101
	BorderTypeBorderTransparent
	BorderTypeBorderDefault
	BorderTypeBorderIsolated
)

var stringsBorderType = [8]string{
	"border-constant",
	"border-replicate",
	"border-reflect",
	"border-wrap",
	"border-reflect101",
	"border-transparent",
	"border-default",
	"border-isolated",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e BorderType) String() string {
	return stringsBorderType[e]
}

// AdaptiveThresholdType represents the enum "wasm:cv/types#adaptive-threshold-type".
//
//	enum adaptive-threshold-type {
//		adaptive-threshold-mean,
//		adaptive-threshold-gaussian
//	}
type AdaptiveThresholdType uint8

const (
	AdaptiveThresholdTypeAdaptiveThresholdMean AdaptiveThresholdType = iota
	AdaptiveThresholdTypeAdaptiveThresholdGaussian
)

var stringsAdaptiveThresholdType = [2]string{
	"adaptive-threshold-mean",
	"adaptive-threshold-gaussian",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e AdaptiveThresholdType) String() string {
	return stringsAdaptiveThresholdType[e]
}

// ThresholdType represents the enum "wasm:cv/types#threshold-type".
//
//	enum threshold-type {
//		threshold-binary,
//		threshold-binary-inv,
//		threshold-trunc,
//		threshold-to-zero,
//		threshold-to-zero-inv,
//		threshold-mask,
//		threshold-otsu,
//		tthreshold-triangle
//	}
type ThresholdType uint8

const (
	ThresholdTypeThresholdBinary ThresholdType = iota
	ThresholdTypeThresholdBinaryInv
	ThresholdTypeThresholdTrunc
	ThresholdTypeThresholdToZero
	ThresholdTypeThresholdToZeroInv
	ThresholdTypeThresholdMask
	ThresholdTypeThresholdOtsu
	ThresholdTypeTthresholdTriangle
)

var stringsThresholdType = [8]string{
	"threshold-binary",
	"threshold-binary-inv",
	"threshold-trunc",
	"threshold-to-zero",
	"threshold-to-zero-inv",
	"threshold-mask",
	"threshold-otsu",
	"tthreshold-triangle",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ThresholdType) String() string {
	return stringsThresholdType[e]
}

// DataLayoutType represents the enum "wasm:cv/types#data-layout-type".
//
//	enum data-layout-type {
//		data-layout-unknown,
//		data-layout-nd,
//		data-layout-nchw,
//		data-layout-ncdhw,
//		data-layout-nhwc,
//		data-layout-ndhwc,
//		data-layout-planar
//	}
type DataLayoutType uint8

const (
	DataLayoutTypeDataLayoutUnknown DataLayoutType = iota
	DataLayoutTypeDataLayoutNd
	DataLayoutTypeDataLayoutNchw
	DataLayoutTypeDataLayoutNcdhw
	DataLayoutTypeDataLayoutNhwc
	DataLayoutTypeDataLayoutNdhwc
	DataLayoutTypeDataLayoutPlanar
)

var stringsDataLayoutType = [7]string{
	"data-layout-unknown",
	"data-layout-nd",
	"data-layout-nchw",
	"data-layout-ncdhw",
	"data-layout-nhwc",
	"data-layout-ndhwc",
	"data-layout-planar",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e DataLayoutType) String() string {
	return stringsDataLayoutType[e]
}

// PaddingModeType represents the enum "wasm:cv/types#padding-mode-type".
//
//	enum padding-mode-type {
//		padding-mode-null,
//		padding-mode-crop-center,
//		padding-mode-letterbox
//	}
type PaddingModeType uint8

const (
	PaddingModeTypePaddingModeNull PaddingModeType = iota
	PaddingModeTypePaddingModeCropCenter
	PaddingModeTypePaddingModeLetterbox
)

var stringsPaddingModeType = [3]string{
	"padding-mode-null",
	"padding-mode-crop-center",
	"padding-mode-letterbox",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e PaddingModeType) String() string {
	return stringsPaddingModeType[e]
}

// BlobParams represents the record "wasm:cv/types#blob-params".
//
//	record blob-params {
//		scale-factor: f32,
//		size: size,
//		mean: scalar,
//		swap-rb: bool,
//		ddepth: u8,
//		data-layout: data-layout-type,
//		padding-mode: padding-mode-type,
//		border: scalar,
//	}
type BlobParams struct {
	_           cm.HostLayout
	ScaleFactor float32
	Size        Size
	Mean        Scalar
	SwapRb      bool
	Ddepth      uint8
	DataLayout  DataLayoutType
	PaddingMode PaddingModeType
	Border      Scalar
}

// MixMaxLocResult represents the record "wasm:cv/types#mix-max-loc-result".
//
//	record mix-max-loc-result {
//		min-val: f32,
//		max-val: f32,
//		min-loc: size,
//		max-loc: size,
//	}
type MixMaxLocResult struct {
	_      cm.HostLayout
	MinVal float32
	MaxVal float32
	MinLoc Size
	MaxLoc Size
}

// HersheyFontType represents the enum "wasm:cv/types#hershey-font-type".
//
//	enum hershey-font-type {
//		hershey-font-simplex,
//		hershey-font-plain,
//		hershey-font-duplex,
//		hershey-font-complex,
//		hershey-font-triplex,
//		hershey-font-complex-small,
//		hershey-font-script-simplex,
//		hershey-font-script-complex,
//		hershey-font-italic
//	}
type HersheyFontType uint8

const (
	HersheyFontTypeHersheyFontSimplex HersheyFontType = iota
	HersheyFontTypeHersheyFontPlain
	HersheyFontTypeHersheyFontDuplex
	HersheyFontTypeHersheyFontComplex
	HersheyFontTypeHersheyFontTriplex
	HersheyFontTypeHersheyFontComplexSmall
	HersheyFontTypeHersheyFontScriptSimplex
	HersheyFontTypeHersheyFontScriptComplex
	HersheyFontTypeHersheyFontItalic
)

var stringsHersheyFontType = [9]string{
	"hershey-font-simplex",
	"hershey-font-plain",
	"hershey-font-duplex",
	"hershey-font-complex",
	"hershey-font-triplex",
	"hershey-font-complex-small",
	"hershey-font-script-simplex",
	"hershey-font-script-complex",
	"hershey-font-italic",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e HersheyFontType) String() string {
	return stringsHersheyFontType[e]
}

// InterpolationType represents the enum "wasm:cv/types#interpolation-type".
//
//	enum interpolation-type {
//		interpolation-nearest,
//		interpolation-linear,
//		interpolation-cubic,
//		interpolation-area,
//		interpolation-lanczos4
//	}
type InterpolationType uint8

const (
	InterpolationTypeInterpolationNearest InterpolationType = iota
	InterpolationTypeInterpolationLinear
	InterpolationTypeInterpolationCubic
	InterpolationTypeInterpolationArea
	InterpolationTypeInterpolationLanczos4
)

var stringsInterpolationType = [5]string{
	"interpolation-nearest",
	"interpolation-linear",
	"interpolation-cubic",
	"interpolation-area",
	"interpolation-lanczos4",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e InterpolationType) String() string {
	return stringsInterpolationType[e]
}
