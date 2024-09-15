// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package net represents the imported interface "wasm:cv/net".
package net

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/wasmvision/wasmcv/components/tinygo/wasm/cv/mat"
)

// NetBackendType represents the enum "wasm:cv/net#net-backend-type".
//
//	enum net-backend-type {
//		net-backend-default,
//		net-backend-halide,
//		net-backend-openvino,
//		net-backend-opencv,
//		net-backend-vkcom,
//		net-backend-cuda
//	}
type NetBackendType uint8

const (
	NetBackendTypeNetBackendDefault NetBackendType = iota
	NetBackendTypeNetBackendHalide
	NetBackendTypeNetBackendOpenvino
	NetBackendTypeNetBackendOpencv
	NetBackendTypeNetBackendVkcom
	NetBackendTypeNetBackendCuda
)

var stringsNetBackendType = [6]string{
	"net-backend-default",
	"net-backend-halide",
	"net-backend-openvino",
	"net-backend-opencv",
	"net-backend-vkcom",
	"net-backend-cuda",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e NetBackendType) String() string {
	return stringsNetBackendType[e]
}

// NetTargetType represents the enum "wasm:cv/net#net-target-type".
//
//	enum net-target-type {
//		net-target-cpu,
//		net-target-fp32,
//		net-target-fp16,
//		net-target-vpu,
//		net-target-vulkan,
//		net-target-fpga,
//		net-target-cuda,
//		net-target-cuda-fp16
//	}
type NetTargetType uint8

const (
	NetTargetTypeNetTargetCPU NetTargetType = iota
	NetTargetTypeNetTargetFp32
	NetTargetTypeNetTargetFp16
	NetTargetTypeNetTargetVpu
	NetTargetTypeNetTargetVulkan
	NetTargetTypeNetTargetFpga
	NetTargetTypeNetTargetCuda
	NetTargetTypeNetTargetCudaFp16
)

var stringsNetTargetType = [8]string{
	"net-target-cpu",
	"net-target-fp32",
	"net-target-fp16",
	"net-target-vpu",
	"net-target-vulkan",
	"net-target-fpga",
	"net-target-cuda",
	"net-target-cuda-fp16",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e NetTargetType) String() string {
	return stringsNetTargetType[e]
}

// DataLayoutType represents the enum "wasm:cv/net#data-layout-type".
//
//	enum data-layout-type {
//		data-layout-unknown,
//		data-layout-nd,
//		data-layout-nchw,
//		data-layout-nhwc,
//		data-layout-ndhwc,
//		data-layout-planar
//	}
type DataLayoutType uint8

const (
	DataLayoutTypeDataLayoutUnknown DataLayoutType = iota
	DataLayoutTypeDataLayoutNd
	DataLayoutTypeDataLayoutNchw
	DataLayoutTypeDataLayoutNhwc
	DataLayoutTypeDataLayoutNdhwc
	DataLayoutTypeDataLayoutPlanar
)

var stringsDataLayoutType = [6]string{
	"data-layout-unknown",
	"data-layout-nd",
	"data-layout-nchw",
	"data-layout-nhwc",
	"data-layout-ndhwc",
	"data-layout-planar",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e DataLayoutType) String() string {
	return stringsDataLayoutType[e]
}

// PaddingModeType represents the enum "wasm:cv/net#padding-mode-type".
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

// Layer represents the imported resource "wasm:cv/net#layer".
//
//	resource layer
type Layer cm.Resource

// ResourceDrop represents the imported resource-drop for resource "layer".
//
// Drops a resource handle.
//
//go:nosplit
func (self Layer) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_LayerResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/net [resource-drop]layer
//go:noescape
func wasmimport_LayerResourceDrop(self0 uint32)

// NewLayer represents the imported constructor for resource "layer".
//
//	constructor()
//
//go:nosplit
func NewLayer() (result Layer) {
	result0 := wasmimport_NewLayer()
	result = cm.Reinterpret[Layer]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/net [constructor]layer
//go:noescape
func wasmimport_NewLayer() (result0 uint32)

// GetName represents the imported method "get-name".
//
// GetName returns the name of the layer.
//
//	get-name: func() -> string
//
//go:nosplit
func (self Layer) GetName() (result string) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_LayerGetName((uint32)(self0), &result)
	return
}

//go:wasmimport wasm:cv/net [method]layer.get-name
//go:noescape
func wasmimport_LayerGetName(self0 uint32, result *string)

// Net represents the imported resource "wasm:cv/net#net".
//
//	resource net
type Net cm.Resource

// ResourceDrop represents the imported resource-drop for resource "net".
//
// Drops a resource handle.
//
//go:nosplit
func (self Net) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/net [resource-drop]net
//go:noescape
func wasmimport_NetResourceDrop(self0 uint32)

// NewNet represents the imported constructor for resource "net".
//
//	constructor()
//
//go:nosplit
func NewNet() (result Net) {
	result0 := wasmimport_NewNet()
	result = cm.Reinterpret[Net]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/net [constructor]net
//go:noescape
func wasmimport_NewNet() (result0 uint32)

// NetReadNetFromOnnx represents the imported static function "read-net-from-onnx".
//
// ReadNetFromONNX reads a network model stored in ONNX framework's format.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga9198ecaac7c32ddf0aa7a1bcbd359567
//
//	read-net-from-onnx: static func(model: string) -> net
//
//go:nosplit
func NetReadNetFromOnnx(model string) (result Net) {
	model0, model1 := cm.LowerString(model)
	result0 := wasmimport_NetReadNetFromOnnx((*uint8)(model0), (uint32)(model1))
	result = cm.Reinterpret[Net]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/net [static]net.read-net-from-onnx
//go:noescape
func wasmimport_NetReadNetFromOnnx(model0 *uint8, model1 uint32) (result0 uint32)

// Close represents the imported method "close".
//
// close the network
//
//	close: func()
//
//go:nosplit
func (self Net) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetClose((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/net [method]net.close
//go:noescape
func wasmimport_NetClose(self0 uint32)

// Empty represents the imported method "empty".
//
// Empty returns true if there are no layers in the network.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#a6a5778787d5b8770deab5eda6968e66c
//
//	empty: func() -> bool
//
//go:nosplit
func (self Net) Empty() (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_NetEmpty((uint32)(self0))
	result = cm.U32ToBool((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/net [method]net.empty
//go:noescape
func wasmimport_NetEmpty(self0 uint32) (result0 uint32)

// Forward represents the imported method "forward".
//
// Forward runs forward pass to compute output of layer with name outputName.
//
// For further details, please see:
// https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a98ed94cb6ef7063d3697259566da310b
//
//	forward: func(output-name: string) -> mat
//
//go:nosplit
func (self Net) Forward(outputName string) (result mat.Mat) {
	self0 := cm.Reinterpret[uint32](self)
	outputName0, outputName1 := cm.LowerString(outputName)
	result0 := wasmimport_NetForward((uint32)(self0), (*uint8)(outputName0), (uint32)(outputName1))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/net [method]net.forward
//go:noescape
func wasmimport_NetForward(self0 uint32, outputName0 *uint8, outputName1 uint32) (result0 uint32)

// GetLayer represents the imported method "get-layer".
//
// GetLayer returns layer with specified id.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#ac944d7f2d3ead5ef9b1b2fa3885f3ff1
//
//	get-layer: func(id: u32) -> layer
//
//go:nosplit
func (self Net) GetLayer(id uint32) (result Layer) {
	self0 := cm.Reinterpret[uint32](self)
	id0 := (uint32)(id)
	result0 := wasmimport_NetGetLayer((uint32)(self0), (uint32)(id0))
	result = cm.Reinterpret[Layer]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/net [method]net.get-layer
//go:noescape
func wasmimport_NetGetLayer(self0 uint32, id0 uint32) (result0 uint32)

// GetLayerNames represents the imported method "get-layer-names".
//
// GetLayerNames returns names of layers in the network.
//
// For further details, please see:
// hhttps://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#a38e67098ae4ae5906bf8d8ea72199c2e
//
//	get-layer-names: func() -> list<string>
//
//go:nosplit
func (self Net) GetLayerNames() (result cm.List[string]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetGetLayerNames((uint32)(self0), &result)
	return
}

//go:wasmimport wasm:cv/net [method]net.get-layer-names
//go:noescape
func wasmimport_NetGetLayerNames(self0 uint32, result *cm.List[string])

// GetUnconnectedOutLayers represents the imported method "get-unconnected-out-layers".
//
// GetUnconnectedOutLayers returns indexes of layers with unconnected outputs.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#ae26f0c29b3733d15d0482098ef9053e3
//
//	get-unconnected-out-layers: func() -> list<u32>
//
//go:nosplit
func (self Net) GetUnconnectedOutLayers() (result cm.List[uint32]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetGetUnconnectedOutLayers((uint32)(self0), &result)
	return
}

//go:wasmimport wasm:cv/net [method]net.get-unconnected-out-layers
//go:noescape
func wasmimport_NetGetUnconnectedOutLayers(self0 uint32, result *cm.List[uint32])

// SetInput represents the imported method "set-input".
//
// SetInput sets the new input value for the network.
//
// For further details, please see:
// https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a672a08ae76444d75d05d7bfea3e4a328
//
//	set-input: func(input: mat, name: string)
//
//go:nosplit
func (self Net) SetInput(input mat.Mat, name string) {
	self0 := cm.Reinterpret[uint32](self)
	input0 := cm.Reinterpret[uint32](input)
	name0, name1 := cm.LowerString(name)
	wasmimport_NetSetInput((uint32)(self0), (uint32)(input0), (*uint8)(name0), (uint32)(name1))
	return
}

//go:wasmimport wasm:cv/net [method]net.set-input
//go:noescape
func wasmimport_NetSetInput(self0 uint32, input0 uint32, name0 *uint8, name1 uint32)
