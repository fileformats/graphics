package elements

import (
	"github.com/fileformats/graphics/jt/model"
)

type ElementType byte

const (
	PartitionNodeType ElementType = iota
	BaseNodeType
	BaseAttributeType
)

type JTElement struct {
	Id model.GUID
	Length int32
	Type model.ObjectBaseType
}

func (e *JTElement) Read(context *model.Context) error {
	if e.Id.Equals(model.EndOfElements) {
		return nil
	}
	length := e.Length - 17
	if length > 0 {
		var data = make([]byte, length)
		context.Data.Unpack(data)
	}
	return context.Data.GetError()
}

type Element interface {
	model.FileObjectReader
	GUID() model.GUID
	BaseElement() *JTElement
}

func New(context *model.Context) Element {
	context.LogGroup("Reading Element")
	defer context.LogGroupEnd()

	var jtElement = &JTElement{}
	jtElement.Length = context.Data.Int32()
	context.Log("Length: %d", jtElement.Length)

	context.Data.Unpack(&jtElement.Id)
	context.Log("GUID: %s (%s)", jtElement.Id, jtElement.Id.Name())

	if !jtElement.Id.Equals(model.EndOfElements) {
		jtElement.Type = model.ObjectBaseType(context.Data.UInt8())
	}
	context.Log("Type: %d (%s)", jtElement.Type, jtElement.Type)

	var element Element
	
	switch jtElement.Id.String() {
	case "ffffffff-ffff-ffff-ff-ff-ff-ff-ff-ff-ff-ff":
		element = &EndOfElements{}
	case "10dd1035-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &BaseNode{}
	case "10dd101b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &GroupNode{}
	case "10dd102a-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &InstanceNode{}
	case "10dd102c-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &LODNode{}
	case "ce357245-38fb-11d1-a5-06-00-60-97-bd-c6-e1":
		element = &MetaDataNode{}
	case "d239e7b6-dd77-4289-a0-7d-b0-ee-79-f7-94-94":
		element = &NullShapeNode{}
	case "ce357244-38fb-11d1-a5-06-00-60-97-bd-c6-e1":
		element = &PartNode{}
	case "10dd103e-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &PartitionNode{}
	case "10dd104c-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &RangeLODNode{}
	case "10dd10f3-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &SwitchNode{}
	case "10dd1059-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &BaseShapeNode{}
	case "98134716-0010-0818-19-98-08-00-09-83-5d-5a":
		element = &PointSetShapeNode{}
	case "10dd1048-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &PolygonSetShapeNode{}
	case "10dd1046-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &PolylineSetShapeNode{}
	case "e40373c1-1ad9-11d3-9d-af-00-a0-c9-c7-dd-c2":
		element = &PrimitiveSetShapeNode{}
	case "10dd1077-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &TriStripSetShapeNode{}
	case "10dd107f-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &VertexShapeNode{}
	case "10dd1001-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &BaseAttribute{}
	case "10dd1014-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &DrawStyleAttribute{}
	case "10dd1083-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &GeometricTransformAttribute{}
	case "10dd1028-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &InfiniteLightAttribute{}
	case "10dd1096-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &LightSetAttribute{}
	case "10dd10c4-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &LinestyleAttribute{}
	case "10dd1030-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &MaterialAttribute{}
	case "10dd1045-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &PointLightAttribute{}
	case "8d57c010-e5cb-11d4-84-0e-00-a0-d2-18-2f-9d":
		element = &PointstyleAttribute{}
	case "10dd1073-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &TextureImageAttribute{}
	case "aa1b831d-6e47-4fee-a8-65-cd-7e-1f-2f-39-dc":
		element = &TextureCoordinateGeneratorAttribute{}
	case "a3cfb921-bdeb-48d7-b3-96-8b-8d-0e-f4-85-a0":
		element = &MappingPlane{}
	case "3e70739d-8cb0-41ef-84-5c-a1-98-d4-00-3b-3f":
		element = &MappingCylinder{}
	case "72475fd1-2823-4219-a0-6c-d9-e6-e3-9a-45-c1":
		element = &MappingSphere{}
	case "92f5b094-6499-4d2d-92-aa-60-d0-5a-44-32-cf":
		element = &MappingTriPlanar{}
	case "10dd104b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &BasePropertyAtom{}
	case "ce357246-38fb-11d1-a5-06-00-60-97-bd-c6-e1":
		element = &DatePropertyAtom{}
	case "10dd102b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &IntegerPropertyAtom{}
	case "10dd1019-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &FloatingPointPropertyAtom{}
	case "e0b05be5-fbbd-11d1-a3-a7-00-aa-00-d1-09-54":
		element = &LateLoadedPropertyAtom{}
	case "10dd1004-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &JTObjectReferencePropertyAtom{}
	case "10dd106e-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &StringPropertyAtom{}
	case "ce357249-38fb-11d1-a5-06-00-60-97-bd-c6-e1":
		element = &PMIManagerMetaData{}
	case "ce357247-38fb-11d1-a5-06-00-60-97-bd-c6-e1":
		element = &PropertyProxyMetaData{}
	case "3e637aed-2a89-41f8-a9-fd-55-37-37-03-96-82":
		element = &NullShapeLOD{}
	case "98134716-0011-0818-19-98-08-00-09-83-5d-5a":
		element = &PointSetShapeLOD{}
	case "10dd10a1-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &PolylineSetShapeLOD{}
	case "e40373c2-1ad9-11d3-9d-af-00-a0-c9-c7-dd-c2":
		element = &PrimitiveSetShape{}
	case "10dd10ab-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &TriStripSetShapeLOD{}
	case "10dd109f-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &PolygonSetLOD{}
	case "10dd10b0-2ac8-11d1-9b-6b-00-80-c7-bb-59-97":
		element = &VertexShapeLOD{}
	default:
		context.Log("Unknown element with id: %s", jtElement.Id.String())
	}
	
	if element != nil {
		be := element.BaseElement()
		be.Id = jtElement.Id
		be.Length = jtElement.Length
		be.Type = jtElement.Type
	}

	return element
}
