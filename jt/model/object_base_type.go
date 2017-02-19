package model

type ObjectBaseType uint8

const (
	BaseGraphNodeObject      ObjectBaseType = 0
	GroupGraphNodeObject     ObjectBaseType = 1
	ShapeGraphNodeObject     ObjectBaseType = 2
	BaseAttributeObject      ObjectBaseType = 3
	ShapeLODObject           ObjectBaseType = 4
	BasePropertyObject       ObjectBaseType = 5
	ReferenceObject          ObjectBaseType = 6
	LateLoadedPropertyObject ObjectBaseType = 8
	JtBaseObject             ObjectBaseType = 9
	NoneObject               ObjectBaseType = 255
)

func (s ObjectBaseType) String() string {
	switch s {
	case 0:
		return "BaseGraphNodeObject"
	case 1:
		return "GroupGraphNodeObject"
	case 2:
		return "ShapeGraphNodeObject"
	case 3:
		return "BaseAttributeObject"
	case 4:
		return "ShapeLODObject"
	case 5:
		return "BasePropertyObject"
	case 6:
		return "ReferenceObject"
	case 8:
		return "LateLoadedPropertyObject"
	case 9:
		return "JtBaseObject"
	case 255:
		return "NoneObject"
	}
	return ""
}
