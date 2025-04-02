// Code generated by wit-bindgen-go. DO NOT EDIT.

package dnn

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasm:cv".

//go:wasmimport wasm:cv/dnn [resource-drop]layer
//go:noescape
func wasmimport_LayerResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/dnn [constructor]layer
//go:noescape
func wasmimport_NewLayer() (result0 uint32)

//go:wasmimport wasm:cv/dnn [method]layer.close
//go:noescape
func wasmimport_LayerClose(self0 uint32)

//go:wasmimport wasm:cv/dnn [method]layer.get-name
//go:noescape
func wasmimport_LayerGetName(self0 uint32, result *cm.Result[string, string, ErrorResult])

//go:wasmimport wasm:cv/dnn [resource-drop]net
//go:noescape
func wasmimport_NetResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/dnn [static]net.read
//go:noescape
func wasmimport_NetRead(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32, result *cm.Result[string, Net, ErrorResult])

//go:wasmimport wasm:cv/dnn [static]net.read-from-ONNX
//go:noescape
func wasmimport_NetReadFromONNX(model0 *uint8, model1 uint32, result *cm.Result[string, Net, ErrorResult])

//go:wasmimport wasm:cv/dnn [method]net.close
//go:noescape
func wasmimport_NetClose(self0 uint32)

//go:wasmimport wasm:cv/dnn [method]net.empty
//go:noescape
func wasmimport_NetEmpty(self0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/dnn [method]net.forward
//go:noescape
func wasmimport_NetForward(self0 uint32, outputName0 *uint8, outputName1 uint32, result *cm.Result[string, Mat, ErrorResult])

//go:wasmimport wasm:cv/dnn [method]net.forward-layers
//go:noescape
func wasmimport_NetForwardLayers(self0 uint32, outputNames0 *string, outputNames1 uint32, result *cm.Result[cm.List[Mat], cm.List[Mat], ErrorResult])

//go:wasmimport wasm:cv/dnn [method]net.get-layer
//go:noescape
func wasmimport_NetGetLayer(self0 uint32, id0 uint32, result *cm.Result[string, Layer, ErrorResult])

//go:wasmimport wasm:cv/dnn [method]net.get-layer-names
//go:noescape
func wasmimport_NetGetLayerNames(self0 uint32, result *cm.Result[cm.List[string], cm.List[string], ErrorResult])

//go:wasmimport wasm:cv/dnn [method]net.get-unconnected-out-layers
//go:noescape
func wasmimport_NetGetUnconnectedOutLayers(self0 uint32, result *cm.Result[cm.List[uint32], cm.List[uint32], ErrorResult])

//go:wasmimport wasm:cv/dnn [method]net.set-input
//go:noescape
func wasmimport_NetSetInput(self0 uint32, input0 uint32, name0 *uint8, name1 uint32, result *cm.Result[ErrorResult, struct{}, ErrorResult])

//go:wasmimport wasm:cv/dnn blob-from-image
//go:noescape
func wasmimport_BlobFromImage(image0 uint32, scaleFactor0 float32, size0 uint32, size1 uint32, mean0 float32, mean1 float32, mean2 float32, mean3 float32, swapRb0 uint32, crop0 uint32, result *cm.Result[string, Mat, ErrorResult])

//go:wasmimport wasm:cv/dnn blob-from-image-with-params
//go:noescape
func wasmimport_BlobFromImageWithParams(image0 uint32, params0 float32, params1 uint32, params2 uint32, params3 float32, params4 float32, params5 float32, params6 float32, params7 uint32, params8 uint32, params9 uint32, params10 uint32, params11 float32, params12 float32, params13 float32, params14 float32, result *cm.Result[string, Mat, ErrorResult])

//go:wasmimport wasm:cv/dnn blob-rects-to-image-rects
//go:noescape
func wasmimport_BlobRectsToImageRects(params *wasmimport_BlobRectsToImageRects_params, result *cm.Result[cm.List[Rect], cm.List[Rect], ErrorResult])

//go:wasmimport wasm:cv/dnn NMS-boxes
//go:noescape
func wasmimport_NMSBoxes(bboxes0 *Rect, bboxes1 uint32, scores0 *float32, scores1 uint32, scoreThreshold0 float32, nmsThreshold0 float32, result *cm.Result[cm.List[uint32], cm.List[uint32], ErrorResult])
