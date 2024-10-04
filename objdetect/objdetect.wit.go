// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package objdetect represents the imported interface "wasm:cv/objdetect".
package objdetect

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"wasmcv.org/wasm/cv/mat"
	"wasmcv.org/wasm/cv/types"
)

// CascadeClassifier represents the imported resource "wasm:cv/objdetect#cascade-classifier".
//
// CascadeClassifier is a cascade classifier class for object detection.
//
//	resource cascade-classifier
type CascadeClassifier cm.Resource

// ResourceDrop represents the imported resource-drop for resource "cascade-classifier".
//
// Drops a resource handle.
//
//go:nosplit
func (self CascadeClassifier) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_CascadeClassifierResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [resource-drop]cascade-classifier
//go:noescape
func wasmimport_CascadeClassifierResourceDrop(self0 uint32)

// NewCascadeClassifier represents the imported constructor for resource "cascade-classifier".
//
// NewCascadeClassifier returns a new CascadeClassifier.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516
//
//	constructor()
//
//go:nosplit
func NewCascadeClassifier() (result CascadeClassifier) {
	result0 := wasmimport_NewCascadeClassifier()
	result = cm.Reinterpret[CascadeClassifier]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [constructor]cascade-classifier
//go:noescape
func wasmimport_NewCascadeClassifier() (result0 uint32)

// Close represents the imported method "close".
//
// Close the CascadeClassifier
//
//	close: func()
//
//go:nosplit
func (self CascadeClassifier) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_CascadeClassifierClose((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.close
//go:noescape
func wasmimport_CascadeClassifierClose(self0 uint32)

// DetectMultiScale represents the imported method "detect-multi-scale".
//
// DetectMultiScale detects objects of different sizes in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498
//
//	detect-multi-scale: func(image: mat) -> list<rect>
//
//go:nosplit
func (self CascadeClassifier) DetectMultiScale(image mat.Mat) (result cm.List[types.Rect]) {
	self0 := cm.Reinterpret[uint32](self)
	image0 := cm.Reinterpret[uint32](image)
	wasmimport_CascadeClassifierDetectMultiScale((uint32)(self0), (uint32)(image0), &result)
	return
}

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.detect-multi-scale
//go:noescape
func wasmimport_CascadeClassifierDetectMultiScale(self0 uint32, image0 uint32, result *cm.List[types.Rect])

// DetectMultiScaleWithParams represents the imported method "detect-multi-scale-with-params".
//
// DetectMultiScaleWithParams detects objects of different sizes in the input Mat
// image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498
//
//	detect-multi-scale-with-params: func(image: mat, scale: f64, min-neighbors: u32,
//	%flags: u32, min-size: size, max-size: size) -> list<rect>
//
//go:nosplit
func (self CascadeClassifier) DetectMultiScaleWithParams(image mat.Mat, scale float64, minNeighbors uint32, flags uint32, minSize types.Size, maxSize types.Size) (result cm.List[types.Rect]) {
	self0 := cm.Reinterpret[uint32](self)
	image0 := cm.Reinterpret[uint32](image)
	scale0 := (float64)(scale)
	minNeighbors0 := (uint32)(minNeighbors)
	flags0 := (uint32)(flags)
	minSize0, minSize1 := lower_Size(minSize)
	maxSize0, maxSize1 := lower_Size(maxSize)
	wasmimport_CascadeClassifierDetectMultiScaleWithParams((uint32)(self0), (uint32)(image0), (float64)(scale0), (uint32)(minNeighbors0), (uint32)(flags0), (uint32)(minSize0), (uint32)(minSize1), (uint32)(maxSize0), (uint32)(maxSize1), &result)
	return
}

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.detect-multi-scale-with-params
//go:noescape
func wasmimport_CascadeClassifierDetectMultiScaleWithParams(self0 uint32, image0 uint32, scale0 float64, minNeighbors0 uint32, flags0 uint32, minSize0 uint32, minSize1 uint32, maxSize0 uint32, maxSize1 uint32, result *cm.List[types.Rect])

// Load represents the imported method "load".
//
// Load cascade classifier from a file.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#a1a5884c8cc749422f9eb77c2471958bc
//
//	load: func(file: string) -> bool
//
//go:nosplit
func (self CascadeClassifier) Load(file string) (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	file0, file1 := cm.LowerString(file)
	result0 := wasmimport_CascadeClassifierLoad((uint32)(self0), (*uint8)(file0), (uint32)(file1))
	result = cm.U32ToBool((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.load
//go:noescape
func wasmimport_CascadeClassifierLoad(self0 uint32, file0 *uint8, file1 uint32) (result0 uint32)

// HOGDescriptor represents the imported resource "wasm:cv/objdetect#HOG-descriptor".
//
// HOGDescriptor is a Histogram Of Gradiants (HOG) for object detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a723b95b709cfd3f95cf9e616de988fc8
//
//	resource HOG-descriptor
type HOGDescriptor cm.Resource

// ResourceDrop represents the imported resource-drop for resource "HOG-descriptor".
//
// Drops a resource handle.
//
//go:nosplit
func (self HOGDescriptor) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_HOGDescriptorResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [resource-drop]HOG-descriptor
//go:noescape
func wasmimport_HOGDescriptorResourceDrop(self0 uint32)

// NewHOGDescriptor represents the imported constructor for resource "HOG-descriptor".
//
// NewHOGDescriptor returns a new HOGDescriptor.
//
//	constructor()
//
//go:nosplit
func NewHOGDescriptor() (result HOGDescriptor) {
	result0 := wasmimport_NewHOGDescriptor()
	result = cm.Reinterpret[HOGDescriptor]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [constructor]HOG-descriptor
//go:noescape
func wasmimport_NewHOGDescriptor() (result0 uint32)

// Close represents the imported method "close".
//
// Close the HOGDescriptor
//
//	close: func()
//
//go:nosplit
func (self HOGDescriptor) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_HOGDescriptorClose((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]HOG-descriptor.close
//go:noescape
func wasmimport_HOGDescriptorClose(self0 uint32)

// DetectMultiScale represents the imported method "detect-multi-scale".
//
// DetectMultiScale detects objects of different sizes in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
//
//	detect-multi-scale: func(image: mat) -> list<rect>
//
//go:nosplit
func (self HOGDescriptor) DetectMultiScale(image mat.Mat) (result cm.List[types.Rect]) {
	self0 := cm.Reinterpret[uint32](self)
	image0 := cm.Reinterpret[uint32](image)
	wasmimport_HOGDescriptorDetectMultiScale((uint32)(self0), (uint32)(image0), &result)
	return
}

//go:wasmimport wasm:cv/objdetect [method]HOG-descriptor.detect-multi-scale
//go:noescape
func wasmimport_HOGDescriptorDetectMultiScale(self0 uint32, image0 uint32, result *cm.List[types.Rect])

// DetectMultiScaleWithParams represents the imported method "detect-multi-scale-with-params".
//
// DetectMultiScaleWithParams detects objects of different sizes in the input Mat
// image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
//
//	detect-multi-scale-with-params: func(image: mat, hit-threshold: f64, win-stride:
//	size, padding: size, scale: f64, final-threshold: f64, use-meanshift-grouping:
//	bool) -> list<rect>
//
//go:nosplit
func (self HOGDescriptor) DetectMultiScaleWithParams(image mat.Mat, hitThreshold float64, winStride types.Size, padding types.Size, scale float64, finalThreshold float64, useMeanshiftGrouping bool) (result cm.List[types.Rect]) {
	self0 := cm.Reinterpret[uint32](self)
	image0 := cm.Reinterpret[uint32](image)
	hitThreshold0 := (float64)(hitThreshold)
	winStride0, winStride1 := lower_Size(winStride)
	padding0, padding1 := lower_Size(padding)
	scale0 := (float64)(scale)
	finalThreshold0 := (float64)(finalThreshold)
	useMeanshiftGrouping0 := cm.BoolToU32(useMeanshiftGrouping)
	wasmimport_HOGDescriptorDetectMultiScaleWithParams((uint32)(self0), (uint32)(image0), (float64)(hitThreshold0), (uint32)(winStride0), (uint32)(winStride1), (uint32)(padding0), (uint32)(padding1), (float64)(scale0), (float64)(finalThreshold0), (uint32)(useMeanshiftGrouping0), &result)
	return
}

//go:wasmimport wasm:cv/objdetect [method]HOG-descriptor.detect-multi-scale-with-params
//go:noescape
func wasmimport_HOGDescriptorDetectMultiScaleWithParams(self0 uint32, image0 uint32, hitThreshold0 float64, winStride0 uint32, winStride1 uint32, padding0 uint32, padding1 uint32, scale0 float64, finalThreshold0 float64, useMeanshiftGrouping0 uint32, result *cm.List[types.Rect])

// FaceDetectorYN represents the imported resource "wasm:cv/objdetect#face-detector-YN".
//
//	resource face-detector-YN
type FaceDetectorYN cm.Resource

// ResourceDrop represents the imported resource-drop for resource "face-detector-YN".
//
// Drops a resource handle.
//
//go:nosplit
func (self FaceDetectorYN) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FaceDetectorYNResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [resource-drop]face-detector-YN
//go:noescape
func wasmimport_FaceDetectorYNResourceDrop(self0 uint32)

// NewFaceDetectorYN represents the imported constructor for resource "face-detector-YN".
//
// Creates an instance of face detector YN with given parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516
//
//	constructor(model: string, config: string, input-size: size)
//
//go:nosplit
func NewFaceDetectorYN(model string, config string, inputSize types.Size) (result FaceDetectorYN) {
	model0, model1 := cm.LowerString(model)
	config0, config1 := cm.LowerString(config)
	inputSize0, inputSize1 := lower_Size(inputSize)
	result0 := wasmimport_NewFaceDetectorYN((*uint8)(model0), (uint32)(model1), (*uint8)(config0), (uint32)(config1), (uint32)(inputSize0), (uint32)(inputSize1))
	result = cm.Reinterpret[FaceDetectorYN]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [constructor]face-detector-YN
//go:noescape
func wasmimport_NewFaceDetectorYN(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32, inputSize0 uint32, inputSize1 uint32) (result0 uint32)

// FaceDetectorYNNewWithParams represents the imported static function "new-with-params".
//
// new: static func(model: string, config: string, input-size: size) -> face-detector-YN;
// Creates an instance of face detector YN with given parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516
//
//	new-with-params: static func(model: string, config: string, input-size: size, score-threshold:
//	f32, nms-threshold: f32, top-k: u32, backend-id: u32, target-id: u32) -> face-detector-YN
//
//go:nosplit
func FaceDetectorYNNewWithParams(model string, config string, inputSize types.Size, scoreThreshold float32, nmsThreshold float32, topK uint32, backendID uint32, targetID uint32) (result FaceDetectorYN) {
	model0, model1 := cm.LowerString(model)
	config0, config1 := cm.LowerString(config)
	inputSize0, inputSize1 := lower_Size(inputSize)
	scoreThreshold0 := (float32)(scoreThreshold)
	nmsThreshold0 := (float32)(nmsThreshold)
	topK0 := (uint32)(topK)
	backendId0 := (uint32)(backendID)
	targetId0 := (uint32)(targetID)
	result0 := wasmimport_FaceDetectorYNNewWithParams((*uint8)(model0), (uint32)(model1), (*uint8)(config0), (uint32)(config1), (uint32)(inputSize0), (uint32)(inputSize1), (float32)(scoreThreshold0), (float32)(nmsThreshold0), (uint32)(topK0), (uint32)(backendId0), (uint32)(targetId0))
	result = cm.Reinterpret[FaceDetectorYN]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [static]face-detector-YN.new-with-params
//go:noescape
func wasmimport_FaceDetectorYNNewWithParams(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32, inputSize0 uint32, inputSize1 uint32, scoreThreshold0 float32, nmsThreshold0 float32, topK0 uint32, backendId0 uint32, targetId0 uint32) (result0 uint32)

// Close represents the imported method "close".
//
// Close the face detector
//
//	close: func()
//
//go:nosplit
func (self FaceDetectorYN) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FaceDetectorYNClose((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.close
//go:noescape
func wasmimport_FaceDetectorYNClose(self0 uint32)

// Detect represents the imported method "detect".
//
// Detects faces in the input image.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#ac05bd075ca3e6edc0e328927aae6f45b
//
//	detect: func(input: string) -> mat
//
//go:nosplit
func (self FaceDetectorYN) Detect(input string) (result mat.Mat) {
	self0 := cm.Reinterpret[uint32](self)
	input0, input1 := cm.LowerString(input)
	result0 := wasmimport_FaceDetectorYNDetect((uint32)(self0), (*uint8)(input0), (uint32)(input1))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.detect
//go:noescape
func wasmimport_FaceDetectorYNDetect(self0 uint32, input0 *uint8, input1 uint32) (result0 uint32)

// GetInputSize represents the imported method "get-input-size".
//
//	get-input-size: func() -> size
//
//go:nosplit
func (self FaceDetectorYN) GetInputSize() (result types.Size) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FaceDetectorYNGetInputSize((uint32)(self0), &result)
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-input-size
//go:noescape
func wasmimport_FaceDetectorYNGetInputSize(self0 uint32, result *types.Size)

// GetNmsThreshold represents the imported method "get-nms-threshold".
//
//	get-nms-threshold: func() -> f32
//
//go:nosplit
func (self FaceDetectorYN) GetNmsThreshold() (result float32) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_FaceDetectorYNGetNmsThreshold((uint32)(self0))
	result = (float32)((float32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-nms-threshold
//go:noescape
func wasmimport_FaceDetectorYNGetNmsThreshold(self0 uint32) (result0 float32)

// GetScoreThreshold represents the imported method "get-score-threshold".
//
//	get-score-threshold: func() -> f32
//
//go:nosplit
func (self FaceDetectorYN) GetScoreThreshold() (result float32) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_FaceDetectorYNGetScoreThreshold((uint32)(self0))
	result = (float32)((float32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-score-threshold
//go:noescape
func wasmimport_FaceDetectorYNGetScoreThreshold(self0 uint32) (result0 float32)

// GetTopk represents the imported method "get-topk".
//
//	get-topk: func() -> u32
//
//go:nosplit
func (self FaceDetectorYN) GetTopk() (result uint32) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_FaceDetectorYNGetTopk((uint32)(self0))
	result = (uint32)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-topk
//go:noescape
func wasmimport_FaceDetectorYNGetTopk(self0 uint32) (result0 uint32)

// SetInputSize represents the imported method "set-input-size".
//
//	set-input-size: func(size: size)
//
//go:nosplit
func (self FaceDetectorYN) SetInputSize(size types.Size) {
	self0 := cm.Reinterpret[uint32](self)
	size0, size1 := lower_Size(size)
	wasmimport_FaceDetectorYNSetInputSize((uint32)(self0), (uint32)(size0), (uint32)(size1))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-input-size
//go:noescape
func wasmimport_FaceDetectorYNSetInputSize(self0 uint32, size0 uint32, size1 uint32)

// SetNmsThreshold represents the imported method "set-nms-threshold".
//
//	set-nms-threshold: func(threshold: f32)
//
//go:nosplit
func (self FaceDetectorYN) SetNmsThreshold(threshold float32) {
	self0 := cm.Reinterpret[uint32](self)
	threshold0 := (float32)(threshold)
	wasmimport_FaceDetectorYNSetNmsThreshold((uint32)(self0), (float32)(threshold0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-nms-threshold
//go:noescape
func wasmimport_FaceDetectorYNSetNmsThreshold(self0 uint32, threshold0 float32)

// SetScoreThreshold represents the imported method "set-score-threshold".
//
//	set-score-threshold: func(threshold: f32)
//
//go:nosplit
func (self FaceDetectorYN) SetScoreThreshold(threshold float32) {
	self0 := cm.Reinterpret[uint32](self)
	threshold0 := (float32)(threshold)
	wasmimport_FaceDetectorYNSetScoreThreshold((uint32)(self0), (float32)(threshold0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-score-threshold
//go:noescape
func wasmimport_FaceDetectorYNSetScoreThreshold(self0 uint32, threshold0 float32)

// SetTopk represents the imported method "set-topk".
//
//	set-topk: func(topk: u32)
//
//go:nosplit
func (self FaceDetectorYN) SetTopk(topk uint32) {
	self0 := cm.Reinterpret[uint32](self)
	topk0 := (uint32)(topk)
	wasmimport_FaceDetectorYNSetTopk((uint32)(self0), (uint32)(topk0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-topk
//go:noescape
func wasmimport_FaceDetectorYNSetTopk(self0 uint32, topk0 uint32)

// FaceDistanceType represents the enum "wasm:cv/objdetect#face-distance-type".
//
//	enum face-distance-type {
//		face-distance-type-cosine,
//		face-distance-norm-l2
//	}
type FaceDistanceType uint8

const (
	FaceDistanceTypeFaceDistanceTypeCosine FaceDistanceType = iota
	FaceDistanceTypeFaceDistanceNormL2
)

var stringsFaceDistanceType = [2]string{
	"face-distance-type-cosine",
	"face-distance-norm-l2",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e FaceDistanceType) String() string {
	return stringsFaceDistanceType[e]
}

// FaceRecognizerSF represents the imported resource "wasm:cv/objdetect#face-recognizer-SF".
//
//	resource face-recognizer-SF
type FaceRecognizerSF cm.Resource

// ResourceDrop represents the imported resource-drop for resource "face-recognizer-SF".
//
// Drops a resource handle.
//
//go:nosplit
func (self FaceRecognizerSF) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FaceRecognizerSFResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [resource-drop]face-recognizer-SF
//go:noescape
func wasmimport_FaceRecognizerSFResourceDrop(self0 uint32)

// NewFaceRecognizerSF represents the imported constructor for resource "face-recognizer-SF".
//
// Creates an instance of FaceRecognizerSF with given parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a04df90b0cd7d26d350acd92621a35743
//
//	constructor(model: string, config: string)
//
//go:nosplit
func NewFaceRecognizerSF(model string, config string) (result FaceRecognizerSF) {
	model0, model1 := cm.LowerString(model)
	config0, config1 := cm.LowerString(config)
	result0 := wasmimport_NewFaceRecognizerSF((*uint8)(model0), (uint32)(model1), (*uint8)(config0), (uint32)(config1))
	result = cm.Reinterpret[FaceRecognizerSF]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [constructor]face-recognizer-SF
//go:noescape
func wasmimport_NewFaceRecognizerSF(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32) (result0 uint32)

// FaceRecognizerSFNewWithParams represents the imported static function "new-with-params".
//
// Creates an instance of FaceRecognizerSF with given parameters.
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a04df90b0cd7d26d350acd92621a35743
//
//	new-with-params: static func(model: string, config: string, backend-id: u32, target-id:
//	u32) -> face-recognizer-SF
//
//go:nosplit
func FaceRecognizerSFNewWithParams(model string, config string, backendID uint32, targetID uint32) (result FaceRecognizerSF) {
	model0, model1 := cm.LowerString(model)
	config0, config1 := cm.LowerString(config)
	backendId0 := (uint32)(backendID)
	targetId0 := (uint32)(targetID)
	result0 := wasmimport_FaceRecognizerSFNewWithParams((*uint8)(model0), (uint32)(model1), (*uint8)(config0), (uint32)(config1), (uint32)(backendId0), (uint32)(targetId0))
	result = cm.Reinterpret[FaceRecognizerSF]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [static]face-recognizer-SF.new-with-params
//go:noescape
func wasmimport_FaceRecognizerSFNewWithParams(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32, backendId0 uint32, targetId0 uint32) (result0 uint32)

// AlignCrop represents the imported method "align-crop".
//
// Aligns detected face with the source input image and crops it.
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a84492908abecbc9362b4ddc8d46b8345
//
//	align-crop: func(src: mat, face-box: mat) -> mat
//
//go:nosplit
func (self FaceRecognizerSF) AlignCrop(src mat.Mat, faceBox mat.Mat) (result mat.Mat) {
	self0 := cm.Reinterpret[uint32](self)
	src0 := cm.Reinterpret[uint32](src)
	faceBox0 := cm.Reinterpret[uint32](faceBox)
	result0 := wasmimport_FaceRecognizerSFAlignCrop((uint32)(self0), (uint32)(src0), (uint32)(faceBox0))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.align-crop
//go:noescape
func wasmimport_FaceRecognizerSFAlignCrop(self0 uint32, src0 uint32, faceBox0 uint32) (result0 uint32)

// Close represents the imported method "close".
//
// Close the face FaceRecognizerSF
//
//	close: func()
//
//go:nosplit
func (self FaceRecognizerSF) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FaceRecognizerSFClose((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.close
//go:noescape
func wasmimport_FaceRecognizerSFClose(self0 uint32)

// Feature represents the imported method "feature".
//
// Feature extracts face feature from aligned image.
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#ab1b4a3c12213e89091a490c573dc5aba
//
//	feature: func(aligned: mat) -> mat
//
//go:nosplit
func (self FaceRecognizerSF) Feature(aligned mat.Mat) (result mat.Mat) {
	self0 := cm.Reinterpret[uint32](self)
	aligned0 := cm.Reinterpret[uint32](aligned)
	result0 := wasmimport_FaceRecognizerSFFeature((uint32)(self0), (uint32)(aligned0))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.feature
//go:noescape
func wasmimport_FaceRecognizerSFFeature(self0 uint32, aligned0 uint32) (result0 uint32)

// Match represents the imported method "match".
//
// Match calculates the distance between two face features.
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a2f0362ca1e64320a1f3ba7e1386d0219
//
//	match: func(face1: mat, face2: mat) -> f32
//
//go:nosplit
func (self FaceRecognizerSF) Match(face1 mat.Mat, face2 mat.Mat) (result float32) {
	self0 := cm.Reinterpret[uint32](self)
	face10 := cm.Reinterpret[uint32](face1)
	face20 := cm.Reinterpret[uint32](face2)
	result0 := wasmimport_FaceRecognizerSFMatch((uint32)(self0), (uint32)(face10), (uint32)(face20))
	result = (float32)((float32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.match
//go:noescape
func wasmimport_FaceRecognizerSFMatch(self0 uint32, face10 uint32, face20 uint32) (result0 float32)

// MatchWithParams represents the imported method "match-with-params".
//
// Match calculates the distance between two face features.
//
// For further details, please see:
// https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a2f0362ca1e64320a1f3ba7e1386d0219
//
//	match-with-params: func(face1: mat, face2: mat, distance: face-distance-type) ->
//	f32
//
//go:nosplit
func (self FaceRecognizerSF) MatchWithParams(face1 mat.Mat, face2 mat.Mat, distance FaceDistanceType) (result float32) {
	self0 := cm.Reinterpret[uint32](self)
	face10 := cm.Reinterpret[uint32](face1)
	face20 := cm.Reinterpret[uint32](face2)
	distance0 := (uint32)(distance)
	result0 := wasmimport_FaceRecognizerSFMatchWithParams((uint32)(self0), (uint32)(face10), (uint32)(face20), (uint32)(distance0))
	result = (float32)((float32)(result0))
	return
}

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.match-with-params
//go:noescape
func wasmimport_FaceRecognizerSFMatchWithParams(self0 uint32, face10 uint32, face20 uint32, distance0 uint32) (result0 float32)
