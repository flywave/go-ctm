package ctm

/*
#include "openctm.h"
#include <stdlib.h>
#cgo CFLAGS: -I ./
#cgo linux LDFLAGS:  -L ./lib -L /usr/lib/x86_64-linux-gnu -Wl,--start-group  -lm -pthread -ldl -Wl,--end-group
extern void ctmLoadStream(CTMcontext aContext, void * aUserData);
*/
import "C"
import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"runtime"
	"unsafe"
)

type Context struct {
	ctx C.CTMcontext
}

func NewContext(mode CTMenum) *Context {
	c := &Context{ctx: C.ctmNewContext(C.CTMenum(mode))}
	runtime.SetFinalizer(c, (*Context).free)
	return c
}

func (c *Context) free() {
	C.ctmFreeContext(c.ctx)
}

func (c *Context) Error() error {
	code := c.errorCode()
	if code != CTM_NONE {
		return errors.New(c.errorString())
	}
	return nil
}

func (c *Context) errorCode() CTMenum {
	return CTMenum(C.ctmGetError(c.ctx))
}

func (c *Context) errorString() string {
	return C.GoString(C.ctmErrorString(C.ctmGetError(c.ctx)))
}

func (c *Context) GetInteger(aProperty CTMenum) uint32 {
	return uint32(C.ctmGetInteger(c.ctx, C.CTMenum(aProperty)))
}

func (c *Context) GetFloat(aProperty CTMenum) float32 {
	return float32(C.ctmGetFloat(c.ctx, C.CTMenum(aProperty)))
}

func (c *Context) GetIntegerArray(aProperty CTMenum) *C.uint {
	return C.ctmGetIntegerArray(c.ctx, C.CTMenum(aProperty))
}

func (c *Context) GetFloatArray(aProperty CTMenum) *C.float {
	return C.ctmGetFloatArray(c.ctx, C.CTMenum(aProperty))
}

func (c *Context) GetNamedUVMap(name string) CTMenum {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return CTMenum(C.ctmGetNamedUVMap(c.ctx, cstr))
}

func (c *Context) GetUVMapString(aUVMap, aProperty CTMenum) string {
	return C.GoString(C.ctmGetUVMapString(c.ctx, C.CTMenum(aUVMap), C.CTMenum(aProperty)))
}

func (c *Context) GetUVMapFloat(aUVMap, aProperty CTMenum) float32 {
	return float32(C.ctmGetUVMapFloat(c.ctx, C.CTMenum(aUVMap), C.CTMenum(aProperty)))
}

func (c *Context) GetNamedAttribMap(name string) CTMenum {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return CTMenum(C.ctmGetNamedAttribMap(c.ctx, cstr))
}

func (c *Context) GetAttribMapString(aAttribMap, aProperty CTMenum) string {
	return C.GoString(C.ctmGetAttribMapString(c.ctx, C.CTMenum(aAttribMap), C.CTMenum(aProperty)))
}

func (c *Context) GetAttribMapFloat(aAttribMap, aProperty CTMenum) float32 {
	return float32(C.ctmGetAttribMapFloat(c.ctx, C.CTMenum(aAttribMap), C.CTMenum(aProperty)))
}

func (c *Context) GetString(aProperty CTMenum) string {
	return C.GoString(C.ctmGetString(c.ctx, C.CTMenum(aProperty)))
}

func (c *Context) CompressionMethod(aMethod CTMenum) {
	C.ctmCompressionMethod(c.ctx, C.CTMenum(aMethod))
}

func (c *Context) CompressionLevel(aLevel CTMuint) {
	C.ctmCompressionLevel(c.ctx, C.CTMuint(aLevel))
}

func (c *Context) VertexPrecision(aPrecision CTMfloat) {
	C.ctmVertexPrecision(c.ctx, C.CTMfloat(aPrecision))
}

func (c *Context) VertexPrecisionRel(aRelPrecision CTMfloat) {
	C.ctmVertexPrecisionRel(c.ctx, C.CTMfloat(aRelPrecision))
}

func (c *Context) NormalPrecision(aPrecision CTMfloat) {
	C.ctmNormalPrecision(c.ctx, C.CTMfloat(aPrecision))
}

func (c *Context) UVCoordPrecision(aUVMap CTMenum, aPrecision CTMfloat) {
	C.ctmUVCoordPrecision(c.ctx, C.CTMenum(aUVMap), C.CTMfloat(aPrecision))
}

func (c *Context) AttribPrecision(aAttribMap CTMenum, aPrecision CTMfloat) {
	C.ctmAttribPrecision(c.ctx, C.CTMenum(aAttribMap), C.CTMfloat(aPrecision))
}

func (c *Context) FileComment(txt string) {
	cstr := C.CString(txt)
	defer C.free(unsafe.Pointer(cstr))
	C.ctmFileComment(c.ctx, cstr)
}

func (c *Context) DefineMesh(vertices []float32, indices []uint32, normals []float32) {
	aVertexCount := len(vertices) / 3
	aTriangleCount := len(indices) / 3
	var norms *C.float
	if normals != nil {
		norms = (*C.CTMfloat)(unsafe.Pointer(&normals[0]))
	}
	C.ctmDefineMesh(c.ctx, (*C.CTMfloat)(unsafe.Pointer(&vertices[0])), C.CTMuint(aVertexCount), (*C.CTMuint)(unsafe.Pointer(&indices[0])), C.CTMuint(aTriangleCount), norms)
}

func (c *Context) AddUVMap(UVCoords []float32, aName string, aFileName string) {
	caName := C.CString(aName)
	defer C.free(unsafe.Pointer(caName))
	caFileName := C.CString(aFileName)
	defer C.free(unsafe.Pointer(caFileName))
	C.ctmAddUVMap(c.ctx, (*C.CTMfloat)(unsafe.Pointer(&UVCoords[0])), caName, caFileName)
}

func (c *Context) AddAttribMap(attribValues []float32, aName string) {
	caName := C.CString(aName)
	defer C.free(unsafe.Pointer(caName))
	C.ctmAddAttribMap(c.ctx, (*C.CTMfloat)(unsafe.Pointer(&attribValues[0])), caName)
}

func (c *Context) Load(aName string) {
	caName := C.CString(aName)
	defer C.free(unsafe.Pointer(caName))
	C.ctmLoad(c.ctx, caName)
}

func (c *Context) Save(aName string) {
	caName := C.CString(aName)
	defer C.free(unsafe.Pointer(caName))
	C.ctmSave(c.ctx, caName)
}

//export readerHelper
func readerHelper(aBuf unsafe.Pointer, aCount C.CTMuint, aUserData unsafe.Pointer) C.CTMuint {
	ctx := (*(**streamContext)(aUserData))

	var bufsSlice []byte
	bufsSHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufsSlice)))
	bufsSHeader.Cap = int(aCount)
	bufsSHeader.Len = int(aCount)
	bufsSHeader.Data = uintptr(unsafe.Pointer(aBuf))

	n, err := ctx.reader.Read(bufsSlice)

	if err != nil {
		return C.CTMuint(0)
	}

	return C.CTMuint(n)
}

type streamContext struct {
	reader io.Reader
}

func (c *Context) LoadFromBuffer(buf []byte) {
	ctx := new(streamContext)
	ctx.reader = bytes.NewBuffer(buf)
	inptr := new(uintptr)
	*inptr = uintptr(unsafe.Pointer(ctx))
	C.ctmLoadStream(c.ctx, (unsafe.Pointer)(inptr))
}

func (c *Context) SaveToBuffer() []byte {
	var si C.size_t
	buf := C.ctmSaveToBuffer(c.ctx, &si)

	defer C.ctmFreeBuffer(unsafe.Pointer(buf))

	var bufsSlice []byte
	bufsSHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufsSlice)))
	bufsSHeader.Cap = int(si)
	bufsSHeader.Len = int(si)
	bufsSHeader.Data = uintptr(unsafe.Pointer(buf))

	ret := make([]byte, int(si))
	copy(ret, bufsSlice)
	return ret
}
