package model

import (
	"fmt"
)

// List of valid GUID
var (

	EndOfElements                              = NewGUID("ffffffff-ffff-ffff-ff-ff-ff-ff-ff-ff-ff-ff")
	BaseNodeElement                            = NewGUID("10dd1035-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	GroupNodeElement                           = NewGUID("10dd101b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	InstanceNodeElement                        = NewGUID("10dd102a-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	LodNodeElement                             = NewGUID("10dd102c-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	MetaDataNodeElement                        = NewGUID("ce357245-38fb-11d1-a5-06-00-60-97-bd-c6-e1")
	NullShapeNodeElement                       = NewGUID("d239e7b6-dd77-4289-a0-7d-b0-ee-79-f7-94-94")
	PartNodeElement                            = NewGUID("ce357244-38fb-11d1-a5-06-00-60-97-bd-c6-e1")
	PartitionNodeElement                       = NewGUID("10dd103e-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	RangeLodNodeElement                        = NewGUID("10dd104c-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	SwitchNodeElement                          = NewGUID("10dd10f3-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	BaseShapeNodeElement                       = NewGUID("10dd1059-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PointSetShapeNodeElement                   = NewGUID("98134716-0010-0818-19-98-08-00-09-83-5d-5a")
	PolygonSetShapeNodeElement                 = NewGUID("10dd1048-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PolylineSetShapeNodeElement                = NewGUID("10dd1046-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PrimitiveSetShapeNodeElement               = NewGUID("e40373c1-1ad9-11d3-9d-af-00-a0-c9-c7-dd-c2")
	TriStripSetShapeNodeElement                = NewGUID("10dd1077-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	VertexShapeNodeElement                     = NewGUID("10dd107f-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	BaseAttributeData                          = NewGUID("10dd1001-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	DrawStyleAttributeElement                  = NewGUID("10dd1014-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	GeometricTransformAttributeElement         = NewGUID("10dd1083-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	InfiniteLightAttributeElement              = NewGUID("10dd1028-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	LightSetAttributeElement                   = NewGUID("10dd1096-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	LinestyleAttributeElement                  = NewGUID("10dd10c4-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	MaterialAttributeElement                   = NewGUID("10dd1030-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PointLightAttributeElement                 = NewGUID("10dd1045-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PointstyleAttributeElement                 = NewGUID("8d57c010-e5cb-11d4-84-0e-00-a0-d2-18-2f-9d")
	TextureImageAttributeElement               = NewGUID("10dd1073-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	TextureCoordinateGeneratorAttributeElement = NewGUID("aa1b831d-6e47-4fee-a8-65-cd-7e-1f-2f-39-dc")
	MappingPlaneElement                        = NewGUID("a3cfb921-bdeb-48d7-b3-96-8b-8d-0e-f4-85-a0")
	MappingCylinderElement                     = NewGUID("3e70739d-8cb0-41ef-84-5c-a1-98-d4-00-3b-3f")
	MappingSphereElement                       = NewGUID("72475fd1-2823-4219-a0-6c-d9-e6-e3-9a-45-c1")
	MappingTriPlanarElement                    = NewGUID("92f5b094-6499-4d2d-92-aa-60-d0-5a-44-32-cf")
	BasePropertyAtomElement                    = NewGUID("10dd104b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	DatePropertyAtomElement                    = NewGUID("ce357246-38fb-11d1-a5-06-00-60-97-bd-c6-e1")
	IntegerPropertyAtomElement                 = NewGUID("10dd102b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	FloatingPointPropertyAtomElement           = NewGUID("10dd1019-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	LateLoadedPropertyAtomElement              = NewGUID("e0b05be5-fbbd-11d1-a3-a7-00-aa-00-d1-09-54")
	JTObjectReferencePropertyAtomElement       = NewGUID("10dd1004-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	StringPropertyAtomElement                  = NewGUID("10dd106e-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	JTBRepElement                              = NewGUID("873a70c0-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PMIManagerMetaDataElement                  = NewGUID("ce357249-38fb-11d1-a5-06-00-60-97-bd-c6-e1")
	PropertyProxyMetaDataElement               = NewGUID("ce357247-38fb-11d1-a5-06-00-60-97-bd-c6-e1")
	NullShapeLODElement                        = NewGUID("3e637aed-2a89-41f8-a9-fd-55-37-37-03-96-82")
	PointSetShapeLODElement                    = NewGUID("98134716-0011-0818-19-98-08-00-09-83-5d-5a")
	PolylineSetShapeLODElement                 = NewGUID("10dd10a1-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PrimitiveSetShapeElement                   = NewGUID("e40373c2-1ad9-11d3-9d-af-00-a0-c9-c7-dd-c2")
	TriStripSetShapeLODElement                 = NewGUID("10dd10ab-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	PolygonSetLODElement                       = NewGUID("10dd109f-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	VertexShapeLODElement                      = NewGUID("10dd10b0-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	XTBRepElement                              = NewGUID("873a70e0-2ac9-11d1-9b-6b-00-80-c7-bb-59-97")
	WireframeRepElement                        = NewGUID("873a70d0-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
	JTULPElement                               = NewGUID("f338a4af-d7d2-41c5-bc-f2-c5-5a-88-b2-1e-73")
	JTLWPAElement                              = NewGUID("d67f8ea8-f524-4879-92-8c-4c-3a-56-1f-b9-3a")
	WireHarnessSetShapeElement                 = NewGUID("4cc7a523-0728-11d3-9d-8b-00-a0-c9-c7-dd-c2")
	BaseShapeLODElement                        = NewGUID("10dd10a4-2ac8-11d1-9b-6b-00-80-c7-bb-59-97")
)



type GUID struct {
	A uint32
	B, C uint16
	D, E, F, G, H, I, J, K byte
}

func NewGUID(val string) (guid GUID) {
	guid = GUID{}
	if len(val) != 0 {
		fmt.Sscanf(val,"%08x-%04x-%04x-%02x-%02x-%02x-%02x-%02x-%02x-%02x-%02x", &guid.A, &guid.B, &guid.C, &guid.D, &guid.E, &guid.F, &guid.G, &guid.H, &guid.I, &guid.J, &guid.K)
	}
	return
}

func (g GUID) Equals(n GUID) bool {
	return g.String() == n.String()
}

func (g GUID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%02x-%02x-%02x-%02x-%02x-%02x-%02x-%02x", g.A, g.B, g.C, g.D, g.E, g.F, g.G, g.H, g.I, g.J, g.K)
}

func (g GUID) Name() string {
	switch g.String() {
		case "ffffffff-ffff-ffff-ff-ff-ff-ff-ff-ff-ff-ff": return "EndOfElements"
		case "10dd1035-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "BaseNodeElement"
		case "10dd101b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "GroupNodeElement"
		case "10dd102a-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "InstanceNodeElement"
		case "10dd102c-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "LodNodeElement"
		case "ce357245-38fb-11d1-a5-06-00-60-97-bd-c6-e1": return "MetaDataNodeElement"
		case "d239e7b6-dd77-4289-a0-7d-b0-ee-79-f7-94-94": return "NullShapeNodeElement"
		case "ce357244-38fb-11d1-a5-06-00-60-97-bd-c6-e1": return "PartNodeElement"
		case "10dd103e-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "PartitionNodeElement"
		case "10dd104c-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "RangeLodNodeElement"
		case "10dd10f3-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "SwitchNodeElement"
		case "10dd1059-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "BaseShapeNodeElement"
		case "98134716-0010-0818-19-98-08-00-09-83-5d-5a": return "PointSetShapeNodeElement"
		case "10dd1048-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "PolygonSetShapeNodeElement"
		case "10dd1046-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "PolylineSetShapeNodeElement"
		case "e40373c1-1ad9-11d3-9d-af-00-a0-c9-c7-dd-c2": return "PrimitiveSetShapeNodeElement"
		case "10dd1077-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "TriStripSetShapeNodeElement"
		case "10dd107f-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "VertexShapeNodeElement"
		case "10dd1001-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "BaseAttributeData"
		case "10dd1014-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "DrawStyleAttributeElement"
		case "10dd1083-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "GeometricTransformAttributeElement"
		case "10dd1028-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "InfiniteLightAttributeElement"
		case "10dd1096-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "LightSetAttributeElement"
		case "10dd10c4-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "LinestyleAttributeElement"
		case "10dd1030-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "MaterialAttributeElement"
		case "10dd1045-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "PointLightAttributeElement"
		case "8d57c010-e5cb-11d4-84-0e-00-a0-d2-18-2f-9d": return "PointstyleAttributeElement"
		case "10dd1073-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "TextureImageAttributeElement"
		case "aa1b831d-6e47-4fee-a8-65-cd-7e-1f-2f-39-dc": return "TextureCoordinateGeneratorAttributeElement"
		case "a3cfb921-bdeb-48d7-b3-96-8b-8d-0e-f4-85-a0": return "MappingPlaneElement"
		case "3e70739d-8cb0-41ef-84-5c-a1-98-d4-00-3b-3f": return "MappingCylinderElement"
		case "72475fd1-2823-4219-a0-6c-d9-e6-e3-9a-45-c1": return "MappingSphereElement"
		case "92f5b094-6499-4d2d-92-aa-60-d0-5a-44-32-cf": return "MappingTriPlanarElement"
		case "10dd104b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "BasePropertyAtomElement"
		case "ce357246-38fb-11d1-a5-06-00-60-97-bd-c6-e1": return "DatePropertyAtomElement"
		case "10dd102b-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "IntegerPropertyAtomElement"
		case "10dd1019-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "FloatingPointPropertyAtomElement"
		case "e0b05be5-fbbd-11d1-a3-a7-00-aa-00-d1-09-54": return "LateLoadedPropertyAtomElement"
		case "10dd1004-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "JTObjectReferencePropertyAtomElement"
		case "10dd106e-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "StringPropertyAtomElement"
		case "873a70c0-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "JTBRepElement"
		case "ce357249-38fb-11d1-a5-06-00-60-97-bd-c6-e1": return "PMIManagerMetaDataElement"
		case "ce357247-38fb-11d1-a5-06-00-60-97-bd-c6-e1": return "PropertyProxyMetaDataElement"
		case "3e637aed-2a89-41f8-a9-fd-55-37-37-03-96-82": return "NullShapeLODElement"
		case "98134716-0011-0818-19-98-08-00-09-83-5d-5a": return "PointSetShapeLODElement"
		case "10dd10a1-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "PolylineSetShapeLODElement"
		case "e40373c2-1ad9-11d3-9d-af-00-a0-c9-c7-dd-c2": return "PrimitiveSetShapeElement"
		case "10dd10ab-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "TriStripSetShapeLODElement"
		case "10dd109f-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "PolygonSetLODElement"
		case "10dd10b0-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "VertexShapeLODElement"
		case "873a70e0-2ac9-11d1-9b-6b-00-80-c7-bb-59-97": return "XTBRepElement"
		case "873a70d0-2ac8-11d1-9b-6b-00-80-c7-bb-59-97": return "WireframeRepElement"
		case "f338a4af-d7d2-41c5-bc-f2-c5-5a-88-b2-1e-73": return "JTULPElement"
		case "d67f8ea8-f524-4879-92-8c-4c-3a-56-1f-b9-3a": return "JTLWPAElement"
		case "4cc7a523-0728-11d3-9d-8b-00-a0-c9-c7-dd-c2": return "WireHarnessSetShapeElement"
	}
	return "Unknown"
}