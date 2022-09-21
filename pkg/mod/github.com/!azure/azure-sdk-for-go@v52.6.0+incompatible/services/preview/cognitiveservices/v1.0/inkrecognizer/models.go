package inkrecognizer

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0/inkrecognizer"

// AlternatePatternItem ...
type AlternatePatternItem struct {
	// Category - Possible values include: 'LeafInkDrawing', 'LeafInkBullet', 'LeafInkWord', 'LeafUnknown'
	Category Leaf `json:"category,omitempty"`
	// Points - Array of point objects that represent points that are relevant to the type of recognition unit. For example, for leaf node of inkDrawing category that represents a triangle, points would include the x,y coordinates of the vertices of the recognized triangle. The points represent the coordinates of points used to create the perfectly drawn shape that is closest to the original input. They may not exactly match.
	Points *[]PointDetailsPattern `json:"points,omitempty"`
	// RotationAngle - The angular orientation of an object relative to the horizontal axis
	RotationAngle *float64 `json:"rotationAngle,omitempty"`
	// Confidence - A number between 0 and 1 which indicates the confidence level in the result
	Confidence *float64 `json:"confidence,omitempty"`
	// RecognizedString - The recognized string from an inkWord or the name of a recognized shape in an inkDrawing object
	RecognizedString *string `json:"recognizedString,omitempty"`
}

// AnalysisRequest this shows the expected contents of a request
type AnalysisRequest struct {
	// ApplicationType - This describes the domain of the client application. Possible values include: 'Drawing', 'Writing', 'Mixed'
	ApplicationType Application `json:"applicationType,omitempty"`
	// InputDeviceKind - This identifies the kind of device used as the writing instrument. Possible values include: 'Digitizer', 'Pen', 'LightPen', 'TouchScreen', 'TouchPad', 'WhiteBoard', 'ThreedDigitizer', 'StereoPlotter', 'ArticulatedArm', 'Armature'
	InputDeviceKind InputDevice `json:"inputDeviceKind,omitempty"`
	// Unit - This is the physical unit of the ink strokes. It is up to the application developer to decide how to convert the device specific units to physical units before calling the service. The conversion factor can be different based on the type of the device used. Possible values include: 'Mm', 'Cm', 'In'
	Unit Unit `json:"unit,omitempty"`
	// UnitMultiple -  This is a scaling factor to be applied to the point coordinates when interpreting them in the physical units specified.
	UnitMultiple *float64 `json:"unitMultiple,omitempty"`
	// Language - The IETF BCP 47 language code (for ex. en-US, en-GB, hi-IN etc.) of the expected language for the handwritten content in the ink strokes. The response will include results from this language.
	Language                *string                   `json:"language,omitempty"`
	InkPointValueAttributes *[]InkPointValueAttribute `json:"inkPointValueAttributes,omitempty"`
	// Strokes - This is the array of strokes sent for recognition. Best results are produced when the order of strokes added in the array matches the order in which the user created them. Changing the stroke order may produce unexpected results.
	Strokes *[]Stroke `json:"strokes,omitempty"`
}

// AnalysisResponse this shows the expected contents of a response from the service
type AnalysisResponse struct {
	autorest.Response `json:"-"`
	// Unit - This is the physical unit of the ink strokes. It is up to the application developer to decide how to convert the device specific units to physical units before calling the service. The conversion factor can be different based on the type of the device used. Possible values include: 'Mm', 'Cm', 'In'
	Unit Unit `json:"unit,omitempty"`
	// UnitMultiple -  This is a scaling factor to be applied to the point coordinates when interpreting them in the physical units specified.
	UnitMultiple *float64 `json:"unitMultiple,omitempty"`
	// Language - This is the language used for recognizing handwriting from the ink strokes in the request.
	Language         *string                `json:"language,omitempty"`
	RecognitionUnits *[]RecognitionUnitItem `json:"recognitionUnits,omitempty"`
}

// DrawingAttributesPattern the properties to use when rendering ink
type DrawingAttributesPattern struct {
	// Width - The width of the stylus used to draw the stroke
	Width *float64 `json:"width,omitempty"`
	// Color - This shows the components of the color in rgba format
	Color *DrawingAttributesPatternColor `json:"color,omitempty"`
	// Height - The height of the stylus used to draw the stroke
	Height *float64 `json:"height,omitempty"`
	// FitToCurve -  This indicates whether Bezier smoothing is used to render the stroke
	FitToCurve *bool `json:"fitToCurve,omitempty"`
	// RasterOp - Possible values include: 'NoOperation', 'CopyPen', 'MaskPen'
	RasterOp RasterOp `json:"rasterOp,omitempty"`
	// IgnorePressure -  This indicates whether the thickness of a rendered Stroke changes according the amount of pressure applied.
	IgnorePressure *bool `json:"ignorePressure,omitempty"`
	// Tip - This specifies the tip to be used to draw a stroke. Possible values include: 'Ellipse', 'Rectangle'
	Tip Tip `json:"tip,omitempty"`
}

// DrawingAttributesPatternColor this shows the components of the color in rgba format
type DrawingAttributesPatternColor struct {
	// R - The red component of the color
	R *float64 `json:"r,omitempty"`
	// G - The green component of the color
	G *float64 `json:"g,omitempty"`
	// B - The blue component of the color
	B *float64 `json:"b,omitempty"`
	// A - The alpha component of the color
	A *float64 `json:"a,omitempty"`
}

// ErrorModel ...
type ErrorModel struct {
	// Code - This represents the error code
	Code *string `json:"code,omitempty"`
	// Message - This represents the error message
	Message *string `json:"message,omitempty"`
	// Target - This represents the target of the error message
	Target *string `json:"target,omitempty"`
	// Details - This gives details of the reason(s) for the error
	Details *[]ErrorModelDetailsItem `json:"details,omitempty"`
}

// ErrorModelDetailsItem ...
type ErrorModelDetailsItem struct {
	// Code - This represents the error code
	Code *string `json:"code,omitempty"`
	// Message - This represents the error message
	Message *string `json:"message,omitempty"`
	// Target - This represents the target of the error message
	Target *string `json:"target,omitempty"`
}

// InkPoint an object containing the properties of an point in the path of an ink stroke. The main
// properties are the x and y values. Other include tip pressure, x tilt etc. For the coordinate values, it
// is recommended to have a precision of 8 digits after the decimal to obtain most accurate recognition
// results. The origin (0,0) of the canvas is assumed to be at the top left corner of the canvas
type InkPoint struct {
	// X - The x coordinate of the pen location on the writing surface.
	X *float64 `json:"x,omitempty"`
	// Y - The y coordinate of the pen location on the writing surface.
	Y *float64 `json:"y,omitempty"`
	// Z - The z coordinate of the pen location on the writing space. This may not be used for recognition.
	Z *float64 `json:"z,omitempty"`
	// TipPressure - The force exerted against the tablet surface by the transducer, typically a stylus. This may not be used for recognition.
	TipPressure *float64 `json:"tipPressure,omitempty"`
	// BarrelPressure - The force exerted directly by the user on a transducer sensor, such as a pressure-sensitive button on the barrel of a stylus. This may not be used for recognition.
	BarrelPressure *float64 `json:"barrelPressure,omitempty"`
	// Timestamp - The time relative to the absolute time the transducer last became active. This may not be used for recognition.
	Timestamp *float64 `json:"timestamp,omitempty"`
	// XTilt - The plane angle between the Y-Z plane and the plane containing the transducer axis and the Y axis. This may not be used for recognition.
	XTilt *float64 `json:"xTilt,omitempty"`
	// YTilt - The angle between the X-Z and transducer-X planes. A positive Y Tilt is toward the user. This may not be used for recognition.
	YTilt *float64 `json:"yTilt,omitempty"`
	// Width - The width of the tip of the writing instrument. This is used by touch screen devices to report the width of the finger contact on the writing surface. This may not be used for recognition.
	Width *float64 `json:"width,omitempty"`
	// Height - The height of the tip of the writing instrument. This is used by touch screen devices to report the height of the finger contact on the writing surface. This may not be used for recognition.
	Height *float64 `json:"height,omitempty"`
	// TipSwitch - A switch located on the tip of a stylus indicating contact of the stylus with a surface. This may not be used for recognition.
	TipSwitch *bool `json:"tipSwitch,omitempty"`
	// Inverted - A value that indicates that the currently sensed position originates from the end of a stylus opposite the tip switch. This may not be used for recognition.
	Inverted *bool `json:"inverted,omitempty"`
	// BarrelSwitch - A non-tip button located on the barrel of a stylus. Its function is typically mapped to a system secondary button. This may not be used for recognition.
	BarrelSwitch *bool `json:"barrelSwitch,omitempty"`
	// Eraser - The control is used for erasing objects. It is typically located opposite the writing end of a stylus. This may not be used for recognition.
	Eraser *bool `json:"eraser,omitempty"`
	// SecondaryTip - A secondary switch used in conjunction with the tip switch to indicate pressure above a certain threshold applied with the stylus. This may not be used for recognition.
	SecondaryTip *bool `json:"secondaryTip,omitempty"`
}

// InkPointValueAttribute a container for the attributes of a value contained in the ink point object.
type InkPointValueAttribute struct {
	// Name - The name of the point attribute.
	Name *string `json:"name,omitempty"`
	// LogicalMinimum - The minimum value for the attribute
	LogicalMinimum *float64 `json:"logicalMinimum,omitempty"`
	// LogicalMaximum - The maximum value for the attribute
	LogicalMaximum *float64 `json:"logicalMaximum,omitempty"`
}

// PointDetailsPattern this holds all the properties of one point
type PointDetailsPattern struct {
	// X - This represents the x coordinate of the point
	X *float64 `json:"x,omitempty"`
	// Y - This represents the y coordinate of the point
	Y *float64 `json:"y,omitempty"`
}

// RecognitionUnitItem this identifies the recognized entity
type RecognitionUnitItem struct {
	// ID - The identifier of the recognition unit. This id is used to indicate parent/child relationship between different recognition units.
	ID *int32 `json:"id,omitempty"`
	// Category - Possible values include: 'Root', 'WritingRegion', 'Paragraph', 'Line', 'InkBullet', 'InkDrawing', 'InkWord', 'Unknown'
	Category   Category                `json:"category,omitempty"`
	Alternates *[]AlternatePatternItem `json:"alternates,omitempty"`
	// Center - The coordinates (x,y) of the center of the recognition unit.
	Center *PointDetailsPattern `json:"center,omitempty"`
	// Points - Array of point objects that represent points that are relevant to the type of recognition unit. For example, for a leaf node of inkDrawing category that represents a triangle, points would include the x, y coordinates of the vertices of the recognized triangle. The points represent the coordinates used to create the perfectly drawn shape that is closest to the original input. They may not exactly match.
	Points *[]PointDetailsPattern `json:"points,omitempty"`
	// ChildIds - An array of integers representing the identifier of each child of the current recognition unit.
	ChildIds *[]int32 `json:"childIds,omitempty"`
	// Class - Possible values include: 'ClassContainer', 'ClassLeaf'
	Class Class `json:"class,omitempty"`
	// ParentID - The id of the parent node in the tree structure of the recognition results. parent = 0 indicates that there is no dedicated parent node for this unit.
	ParentID *int32 `json:"parentId,omitempty"`
	// BoundingRectangle - The bounding rectangle of the recognition unit represented by the coordinates of the top left corner (topX,topY) along with width and height of the rectangle. Note that this rectangle is not rotated. So for  rotated objects such as slanted handwriting, it will cover the entire object. The unit will be matched to the one specified in the original request (mm by default.)
	BoundingRectangle *RecognitionUnitItemBoundingRectangle `json:"boundingRectangle,omitempty"`
	// RotatedBoundingRectangle - This is the rotated bounding rectangle that covers the entire recognized object along the angle of rotation of the object. Note that this is NOT the same as rotating the boundingRectangle by the rotation angle.
	RotatedBoundingRectangle *[]PointDetailsPattern `json:"rotatedBoundingRectangle,omitempty"`
	// StrokeIds - This is an array of integers representing the list of stroke Identifiers from the input request body that belong to this recognition unit.
	StrokeIds *[]int32 `json:"strokeIds,omitempty"`
	// RecognizedText - The string contains the text that was recognized. It can be an empty string if the recognizer cannot determine the text.
	RecognizedText *string `json:"recognizedText,omitempty"`
	// Confidence - A number between 0 and 1 which indicates the confidence level in the result.
	Confidence *float64 `json:"confidence,omitempty"`
	// RotationAngle - This is the angle at which the unit is rotated in degrees with respect to the positive X axis.
	RotationAngle *float64 `json:"rotationAngle,omitempty"`
	// RecognizedObject - Possible values include: 'ShapeDrawing', 'ShapeSquare', 'ShapeRectangle', 'ShapeCircle', 'ShapeEllipse', 'ShapeTriangle', 'ShapeIsoscelesTriangle', 'ShapeEquilateralTriangle', 'ShapeRightTriangle', 'ShapeQuadrilateral', 'ShapeDiamond', 'ShapeTrapezoid', 'ShapeParallelogram', 'ShapePentagon', 'ShapeHexagon', 'ShapeBlockArrow', 'ShapeHeart', 'ShapeStarSimple', 'ShapeStarCrossed', 'ShapeCloud', 'ShapeLine', 'ShapeCurve', 'ShapePolyLine'
	RecognizedObject Shape `json:"recognizedObject,omitempty"`
}

// RecognitionUnitItemBoundingRectangle the bounding rectangle of the recognition unit represented by the
// coordinates of the top left corner (topX,topY) along with width and height of the rectangle. Note that
// this rectangle is not rotated. So for  rotated objects such as slanted handwriting, it will cover the
// entire object. The unit will be matched to the one specified in the original request (mm by default.)
type RecognitionUnitItemBoundingRectangle struct {
	// TopX - This is the top left x coordinate
	TopX *float64 `json:"topX,omitempty"`
	// TopY - This is the top left y coordinate
	TopY *float64 `json:"topY,omitempty"`
	// Width - This is width of the bounding rectangle
	Width *float64 `json:"width,omitempty"`
	// Height - The is the height of the bounding rectangle
	Height *float64 `json:"height,omitempty"`
}

// Stroke ...
type Stroke struct {
	// ID - This is treated as a unique identifier for each stroke within a request. If the id is repeated within the same request, the service will return an error.
	ID *int32 `json:"id,omitempty"`
	// Language - The IETF BCP 47 language code (for ex. en-US, en-GB, hi-IN etc.) of the expected language for the handwritten content in this stroke. The response will include results from this language.
	Language          *string                   `json:"language,omitempty"`
	Points            *[]InkPoint               `json:"points,omitempty"`
	DrawingAttributes *DrawingAttributesPattern `json:"drawingAttributes,omitempty"`
	// Kind - This is an optional property which influences the decision about what the stroke kind is between inkWriting and inkDrawing. This property should be set ONLY if the type of user content is known ahead of time. Not setting this value implies the kind is not known ahead of time. Kind represents the type of content the stroke is a part of. Possible values include: 'KindInkDrawing', 'KindInkWriting'
	Kind Kind `json:"kind,omitempty"`
}