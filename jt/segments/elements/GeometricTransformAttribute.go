package elements

import "github.com/fileformats/graphics/jt/model"

// Geometric Transform Attribute Element contains a 4x4 homogeneous transformation matrix that positions the associated
// LSG node’s coordinate system relative to its parent LSG node.
// JT format LSG traversal semantics state that geometric transform attributes accumulate down the LSG through matrix
// multiplication as follows:
//          p’ = pAM
// Where p is a point of the model,  p’ is the transformed point, M is the current modeling transformation matrix
// inherited from ancestor LSG nodes and previous Geometric Transform Attribute Element, and A is the transformation
// matrix of this Geometric Transform Attribute Element.
type GeometricTransformAttribute struct {
	BaseAttribute
	// Version Number is the version identifier for this element
	VersionNumber uint8
	// Stored Values mask is a 16-bit mask where each bit is a flag indicating whether the corresponding element in
	// the matrix is different from the identity matrix. Only elements which are different from the identity matrix are
	// actually stored. The bits are assigned to matrix elements as follows
	// Bit15    Bit14    Bit13    Bit12
	// Bit11    Bit10     Bit9     Bit8
	//  Bit      Bit6     Bit5     Bit4
	//  Bit      Bit2     Bit1     Bit0
	StoredValueMask uint16
	// Element Value specifies a particular matrix element value.
	ElementValue float32
	// Computed transformation matrix
	TransformationMatrix model.Matrix4F32
}

func (n GeometricTransformAttribute) GUID() model.GUID {
	return model.GeometricTransformAttributeElement
}

func (n *GeometricTransformAttribute) Read(c *model.Context) error {
	c.LogGroup("GeometricTransformAttribute")
	defer c.LogGroupEnd()

	if err := (&n.BaseAttribute).Read(c); err != nil {
		return err
	}
	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.GreaterEqThan(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
		}
	}

	n.StoredValueMask = c.Data.UInt16()
	tmp := n.StoredValueMask
	c.Log("StoreValueMask: %d", n.StoredValueMask)

	n.TransformationMatrix = model.Matrix4F32{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
	total := 0
	for i := 0; i < 16; i++ {
		if tmp & 0x8000 != 0 {
			n.TransformationMatrix[i] = c.Data.Float32()
			total++
		}
		tmp = tmp << 1
	}

	// TODO: investigate here if f64 matrix is required

	return c.Data.GetError()
}

func (n *GeometricTransformAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *GeometricTransformAttribute) BaseElement() *JTElement {
	return &n.JTElement
}
