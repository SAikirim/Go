// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package computervision

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.1/computervision"

type DescriptionExclude = original.DescriptionExclude

const (
	Celebrities DescriptionExclude = original.Celebrities
	Landmarks   DescriptionExclude = original.Landmarks
)

type Details = original.Details

const (
	DetailsCelebrities Details = original.DetailsCelebrities
	DetailsLandmarks   Details = original.DetailsLandmarks
)

type Gender = original.Gender

const (
	Female Gender = original.Female
	Male   Gender = original.Male
)

type OcrDetectionLanguage = original.OcrDetectionLanguage

const (
	De OcrDetectionLanguage = original.De
	En OcrDetectionLanguage = original.En
	Es OcrDetectionLanguage = original.Es
	Fr OcrDetectionLanguage = original.Fr
	It OcrDetectionLanguage = original.It
	Nl OcrDetectionLanguage = original.Nl
	Pt OcrDetectionLanguage = original.Pt
)

type OcrLanguages = original.OcrLanguages

const (
	OcrLanguagesAr     OcrLanguages = original.OcrLanguagesAr
	OcrLanguagesCs     OcrLanguages = original.OcrLanguagesCs
	OcrLanguagesDa     OcrLanguages = original.OcrLanguagesDa
	OcrLanguagesDe     OcrLanguages = original.OcrLanguagesDe
	OcrLanguagesEl     OcrLanguages = original.OcrLanguagesEl
	OcrLanguagesEn     OcrLanguages = original.OcrLanguagesEn
	OcrLanguagesEs     OcrLanguages = original.OcrLanguagesEs
	OcrLanguagesFi     OcrLanguages = original.OcrLanguagesFi
	OcrLanguagesFr     OcrLanguages = original.OcrLanguagesFr
	OcrLanguagesHu     OcrLanguages = original.OcrLanguagesHu
	OcrLanguagesIt     OcrLanguages = original.OcrLanguagesIt
	OcrLanguagesJa     OcrLanguages = original.OcrLanguagesJa
	OcrLanguagesKo     OcrLanguages = original.OcrLanguagesKo
	OcrLanguagesNb     OcrLanguages = original.OcrLanguagesNb
	OcrLanguagesNl     OcrLanguages = original.OcrLanguagesNl
	OcrLanguagesPl     OcrLanguages = original.OcrLanguagesPl
	OcrLanguagesPt     OcrLanguages = original.OcrLanguagesPt
	OcrLanguagesRo     OcrLanguages = original.OcrLanguagesRo
	OcrLanguagesRu     OcrLanguages = original.OcrLanguagesRu
	OcrLanguagesSk     OcrLanguages = original.OcrLanguagesSk
	OcrLanguagesSrCyrl OcrLanguages = original.OcrLanguagesSrCyrl
	OcrLanguagesSrLatn OcrLanguages = original.OcrLanguagesSrLatn
	OcrLanguagesSv     OcrLanguages = original.OcrLanguagesSv
	OcrLanguagesTr     OcrLanguages = original.OcrLanguagesTr
	OcrLanguagesUnk    OcrLanguages = original.OcrLanguagesUnk
	OcrLanguagesZhHans OcrLanguages = original.OcrLanguagesZhHans
	OcrLanguagesZhHant OcrLanguages = original.OcrLanguagesZhHant
)

type OperationStatusCodes = original.OperationStatusCodes

const (
	Failed     OperationStatusCodes = original.Failed
	NotStarted OperationStatusCodes = original.NotStarted
	Running    OperationStatusCodes = original.Running
	Succeeded  OperationStatusCodes = original.Succeeded
)

type TextRecognitionResultDimensionUnit = original.TextRecognitionResultDimensionUnit

const (
	Inch  TextRecognitionResultDimensionUnit = original.Inch
	Pixel TextRecognitionResultDimensionUnit = original.Pixel
)

type VisualFeatureTypes = original.VisualFeatureTypes

const (
	VisualFeatureTypesAdult       VisualFeatureTypes = original.VisualFeatureTypesAdult
	VisualFeatureTypesBrands      VisualFeatureTypes = original.VisualFeatureTypesBrands
	VisualFeatureTypesCategories  VisualFeatureTypes = original.VisualFeatureTypesCategories
	VisualFeatureTypesColor       VisualFeatureTypes = original.VisualFeatureTypesColor
	VisualFeatureTypesDescription VisualFeatureTypes = original.VisualFeatureTypesDescription
	VisualFeatureTypesFaces       VisualFeatureTypes = original.VisualFeatureTypesFaces
	VisualFeatureTypesImageType   VisualFeatureTypes = original.VisualFeatureTypesImageType
	VisualFeatureTypesObjects     VisualFeatureTypes = original.VisualFeatureTypesObjects
	VisualFeatureTypesTags        VisualFeatureTypes = original.VisualFeatureTypesTags
)

type AdultInfo = original.AdultInfo
type AnalyzeResults = original.AnalyzeResults
type AreaOfInterestResult = original.AreaOfInterestResult
type BaseClient = original.BaseClient
type BoundingRect = original.BoundingRect
type Category = original.Category
type CategoryDetail = original.CategoryDetail
type CelebritiesModel = original.CelebritiesModel
type CelebrityResults = original.CelebrityResults
type ColorInfo = original.ColorInfo
type DetectResult = original.DetectResult
type DetectedBrand = original.DetectedBrand
type DetectedObject = original.DetectedObject
type DomainModelResults = original.DomainModelResults
type Error = original.Error
type FaceDescription = original.FaceDescription
type FaceRectangle = original.FaceRectangle
type ImageAnalysis = original.ImageAnalysis
type ImageCaption = original.ImageCaption
type ImageDescription = original.ImageDescription
type ImageDescriptionDetails = original.ImageDescriptionDetails
type ImageMetadata = original.ImageMetadata
type ImageTag = original.ImageTag
type ImageType = original.ImageType
type ImageURL = original.ImageURL
type LandmarkResults = original.LandmarkResults
type LandmarksModel = original.LandmarksModel
type Line = original.Line
type ListModelsResult = original.ListModelsResult
type ModelDescription = original.ModelDescription
type ObjectHierarchy = original.ObjectHierarchy
type OcrLine = original.OcrLine
type OcrRegion = original.OcrRegion
type OcrResult = original.OcrResult
type OcrWord = original.OcrWord
type ReadCloser = original.ReadCloser
type ReadOperationResult = original.ReadOperationResult
type ReadResult = original.ReadResult
type TagResult = original.TagResult
type Word = original.Word

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleDescriptionExcludeValues() []DescriptionExclude {
	return original.PossibleDescriptionExcludeValues()
}
func PossibleDetailsValues() []Details {
	return original.PossibleDetailsValues()
}
func PossibleGenderValues() []Gender {
	return original.PossibleGenderValues()
}
func PossibleOcrDetectionLanguageValues() []OcrDetectionLanguage {
	return original.PossibleOcrDetectionLanguageValues()
}
func PossibleOcrLanguagesValues() []OcrLanguages {
	return original.PossibleOcrLanguagesValues()
}
func PossibleOperationStatusCodesValues() []OperationStatusCodes {
	return original.PossibleOperationStatusCodesValues()
}
func PossibleTextRecognitionResultDimensionUnitValues() []TextRecognitionResultDimensionUnit {
	return original.PossibleTextRecognitionResultDimensionUnitValues()
}
func PossibleVisualFeatureTypesValues() []VisualFeatureTypes {
	return original.PossibleVisualFeatureTypesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
