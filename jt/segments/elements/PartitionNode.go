package elements

import (
	"github.com/fileformats/graphics/jt/model"
)

// A partition in a JT file must always be either the root or leaf node.
// A leaf partition node represents an external JT file reference and provides a means to partition a model into multiple physical JT files
type PartitionNode struct {
	GroupNode
	// Partition Flags is a collection of flags.  The flags are combined using the binary OR operator.
	// These flags store various state information of the Partition Node Object such as indicating
	// the presence of optional data.
	// All bits fields that are not defined as in use should be set to ―0
	PartitionFlags int32
	// File Name is the relative path portion of the Partition‘s file location. Where ―rleative path
	// should be interpreted to mean the string contains the file name along with any additional path
	// information that locates the partition JT file relative to the location of the referencing JT file
	FileName model.MbString
	// The Transformed BBox is an MCS axis aligned bounding box and represents the transformed geometry
	// extents for all geometry contained in the Partition Node.  This bounding box information may be
	// used by a renderer of JT data to determine whether to load the data contained within the Partition
	// node (i.e. is any part of the bounding box within the view frustum)
	TransformedBBox   model.BoundingBox
	Reserved          model.BoundingBox
	UntransformedBBox model.BoundingBox
	// Area is the total surface area for this node and all of its descendents.
	// This value is stored in MCS coordinate space (i.e. values scaled by MCS scaling)
	Area float32
	// Vertex Count Range is the aggregate minimum and maximum vertex count for all descendants of the Partition Node
	VertexCountRange model.Int32Range
	// Node Count Range is the aggregate minimum and maximum count of all node descendants of the Partition Node.
	NodeCountRange model.Int32Range
	// Polygon Count Range is the aggregate minimum and maximum polygon count for all descendants of the Partition Node.
	PolygonCountRange model.Int32Range
}

func (n *PartitionNode) GUID() model.GUID {
	return model.PartitionNodeElement
}

func (n *PartitionNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *PartitionNode) Read(c *model.Context) error {
	c.LogGroup("PartitionNode")
	defer c.LogGroupEnd()

	pn := &n.GroupNode
	pn.Read(c)

	n.PartitionFlags = c.Data.Int32()
	c.Log("PartitionFlags: %d", n.PartitionFlags)

	(&n.FileName).Read(c)
	c.Log("FileName: %s", n.FileName)

	if n.PartitionFlags&1 == 0 {
		c.Data.Unpack(&n.TransformedBBox)
	} else {
		c.Data.Unpack(&n.Reserved)
	}
	c.Log("TransformationBox: %s", n.TransformedBBox)
	c.Log("Reserved: %s", n.Reserved)

	n.Area = c.Data.Float32()
	c.Log("Area: %f", n.Area)

	c.Data.Unpack(&n.VertexCountRange)
	c.Log("VertexCountRange: %v", n.VertexCountRange)

	c.Data.Unpack(&n.NodeCountRange)
	c.Log("NodeCountRange: %v", n.NodeCountRange)

	c.Data.Unpack(&n.PolygonCountRange)
	c.Log("PolygonCountRange: %v", n.PolygonCountRange)

	if n.PartitionFlags&1 != 0 {
		c.Data.Unpack(&n.UntransformedBBox)
	}
	c.Log("UntransformedBBox: %s", n.UntransformedBBox)

	return c.Data.GetError()
}

func (n *PartitionNode) BaseElement() *JTElement {
	return &n.JTElement
}
