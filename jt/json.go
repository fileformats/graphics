package jt

import "github.com/fileformats/graphics/jt/model"

type JSONGeometry struct {
	Vertices []float32 `json:"vertices"`
	Normals  []float32 `json:"normals"`
	Colors   []int32   `json:"colors"`
	Faces    []int32   `json:"faces"`
}

type JSONMeta struct {
	BBox           model.BoundingBox      `json:"bbox"`
	Area           float32                `json:"area"`
	Filename       string                 `json:"filename"`
	FileVersion    string                 `json:"version,omitempty"`
	Properties     map[string]interface{} `json:"properties,omitempty"`
	LevelOfDetails LevelOfDetail          `json:"lod"`
	AvailableLOD   []LevelOfDetail        `json:"availableLOD"`
	Transformation []float32              `json:"transformation"`
	NumberOfJSONs  int                    `json:"numberOfJsons"`
}

type JSONJt struct {
	MetaData JSONMeta     `json:"metadata"`
	Geometry JSONGeometry `json:"geometry"`
}
