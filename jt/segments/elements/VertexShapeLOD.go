package elements

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"

	"github.com/cthackers/go/bitstream"
	"github.com/fileformats/graphics/jt/codec"
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments/quantize"
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
	Vertex []model.Vector3D
	Faces  []int32

	normals []float32
	vertex []float32
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

	if err := (&n.QuantizationParam).Read(c); err != nil {
		return err
	}

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
	}

	var numVertices int32
	var numFaces int32

	for i := 0; i < len(n.PrimitiveListIndices) - 1; i++ {
		start := n.PrimitiveListIndices[i]
		end := n.PrimitiveListIndices[i + 1]
		numVertices += end - start
		numFaces += end - start - 2
	}

	n.Faces = make([]int32, numFaces*3)

	if n.Uncompressed {
		var k int32
		for i := 0; i < len(n.PrimitiveListIndices) - 1; i++ {
			start := n.PrimitiveListIndices[i]
			end := n.PrimitiveListIndices[i + 1]
			for f := start; f < end-2; f++ {
				if f % 2 == 0 {
					n.Faces[k] = f
					n.Faces[k+1] = f + 1
					n.Faces[k+2] = f + 2
				} else {
					n.Faces[k] = f
					n.Faces[k+2] = f + 1
					n.Faces[k+1] = f + 2
				}
				k+=3
			}
		}
	} else {
		n.vertex = make([]float32, numVertices*3)
		n.normals = make([]float32, numVertices*3)

		var k int32
		for i := 0; i < len(n.PrimitiveListIndices) - 1; i++ {
			start := n.PrimitiveListIndices[i]
			end := n.PrimitiveListIndices[i + 1]

			for v := start; v < end; v++ {
				j := v*3
				idx := n.RawVertexData.VertexDataIndices[v]
				n.vertex[j] = n.RawVertexData.QuantVertexCoord.XVertexCoordCodes[idx]
				n.vertex[j+1] = n.RawVertexData.QuantVertexCoord.YVertexCoordCodes[idx]
				n.vertex[j+2] = n.RawVertexData.QuantVertexCoord.ZVertexCoordCodes[idx]
				n.Vertex = append(n.Vertex, model.Vector3D{n.vertex[j], n.vertex[j+1], n.vertex[j+2]})

				n.normals[j] = n.RawVertexData.QuantVertexNorm.Normals[idx].X
				n.normals[j+1] = n.RawVertexData.QuantVertexNorm.Normals[idx].Y
				n.normals[j+2] = n.RawVertexData.QuantVertexNorm.Normals[idx].Z
				n.Normals = append(n.Normals, model.Vector3D{n.normals[j], n.normals[j+1], n.normals[j+2]})
			}
			for f := start; f < end - 2; f++ {
				if f % 2 == 0 {
					n.Faces[k] = f
					n.Faces[k+1] = f + 1
					n.Faces[k+2] = f + 2
				} else {
					n.Faces[k] = f
					n.Faces[k+2] = f + 1
					n.Faces[k+1] = f + 2
				}
				k+=3
			}
		}
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
		c.Data.Unpack(data)
		fr, _ := zlib.NewReader(bytes.NewReader(data))
		// fr := flate.NewReader(bytes.NewReader(data))
		defer fr.Close()

		data, err := ioutil.ReadAll(fr)
		if err != nil {
			return err
		}
		r = bitstream.NewReaderBE(bytes.NewReader(data))
		r.ByteOrder(c.Data.GetByteOrder())
	}

	numFaces := len(n.PrimitiveListIndices) - 1
	c.Log("NumFaces: %d", numFaces)
	numVertices := int(n.PrimitiveListIndices[numFaces])
	c.Log("NumVertices: %d", numVertices)

	n.normals = make([]float32, numVertices * 3)
	n.vertex = make([]float32, numVertices * 3)

	for i := 0; i < numVertices; i++ {
		j := i * 3
		if n.TextureCoordBinding == 1 {

		}
		if n.ColorBinding == 1 {
			x := r.Float32()
			x = r.Float32()
			x = r.Float32()
			_ = x
		}
		if n.NormalBinding == 1 {
			n.normals[j] = r.Float32()
			n.normals[j+1] = r.Float32()
			n.normals[j+2] = r.Float32()
			n.Normals = append(n.Normals, model.Vector3D{n.normals[j], n.normals[j+1], n.normals[j+2]})
		}

		n.vertex[j] = r.Float32()
		n.vertex[j+1] = r.Float32()
		n.vertex[j+2] = r.Float32()
		n.Vertex = append(n.Vertex, model.Vector3D{n.vertex[j], n.vertex[j+1], n.vertex[j+2]})
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
