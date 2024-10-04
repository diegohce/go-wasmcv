// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package mat represents the imported interface "wasm:cv/mat".
//
// mat resource is a matrix of bytes, wrapper around the cv::Mat type.
package mat

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"wasmcv.org/wasm/cv/types"
)

// Mattype represents the enum "wasm:cv/mat#mattype".
//
//	enum mattype {
//		cv8u,
//		cv8s,
//		cv16u,
//		cv16s,
//		cv32s,
//		cv32f,
//		cv64f
//	}
type Mattype uint8

const (
	MattypeCv8u Mattype = iota
	MattypeCv8s
	MattypeCv16u
	MattypeCv16s
	MattypeCv32s
	MattypeCv32f
	MattypeCv64f
)

var stringsMattype = [7]string{
	"cv8u",
	"cv8s",
	"cv16u",
	"cv16s",
	"cv32s",
	"cv32f",
	"cv64f",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e Mattype) String() string {
	return stringsMattype[e]
}

// Mat represents the imported resource "wasm:cv/mat#mat".
//
//	resource mat
type Mat cm.Resource

// ResourceDrop represents the imported resource-drop for resource "mat".
//
// Drops a resource handle.
//
//go:nosplit
func (self Mat) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_MatResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/mat [resource-drop]mat
//go:noescape
func wasmimport_MatResourceDrop(self0 uint32)

// NewMat represents the imported constructor for resource "mat".
//
// Create a new Mat.
//
//	constructor()
//
//go:nosplit
func NewMat() (result Mat) {
	result0 := wasmimport_NewMat()
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [constructor]mat
//go:noescape
func wasmimport_NewMat() (result0 uint32)

// MatNewWithSize represents the imported static function "new-with-size".
//
// Create a new Mat with the specified size and type.
//
//	new-with-size: static func(cols: u32, rows: u32, mattype: mattype) -> mat
//
//go:nosplit
func MatNewWithSize(cols uint32, rows uint32, mattype Mattype) (result Mat) {
	cols0 := (uint32)(cols)
	rows0 := (uint32)(rows)
	mattype0 := (uint32)(mattype)
	result0 := wasmimport_MatNewWithSize((uint32)(cols0), (uint32)(rows0), (uint32)(mattype0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [static]mat.new-with-size
//go:noescape
func wasmimport_MatNewWithSize(cols0 uint32, rows0 uint32, mattype0 uint32) (result0 uint32)

// Clone represents the imported method "clone".
//
// Clone returns a cloned full copy of the Mat.
//
//	clone: func() -> mat
//
//go:nosplit
func (self Mat) Clone() (result Mat) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatClone((uint32)(self0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.clone
//go:noescape
func wasmimport_MatClone(self0 uint32) (result0 uint32)

// Close represents the imported method "close".
//
// Close the Mat
//
//	close: func()
//
//go:nosplit
func (self Mat) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_MatClose((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.close
//go:noescape
func wasmimport_MatClose(self0 uint32)

// ColRange represents the imported method "col-range".
//
// ColRange creates a matrix header for the specified column span.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#aadc8f9210fe4dec50513746c246fa8d9
//
//	col-range: func(start: u32, end: u32) -> mat
//
//go:nosplit
func (self Mat) ColRange(start uint32, end uint32) (result Mat) {
	self0 := cm.Reinterpret[uint32](self)
	start0 := (uint32)(start)
	end0 := (uint32)(end)
	result0 := wasmimport_MatColRange((uint32)(self0), (uint32)(start0), (uint32)(end0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.col-range
//go:noescape
func wasmimport_MatColRange(self0 uint32, start0 uint32, end0 uint32) (result0 uint32)

// Cols represents the imported method "cols".
//
// Cols returns the number of columns for this Mat.
//
//	cols: func() -> u32
//
//go:nosplit
func (self Mat) Cols() (result uint32) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatCols((uint32)(self0))
	result = (uint32)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.cols
//go:noescape
func wasmimport_MatCols(self0 uint32) (result0 uint32)

// CopyTo represents the imported method "copy-to".
//
// CopyTo copies Mat into destination Mat.
//
//	copy-to: func(dst: mat)
//
//go:nosplit
func (self Mat) CopyTo(dst Mat) {
	self0 := cm.Reinterpret[uint32](self)
	dst0 := cm.Reinterpret[uint32](dst)
	wasmimport_MatCopyTo((uint32)(self0), (uint32)(dst0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.copy-to
//go:noescape
func wasmimport_MatCopyTo(self0 uint32, dst0 uint32)

// Empty represents the imported method "empty".
//
// Empty returns true if the Mat is empty.
//
//	empty: func() -> bool
//
//go:nosplit
func (self Mat) Empty() (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatEmpty((uint32)(self0))
	result = cm.U32ToBool((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.empty
//go:noescape
func wasmimport_MatEmpty(self0 uint32) (result0 uint32)

// GetFloatAt represents the imported method "get-float-at".
//
// GetFloatAt returns the value at the specified row and column as a f32.
//
//	get-float-at: func(row: u32, col: u32) -> f32
//
//go:nosplit
func (self Mat) GetFloatAt(row uint32, col uint32) (result float32) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	result0 := wasmimport_MatGetFloatAt((uint32)(self0), (uint32)(row0), (uint32)(col0))
	result = (float32)((float32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-float-at
//go:noescape
func wasmimport_MatGetFloatAt(self0 uint32, row0 uint32, col0 uint32) (result0 float32)

// GetFloatAt3 represents the imported method "get-float-at3".
//
// GetFloatAt3 returns the value at the specified x, y, z as a f32.
//
//	get-float-at3: func(x: u32, y: u32, z: u32) -> f32
//
//go:nosplit
func (self Mat) GetFloatAt3(x uint32, y uint32, z uint32) (result float32) {
	self0 := cm.Reinterpret[uint32](self)
	x0 := (uint32)(x)
	y0 := (uint32)(y)
	z0 := (uint32)(z)
	result0 := wasmimport_MatGetFloatAt3((uint32)(self0), (uint32)(x0), (uint32)(y0), (uint32)(z0))
	result = (float32)((float32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-float-at3
//go:noescape
func wasmimport_MatGetFloatAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32) (result0 float32)

// GetIntAt represents the imported method "get-int-at".
//
// GetIntAt returns the value at the specified row and column as a s32.
//
//	get-int-at: func(row: u32, col: u32) -> s32
//
//go:nosplit
func (self Mat) GetIntAt(row uint32, col uint32) (result int32) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	result0 := wasmimport_MatGetIntAt((uint32)(self0), (uint32)(row0), (uint32)(col0))
	result = (int32)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-int-at
//go:noescape
func wasmimport_MatGetIntAt(self0 uint32, row0 uint32, col0 uint32) (result0 uint32)

// GetIntAt3 represents the imported method "get-int-at3".
//
// GetIntAt3 returns the value at the specified x, y, z as a s32.
//
//	get-int-at3: func(x: u32, y: u32, z: u32) -> s32
//
//go:nosplit
func (self Mat) GetIntAt3(x uint32, y uint32, z uint32) (result int32) {
	self0 := cm.Reinterpret[uint32](self)
	x0 := (uint32)(x)
	y0 := (uint32)(y)
	z0 := (uint32)(z)
	result0 := wasmimport_MatGetIntAt3((uint32)(self0), (uint32)(x0), (uint32)(y0), (uint32)(z0))
	result = (int32)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-int-at3
//go:noescape
func wasmimport_MatGetIntAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32) (result0 uint32)

// GetUcharAt represents the imported method "get-uchar-at".
//
// GetUCharAt returns the value at the specified row and column as a u8.
//
//	get-uchar-at: func(row: u32, col: u32) -> u8
//
//go:nosplit
func (self Mat) GetUcharAt(row uint32, col uint32) (result uint8) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	result0 := wasmimport_MatGetUcharAt((uint32)(self0), (uint32)(row0), (uint32)(col0))
	result = (uint8)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-uchar-at
//go:noescape
func wasmimport_MatGetUcharAt(self0 uint32, row0 uint32, col0 uint32) (result0 uint32)

// GetUcharAt3 represents the imported method "get-uchar-at3".
//
// GetUCharAt3 returns the value at the specified x, y, z as a u8.
//
//	get-uchar-at3: func(x: u32, y: u32, z: u32) -> u8
//
//go:nosplit
func (self Mat) GetUcharAt3(x uint32, y uint32, z uint32) (result uint8) {
	self0 := cm.Reinterpret[uint32](self)
	x0 := (uint32)(x)
	y0 := (uint32)(y)
	z0 := (uint32)(z)
	result0 := wasmimport_MatGetUcharAt3((uint32)(self0), (uint32)(x0), (uint32)(y0), (uint32)(z0))
	result = (uint8)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-uchar-at3
//go:noescape
func wasmimport_MatGetUcharAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32) (result0 uint32)

// GetVecbAt represents the imported method "get-vecb-at".
//
// GetVecbAt returns a vector of bytes. Its size corresponds to the number of channels
// of the Mat.
//
//	get-vecb-at: func(row: u32, col: u32) -> list<u8>
//
//go:nosplit
func (self Mat) GetVecbAt(row uint32, col uint32) (result cm.List[uint8]) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	wasmimport_MatGetVecbAt((uint32)(self0), (uint32)(row0), (uint32)(col0), &result)
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-vecb-at
//go:noescape
func wasmimport_MatGetVecbAt(self0 uint32, row0 uint32, col0 uint32, result *cm.List[uint8])

// GetVecfAt represents the imported method "get-vecf-at".
//
// GetVecfAt returns a vector of floats. Its size corresponds to the number of channels
// of the Mat.
//
//	get-vecf-at: func(row: u32, col: u32) -> list<f32>
//
//go:nosplit
func (self Mat) GetVecfAt(row uint32, col uint32) (result cm.List[float32]) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	wasmimport_MatGetVecfAt((uint32)(self0), (uint32)(row0), (uint32)(col0), &result)
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-vecf-at
//go:noescape
func wasmimport_MatGetVecfAt(self0 uint32, row0 uint32, col0 uint32, result *cm.List[float32])

// GetVeciAt represents the imported method "get-veci-at".
//
// GetVeciAt returns a vector of s32. Its size corresponds to the number of channels
// of the Mat.
//
//	get-veci-at: func(row: u32, col: u32) -> list<s32>
//
//go:nosplit
func (self Mat) GetVeciAt(row uint32, col uint32) (result cm.List[int32]) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	wasmimport_MatGetVeciAt((uint32)(self0), (uint32)(row0), (uint32)(col0), &result)
	return
}

//go:wasmimport wasm:cv/mat [method]mat.get-veci-at
//go:noescape
func wasmimport_MatGetVeciAt(self0 uint32, row0 uint32, col0 uint32, result *cm.List[int32])

// Mattype represents the imported method "mattype".
//
// MatType returns the type of the Mat.
//
//	mattype: func() -> mattype
//
//go:nosplit
func (self Mat) Mattype() (result Mattype) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatMattype((uint32)(self0))
	result = (Mattype)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.mattype
//go:noescape
func wasmimport_MatMattype(self0 uint32) (result0 uint32)

// MinMaxLoc represents the imported method "min-max-loc".
//
// MinMaxLoc finds the global minimum and maximum in an array.
//
// For further details, please see:
// https://docs.opencv.org/trunk/d2/de8/group__core__array.html#gab473bf2eb6d14ff97e89b355dac20707
//
//	min-max-loc: func() -> mix-max-loc-result
//
//go:nosplit
func (self Mat) MinMaxLoc() (result types.MixMaxLocResult) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_MatMinMaxLoc((uint32)(self0), &result)
	return
}

//go:wasmimport wasm:cv/mat [method]mat.min-max-loc
//go:noescape
func wasmimport_MatMinMaxLoc(self0 uint32, result *types.MixMaxLocResult)

// Region represents the imported method "region".
//
// Region returns a new Mat that points to a region of this Mat. Changes made to the
// region Mat will affect the original Mat, since they are pointers to the underlying
// OpenCV Mat object.
//
//	region: func(rect: rect) -> mat
//
//go:nosplit
func (self Mat) Region(rect types.Rect) (result Mat) {
	self0 := cm.Reinterpret[uint32](self)
	rect0, rect1, rect2, rect3 := lower_Rect(rect)
	result0 := wasmimport_MatRegion((uint32)(self0), (uint32)(rect0), (uint32)(rect1), (uint32)(rect2), (uint32)(rect3))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.region
//go:noescape
func wasmimport_MatRegion(self0 uint32, rect0 uint32, rect1 uint32, rect2 uint32, rect3 uint32) (result0 uint32)

// Reshape represents the imported method "reshape".
//
// Reshape changes the shape and/or the number of channels of a 2D matrix without
// copying the data.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#a4eb96e3251417fa88b78e2abd6cfd7d8
//
//	reshape: func(channels: u32, rows: u32) -> mat
//
//go:nosplit
func (self Mat) Reshape(channels uint32, rows uint32) (result Mat) {
	self0 := cm.Reinterpret[uint32](self)
	channels0 := (uint32)(channels)
	rows0 := (uint32)(rows)
	result0 := wasmimport_MatReshape((uint32)(self0), (uint32)(channels0), (uint32)(rows0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.reshape
//go:noescape
func wasmimport_MatReshape(self0 uint32, channels0 uint32, rows0 uint32) (result0 uint32)

// RowRange represents the imported method "row-range".
//
// RowRange creates a matrix header for the specified row span.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#aa6542193430356ad631a9beabc624107
//
//	row-range: func(start: u32, end: u32) -> mat
//
//go:nosplit
func (self Mat) RowRange(start uint32, end uint32) (result Mat) {
	self0 := cm.Reinterpret[uint32](self)
	start0 := (uint32)(start)
	end0 := (uint32)(end)
	result0 := wasmimport_MatRowRange((uint32)(self0), (uint32)(start0), (uint32)(end0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.row-range
//go:noescape
func wasmimport_MatRowRange(self0 uint32, start0 uint32, end0 uint32) (result0 uint32)

// Rows represents the imported method "rows".
//
// Rows returns the number of rows for this Mat.
//
//	rows: func() -> u32
//
//go:nosplit
func (self Mat) Rows() (result uint32) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatRows((uint32)(self0))
	result = (uint32)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.rows
//go:noescape
func wasmimport_MatRows(self0 uint32) (result0 uint32)

// SetFloatAt represents the imported method "set-float-at".
//
// SetFloatAt sets the value at the specified row and column as a f32.
//
//	set-float-at: func(row: u32, col: u32, val: f32)
//
//go:nosplit
func (self Mat) SetFloatAt(row uint32, col uint32, val float32) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	val0 := (float32)(val)
	wasmimport_MatSetFloatAt((uint32)(self0), (uint32)(row0), (uint32)(col0), (float32)(val0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.set-float-at
//go:noescape
func wasmimport_MatSetFloatAt(self0 uint32, row0 uint32, col0 uint32, val0 float32)

// SetFloatAt3 represents the imported method "set-float-at3".
//
// SetFloatAt3 sets the value at the specified x, y, z as a f32.
//
//	set-float-at3: func(x: u32, y: u32, z: u32, val: f32)
//
//go:nosplit
func (self Mat) SetFloatAt3(x uint32, y uint32, z uint32, val float32) {
	self0 := cm.Reinterpret[uint32](self)
	x0 := (uint32)(x)
	y0 := (uint32)(y)
	z0 := (uint32)(z)
	val0 := (float32)(val)
	wasmimport_MatSetFloatAt3((uint32)(self0), (uint32)(x0), (uint32)(y0), (uint32)(z0), (float32)(val0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.set-float-at3
//go:noescape
func wasmimport_MatSetFloatAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32, val0 float32)

// SetIntAt represents the imported method "set-int-at".
//
// SetIntAt sets the value at the specified row and column as a s32.
//
//	set-int-at: func(row: u32, col: u32, val: s32)
//
//go:nosplit
func (self Mat) SetIntAt(row uint32, col uint32, val int32) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	val0 := (uint32)(val)
	wasmimport_MatSetIntAt((uint32)(self0), (uint32)(row0), (uint32)(col0), (uint32)(val0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.set-int-at
//go:noescape
func wasmimport_MatSetIntAt(self0 uint32, row0 uint32, col0 uint32, val0 uint32)

// SetIntAt3 represents the imported method "set-int-at3".
//
// SetIntAt3 sets the value at the specified x, y, z as a s32.
//
//	set-int-at3: func(x: u32, y: u32, z: u32, val: s32)
//
//go:nosplit
func (self Mat) SetIntAt3(x uint32, y uint32, z uint32, val int32) {
	self0 := cm.Reinterpret[uint32](self)
	x0 := (uint32)(x)
	y0 := (uint32)(y)
	z0 := (uint32)(z)
	val0 := (uint32)(val)
	wasmimport_MatSetIntAt3((uint32)(self0), (uint32)(x0), (uint32)(y0), (uint32)(z0), (uint32)(val0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.set-int-at3
//go:noescape
func wasmimport_MatSetIntAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32, val0 uint32)

// SetUcharAt represents the imported method "set-uchar-at".
//
// SetUCharAt sets the value at the specified row and column as a u8.
//
//	set-uchar-at: func(row: u32, col: u32, val: u8)
//
//go:nosplit
func (self Mat) SetUcharAt(row uint32, col uint32, val uint8) {
	self0 := cm.Reinterpret[uint32](self)
	row0 := (uint32)(row)
	col0 := (uint32)(col)
	val0 := (uint32)(val)
	wasmimport_MatSetUcharAt((uint32)(self0), (uint32)(row0), (uint32)(col0), (uint32)(val0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.set-uchar-at
//go:noescape
func wasmimport_MatSetUcharAt(self0 uint32, row0 uint32, col0 uint32, val0 uint32)

// SetUcharAt3 represents the imported method "set-uchar-at3".
//
// SetUCharAt3 sets the value at the specified x, y, z as a u8.
//
//	set-uchar-at3: func(x: u32, y: u32, z: u32, val: u8)
//
//go:nosplit
func (self Mat) SetUcharAt3(x uint32, y uint32, z uint32, val uint8) {
	self0 := cm.Reinterpret[uint32](self)
	x0 := (uint32)(x)
	y0 := (uint32)(y)
	z0 := (uint32)(z)
	val0 := (uint32)(val)
	wasmimport_MatSetUcharAt3((uint32)(self0), (uint32)(x0), (uint32)(y0), (uint32)(z0), (uint32)(val0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.set-uchar-at3
//go:noescape
func wasmimport_MatSetUcharAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32, val0 uint32)

// Size represents the imported method "size".
//
// Size returns an array with one element for each dimension containing the size of
// that dimension for the Mat.
//
//	size: func() -> list<u32>
//
//go:nosplit
func (self Mat) Size() (result cm.List[uint32]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_MatSize((uint32)(self0), &result)
	return
}

//go:wasmimport wasm:cv/mat [method]mat.size
//go:noescape
func wasmimport_MatSize(self0 uint32, result *cm.List[uint32])
