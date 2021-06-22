package ctm

/*
#include "openctm.h"
#cgo CFLAGS: -I ./
*/
import "C"

type CTMint C.int
type CTMfloat C.float
type CTMuint C.uint

const (
	CTM_API_VERSION = 0x00000100
	CTM_TRUE        = 1
	CTM_FALSE       = 0
)

type CTMenum uint32

const (
	CTM_NONE                       CTMenum = 0x0000
	CTM_INVALID_CONTEXT            CTMenum = 0x0001
	CTM_INVALID_ARGUMENT           CTMenum = 0x0002
	CTM_INVALID_OPERATION          CTMenum = 0x0003
	CTM_INVALID_MESH               CTMenum = 0x0004
	CTM_OUT_OF_MEMORY              CTMenum = 0x0005
	CTM_FILE_ERROR                 CTMenum = 0x0006
	CTM_BAD_FORMAT                 CTMenum = 0x0007
	CTM_LZMA_ERROR                 CTMenum = 0x0008
	CTM_INTERNAL_ERROR             CTMenum = 0x0009
	CTM_UNSUPPORTED_FORMAT_VERSION CTMenum = 0x000A
	CTM_IMPORT                     CTMenum = 0x0101
	CTM_EXPORT                     CTMenum = 0x0102
	CTM_METHOD_RAW                 CTMenum = 0x0201
	CTM_METHOD_MG1                 CTMenum = 0x0202
	CTM_METHOD_MG2                 CTMenum = 0x0203
	CTM_VERTEX_COUNT               CTMenum = 0x0301
	CTM_TRIANGLE_COUNT             CTMenum = 0x0302
	CTM_HAS_NORMALS                CTMenum = 0x0303
	CTM_UV_MAP_COUNT               CTMenum = 0x0304
	CTM_ATTRIB_MAP_COUNT           CTMenum = 0x0305
	CTM_VERTEX_PRECISION           CTMenum = 0x0306
	CTM_NORMAL_PRECISION           CTMenum = 0x0307
	CTM_COMPRESSION_METHOD         CTMenum = 0x0308
	CTM_FILE_COMMENT               CTMenum = 0x0309
	CTM_NAME                       CTMenum = 0x0501
	CTM_FILE_NAME                  CTMenum = 0x0502
	CTM_PRECISION                  CTMenum = 0x0503
	CTM_INDICES                    CTMenum = 0x0601
	CTM_VERTICES                   CTMenum = 0x0602
	CTM_NORMALS                    CTMenum = 0x0603
	CTM_UV_MAP_1                   CTMenum = 0x0700
	CTM_UV_MAP_2                   CTMenum = 0x0701
	CTM_UV_MAP_3                   CTMenum = 0x0702
	CTM_UV_MAP_4                   CTMenum = 0x0703
	CTM_UV_MAP_5                   CTMenum = 0x0704
	CTM_UV_MAP_6                   CTMenum = 0x0705
	CTM_UV_MAP_7                   CTMenum = 0x0706
	CTM_UV_MAP_8                   CTMenum = 0x0707
	CTM_ATTRIB_MAP_1               CTMenum = 0x0800
	CTM_ATTRIB_MAP_2               CTMenum = 0x0801
	CTM_ATTRIB_MAP_3               CTMenum = 0x0802
	CTM_ATTRIB_MAP_4               CTMenum = 0x0803
	CTM_ATTRIB_MAP_5               CTMenum = 0x0804
	CTM_ATTRIB_MAP_6               CTMenum = 0x0805
	CTM_ATTRIB_MAP_7               CTMenum = 0x0806
	CTM_ATTRIB_MAP_8               CTMenum = 0x0807
)
