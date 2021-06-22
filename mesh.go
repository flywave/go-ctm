package ctm

import (
	"reflect"
	"unsafe"

	"github.com/flywave/go3d/vec2"
	"github.com/flywave/go3d/vec3"
)

type Mesh struct {
	ctx *Context
}

func LoadMesh(path string) *Mesh {
	m := &Mesh{ctx: NewContext(CTM_IMPORT)}
	m.ctx.Load(path)
	return m
}

func NewMesh(vertices []vec3.T, indices [][3]uint32, normals []vec3.T) *Mesh {
	var verticesSlice []float32
	verticesHeader := (*reflect.SliceHeader)((unsafe.Pointer(&verticesSlice)))
	verticesHeader.Cap = int(len(vertices) * 3)
	verticesHeader.Len = int(len(vertices) * 3)
	verticesHeader.Data = uintptr(unsafe.Pointer(&vertices[0]))

	var indicesSlice []uint32
	indicesHeader := (*reflect.SliceHeader)((unsafe.Pointer(&indicesSlice)))
	indicesHeader.Cap = int(len(normals) * 3)
	indicesHeader.Len = int(len(normals) * 3)
	indicesHeader.Data = uintptr(unsafe.Pointer(&indices[0]))

	var normalsSlice []float32
	normalsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&normalsSlice)))
	normalsHeader.Cap = int(len(normals) * 3)
	normalsHeader.Len = int(len(normals) * 3)
	normalsHeader.Data = uintptr(unsafe.Pointer(&normals[0]))

	m := &Mesh{ctx: NewContext(CTM_IMPORT)}
	m.ctx.DefineMesh(verticesSlice, indicesSlice, normalsSlice)
	return m
}

func (m *Mesh) GetContext() *Context {
	return m.ctx
}

func (m *Mesh) GetVertCount() uint32 {
	return m.ctx.GetInteger(CTM_VERTEX_COUNT)
}

func (m *Mesh) GetFaceCount() uint32 {
	return m.ctx.GetInteger(CTM_TRIANGLE_COUNT)
}

func (m *Mesh) GetUVMapCount() uint32 {
	return m.ctx.GetInteger(CTM_UV_MAP_COUNT)
}

func (m *Mesh) GetAttribMapCount() uint32 {
	return m.ctx.GetInteger(CTM_ATTRIB_MAP_COUNT)
}

func (m *Mesh) HasNormals() bool {
	return m.ctx.GetInteger(CTM_HAS_NORMALS) == uint32(CTM_TRUE)
}

func (m *Mesh) GetCompressionMethod() uint32 {
	return m.ctx.GetInteger(CTM_COMPRESSION_METHOD)
}

func (m *Mesh) SetCompressionMethod(aMethod CTMenum) {
	m.ctx.CompressionMethod(aMethod)
}

func (m *Mesh) SetCompressionLevel(aLevel CTMuint) {
	m.ctx.CompressionLevel(aLevel)
}

func (m *Mesh) GetVertexPrecision() float32 {
	return m.ctx.GetFloat(CTM_VERTEX_PRECISION)
}

func (m *Mesh) SetVertexPrecision(aPrecision float32) {
	m.ctx.VertexPrecision(CTMfloat(aPrecision))
}

func (m *Mesh) GetNormalPrecision() float32 {
	return m.ctx.GetFloat(CTM_NORMAL_PRECISION)
}

func (m *Mesh) SetNormalPrecision(aPrecision float32) {
	m.ctx.NormalPrecision(CTMfloat(aPrecision))
}

func (m *Mesh) GetUVMapPrecision(aUVMap CTMenum) float32 {
	return m.ctx.GetUVMapFloat(aUVMap, CTM_PRECISION)
}

func (m *Mesh) SetUVMapPrecision(aUVMap CTMenum, aPrecision float32) {
	m.ctx.UVCoordPrecision(aUVMap, CTMfloat(aPrecision))
}

func (m *Mesh) GetAttribMapPrecision(aAttribMap CTMenum) float32 {
	return m.ctx.GetAttribMapFloat(aAttribMap, CTM_PRECISION)
}

func (m *Mesh) SetAttribMapPrecision(aAttribMap CTMenum, aPrecision float32) {
	m.ctx.AttribPrecision(aAttribMap, CTMfloat(aPrecision))
}

func (m *Mesh) GetVertices() []vec3.T {
	size := m.GetVertCount()
	data := m.ctx.GetFloatArray(CTM_VERTICES)

	var bufSlice []vec3.T
	bufHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufSlice)))
	bufHeader.Cap = int(size / 3)
	bufHeader.Len = int(size / 3)
	bufHeader.Data = uintptr(unsafe.Pointer(data))

	return bufSlice
}

func (m *Mesh) GetNormals() []vec3.T {
	size := m.GetVertCount()
	data := m.ctx.GetFloatArray(CTM_NORMALS)

	var bufSlice []vec3.T
	bufHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufSlice)))
	bufHeader.Cap = int(size / 3)
	bufHeader.Len = int(size / 3)
	bufHeader.Data = uintptr(unsafe.Pointer(data))

	return bufSlice
}

func (m *Mesh) GetFaces() [][3]uint32 {
	size := m.GetFaceCount()
	data := m.ctx.GetIntegerArray(CTM_INDICES)

	var bufSlice [][3]uint32
	bufHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufSlice)))
	bufHeader.Cap = int(size / 3)
	bufHeader.Len = int(size / 3)
	bufHeader.Data = uintptr(unsafe.Pointer(data))

	return bufSlice
}

func (m *Mesh) AddUVMap(UVCoords []vec2.T, aName string, aFileName string) {
	var uvsSlice []float32
	uvsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&uvsSlice)))
	uvsHeader.Cap = int(len(UVCoords) * 2)
	uvsHeader.Len = int(len(UVCoords) * 2)
	uvsHeader.Data = uintptr(unsafe.Pointer(&UVCoords[0]))

	m.ctx.AddUVMap(uvsSlice, aName, aFileName)
}

func (m *Mesh) GetUVMap(aName string) []vec2.T {
	mod := m.ctx.GetNamedUVMap(aName)
	if mod != CTM_NONE {
		size := m.GetVertCount()
		data := m.ctx.GetFloatArray(mod)

		var bufSlice []vec2.T
		bufHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufSlice)))
		bufHeader.Cap = int(size)
		bufHeader.Len = int(size)
		bufHeader.Data = uintptr(unsafe.Pointer(data))

		return bufSlice
	}
	return nil
}

func (m *Mesh) AddAttribMap(attribValues []float32, aName string) {
	m.ctx.AddAttribMap(attribValues, aName)
}

func (m *Mesh) GetAttribMap(aName string) []float32 {
	mod := m.ctx.GetNamedAttribMap(aName)
	if mod != CTM_NONE {
		size := m.GetVertCount()
		data := m.ctx.GetFloatArray(mod)

		var bufSlice []float32
		bufHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufSlice)))
		bufHeader.Cap = int(size)
		bufHeader.Len = int(size)
		bufHeader.Data = uintptr(unsafe.Pointer(data))

		return bufSlice
	}
	return nil
}

func (m *Mesh) AddColor(colors [][4]float32) {
	var colorsSlice []float32
	colorsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&colorsSlice)))
	colorsHeader.Cap = int(len(colors) * 4)
	colorsHeader.Len = int(len(colors) * 4)
	colorsHeader.Data = uintptr(unsafe.Pointer(&colors[0]))

	m.ctx.AddAttribMap(colorsSlice, "Color")
}

func (m *Mesh) GetColor() [][4]float32 {
	mod := m.ctx.GetNamedAttribMap("Color")
	if mod != CTM_NONE {
		size := m.GetVertCount()
		data := m.ctx.GetFloatArray(mod)

		var bufSlice [][4]float32
		bufHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufSlice)))
		bufHeader.Cap = int(size)
		bufHeader.Len = int(size)
		bufHeader.Data = uintptr(unsafe.Pointer(data))

		return bufSlice
	}
	return nil
}
