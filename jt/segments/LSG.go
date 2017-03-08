package segments

import (
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments/elements"
)

// The LSG segment contains the main structural information of the file.
// The Scene graph contain references to the metadata, attributes (material, transformation)
// and properties, plus all the references to the actual geometric data.
type LSGSegment struct {
	*SegmentData
	// Root node
	PartitionNode *elements.PartitionNode

	Nodes      map[int32]elements.BaseNodeElement
	Attributes map[int32]elements.BaseAttributeElement
	Properties map[int32]elements.BasePropertyAtomElement

	// map elements on their guid
	Elements map[string]*elements.Element
}

func (l *LSGSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("LSGSegment")
	defer context.LogGroupEnd()

	l.Nodes = make(map[int32]elements.BaseNodeElement)
	l.Attributes = make(map[int32]elements.BaseAttributeElement)
	l.Properties = make(map[int32]elements.BasePropertyAtomElement)

	l.readGraphElements(context)
	l.readPropertyElements(context)
	l.generateGraph(context)
	l.extractGeometry(context)

	return context.Data.GetError()
}

func (l *LSGSegment) readGraphElements(context *model.Context) {
	for {
		element := elements.New(context)
		if context.Data.GetError() != nil {
			break
		}

		if element == nil {
			context.Log("Unknown element")
			continue
		}
		if err := element.Read(context); err != nil {
			continue
		}
		if element.GUID().Equals(model.EndOfElements) {
			break
		}
		if element.GUID().Equals(model.PartitionNodeElement) {
			l.PartitionNode = element.(*elements.PartitionNode)
		}
		if n, ok := element.(elements.BaseAttributeElement); ok {
			l.Attributes[n.GetBaseAttribute().ObjectId] = n
		}
		if n, ok := element.(elements.BaseNodeElement); ok {
			l.Nodes[n.GetBaseNode().ObjectId] = n
		}
	}
}

func (l *LSGSegment) readPropertyElements(context *model.Context) {
	for {
		element := elements.New(context)
		if context.Data.GetError() != nil {
			break
		}
		if element == nil {
			context.Log("Unknown element")
			continue
		}
		if err := element.Read(context); err != nil {
			continue
		}
		if element.GUID().Equals(model.EndOfElements) {
			break
		}
		if n, ok := element.(elements.BasePropertyAtomElement); ok {
			l.Properties[n.GetPropertyAtom().ObjectId] = n
		}
	}
}

func (l *LSGSegment) generateGraph(context *model.Context) error {
	context.LogGroup("Generate LSG graph")
	defer context.LogGroupEnd()

	propertiesTable := elements.PropertyTable{}
	if err := propertiesTable.Read(context); err != nil {
		return err
	}

	for objectId, node := range l.Nodes {
		baseNode := node.GetBaseNode()
		element, ok := node.(elements.Element)
		if !ok {
			continue
		}

		// attach base node attributes to object
		for i := 0; i < int(baseNode.AttributeCount); i++ {
			attributeId := baseNode.AttributeObjectId[i]
			if attr, ok := l.Attributes[attributeId]; ok {
				baseNode.Attributes = append(baseNode.Attributes, attr)
			}
		}
		// attach child nodes to node
		if gn, ok := element.(*elements.GroupNode); ok {
			for _, childId := range gn.ChildNodeIds {
				if childNode, ok := l.Nodes[childId]; ok {
					baseNode.Children = append(baseNode.Children, childNode)
				}
			}
		}

		if in, ok := element.(*elements.InstanceNode); ok {
			if childNode, ok := l.Nodes[in.ChildNodeObjectId]; ok {
				baseNode.Children = append(baseNode.Children, childNode)
			}
		}

		if properties, ok := propertiesTable.Properties[objectId]; ok {
			for k, v := range properties.Values {
				if pk, ok := l.Properties[k]; ok {
					if pv, ok := l.Properties[v]; ok {
						if baseNode.Properties == nil {
							baseNode.Properties =  map[elements.BasePropertyAtomElement]elements.BasePropertyAtomElement{}
						}
						baseNode.Properties[pk] = pv
					}
				}
			}
		}

	}
	return context.Data.GetError()
}

func (l *LSGSegment) extractGeometry(context *model.Context) error {
	return context.Data.GetError()
}

func (l *LSGSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}
