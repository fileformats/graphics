package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments/quantize"
	"github.com/fileformats/graphics/jt/codec"
	"github.com/cthackers/go/bitstream"
	"bytes"
	"compress/flate"
)

// Vertex Shape LOD Element represents LODs defined by collections of vertices.
type VertexShapeLOD struct {
	BaseShapeLOD
	// Version Number is the version identifier for this Vertex Shape LOD Data
	VersionNumber uint8
	// V8 JT Files. Binding Attributes is a collection of normal, texture coordinate, and color binding information
	BindingAttributes uint64
	// Normal Binding specifies how (at what granularity) normal vector
	// data is supplied (“bound”) for the Shape Rep in either the Lossless
	// Compressed Raw Vertex Data
	NormalBinding uint8
	// Texture Coord Binding specifies how (at what granularity) texture
    // coordinate data is supplied (“bound”) for the Shape Rep in either
    // the Lossless Compressed Raw Vertex Data or Lossy Quantized Raw
    // Vertex Data collections.
	TextureCoordBinding uint8
	// Color Binding specifies how (at what granularity) color data is
	// supplied (“bound”) for the Shape Rep in either the Lossless
	// Compressed Raw Vertex Data or Lossy Quantized Raw Vertex Data
	// collections.
	ColorBinding uint8
	//
	QuantizationParam quantize.QuantizationParam
	//
	PrimitiveListIndices []int32
	//
	RawVertexData quantize.LossyQuantizedRawVertex
	//
	TopoMeshLOD              TopoMeshLOD
	TopoMeshTopologicallyLOD TopoMeshTopologicallyCompressedLOD
	Uncompressed bool

	Normals []model.Vector3D
	Normal []float32
	Vertex []float32
}

func (n VertexShapeLOD) GUID() model.GUID {
	return model.VertexShapeLODElement
}

func (n *VertexShapeLOD) Read(c *model.Context) error {
	c.LogGroup("VertexShapeLOD")
	defer c.LogGroupEnd()

	if err := n.BaseShapeLOD.Read(c); err != nil {
		return err
	}

	if c.Version.Equal(model.V8) {
		n.BindingAttributes = uint64(c.Data.Int32())
		c.Log("BindingAttributes: %d", n.BindingAttributes)
		n.QuantizationParam.Read(c)
	} else {
		n.BindingAttributes = c.Data.UInt64()
		c.Log("BindingAttributes: %d", n.BindingAttributes)
	}

	return c.Data.GetError()
}

func (n *VertexShapeLOD) ReadCompressedData(c *model.Context) (err error) {
	c.LogGroup("VertexShapeCompressedData")
	n.VersionNumber = uint8(c.Data.Int16())
	c.Log("VersionNumber: %d", n.VersionNumber)
	n.NormalBinding = c.Data.UInt8()
	c.Log("NormalBinding: %d", n.NormalBinding)
	n.TextureCoordBinding = c.Data.UInt8()
	c.Log("TextureCoordBinding: %d", n.TextureCoordBinding)
	n.ColorBinding = c.Data.UInt8()
	c.Log("ColorBinding: %d", n.ColorBinding)

	(&n.QuantizationParam).Read(c)

	if n.PrimitiveListIndices, err = (&codec.Int32CDP{}).ReadVecI32(c); err != nil {
		return err
	} else {
		codec.UnpackResidual(n.PrimitiveListIndices, codec.Stride1)
	}
	c.Log("PrimitiveListIndices: %v", n.PrimitiveListIndices)

	if n.QuantizationParam.BitsPerVertex == 0 {
		n.readLosslessRawVertexData(c)
		n.Uncompressed = true
	} else {
		n.Uncompressed = false
		if err := (&n.RawVertexData).Read(c, n.NormalBinding, n.TextureCoordBinding, n.ColorBinding); err != nil {
			return err
		}
		n.Normals = n.RawVertexData.QuantVertexNorm.Normals
	}
	c.Log("Uncompressed: %v", n.Uncompressed)
	c.Log("RawVertexData: %v", n.RawVertexData.QuantVertexCoord)

	return c.Data.GetError()
}

func (n *VertexShapeLOD) readLosslessRawVertexData(c *model.Context) error {
	c.Log("ReadLosslessRawVertexData")
	uncompressedSize := c.Data.Int32()
	c.Log("UncompressedSize: %d", uncompressedSize)
	compressedSize := c.Data.Int32()
	c.Log("CompressedSize: %d", compressedSize)
	var r = c.Data
	if compressedSize > 0 {
		data := make([]byte, compressedSize)
		fr := flate.NewReader(bytes.NewReader(data))
		defer fr.Close()

		r = bitstream.NewReaderBE(fr)
		r.ByteOrder(c.Data.GetByteOrder())
	}

	numFaces := len(n.PrimitiveListIndices) - 1
	c.Log("NumFaces: %d", numFaces)
	numVertices := int(n.PrimitiveListIndices[numFaces])
	c.Log("NumVertices: %d", numVertices)

	normal := make([]float32, numVertices*3)
	vertex := make([]float32, numVertices*3)

	for i := 0; i < numVertices; i++ {
		j := i * 3
		if n.TextureCoordBinding == 1 {

		}
		if n.ColorBinding == 1 {
			c.Log("Color:", r.Float32())
			c.Log("Color:", r.Float32())
			c.Log("Color:", r.Float32())
		}
		if n.NormalBinding == 1 {
			normal[j] = r.Float32()
			normal[j+1] = r.Float32()
			normal[j+2] = r.Float32()
		}
		vertex[j] = r.Float32()
		vertex[j+1] = r.Float32()
		vertex[j+2] = r.Float32()
	}

	return c.Data.GetError()
}

func (n *VertexShapeLOD) ReadData(c *model.Context, isTriSetShape bool) error {
	if isTriSetShape {
		if err := (&n.TopoMeshTopologicallyLOD).Read(c); err != nil {
			return err
		}
	} else {
		if err := (&n.TopoMeshLOD).Read(c); err != nil {
			return err
		}
	}
	return c.Data.GetError()
}

func (n *VertexShapeLOD) BaseElement() *JTElement {
	return &n.JTElement
}
