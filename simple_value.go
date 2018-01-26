package coven

//this file is generated by gen.py

import (
	"reflect"
	"unsafe"
)

type convertKind struct {
	srcTyp reflect.Kind
	dstTyp reflect.Kind
}

type cvtOp func(unsafe.Pointer, unsafe.Pointer)

var cvtOps = make(map[convertKind]cvtOp)

func init() {
	cvtOps[convertKind{reflect.Bool, reflect.Bool}] = cvtBoolBool
	cvtOps[convertKind{reflect.Int, reflect.Int}] = cvtIntInt
	cvtOps[convertKind{reflect.Int, reflect.Uint}] = cvtIntUint
	cvtOps[convertKind{reflect.Int, reflect.Int8}] = cvtIntInt8
	cvtOps[convertKind{reflect.Int, reflect.Uint8}] = cvtIntUint8
	cvtOps[convertKind{reflect.Int, reflect.Int16}] = cvtIntInt16
	cvtOps[convertKind{reflect.Int, reflect.Uint16}] = cvtIntUint16
	cvtOps[convertKind{reflect.Int, reflect.Int32}] = cvtIntInt32
	cvtOps[convertKind{reflect.Int, reflect.Uint32}] = cvtIntUint32
	cvtOps[convertKind{reflect.Int, reflect.Int64}] = cvtIntInt64
	cvtOps[convertKind{reflect.Int, reflect.Uint64}] = cvtIntUint64
	cvtOps[convertKind{reflect.Int, reflect.Float32}] = cvtIntFloat32
	cvtOps[convertKind{reflect.Int, reflect.Float64}] = cvtIntFloat64
	cvtOps[convertKind{reflect.Int, reflect.Uintptr}] = cvtIntUintptr
	cvtOps[convertKind{reflect.Int, reflect.String}] = cvtIntString
	cvtOps[convertKind{reflect.Uint, reflect.Int}] = cvtUintInt
	cvtOps[convertKind{reflect.Uint, reflect.Uint}] = cvtUintUint
	cvtOps[convertKind{reflect.Uint, reflect.Int8}] = cvtUintInt8
	cvtOps[convertKind{reflect.Uint, reflect.Uint8}] = cvtUintUint8
	cvtOps[convertKind{reflect.Uint, reflect.Int16}] = cvtUintInt16
	cvtOps[convertKind{reflect.Uint, reflect.Uint16}] = cvtUintUint16
	cvtOps[convertKind{reflect.Uint, reflect.Int32}] = cvtUintInt32
	cvtOps[convertKind{reflect.Uint, reflect.Uint32}] = cvtUintUint32
	cvtOps[convertKind{reflect.Uint, reflect.Int64}] = cvtUintInt64
	cvtOps[convertKind{reflect.Uint, reflect.Uint64}] = cvtUintUint64
	cvtOps[convertKind{reflect.Uint, reflect.Float32}] = cvtUintFloat32
	cvtOps[convertKind{reflect.Uint, reflect.Float64}] = cvtUintFloat64
	cvtOps[convertKind{reflect.Uint, reflect.Uintptr}] = cvtUintUintptr
	cvtOps[convertKind{reflect.Uint, reflect.String}] = cvtUintString
	cvtOps[convertKind{reflect.Int8, reflect.Int}] = cvtInt8Int
	cvtOps[convertKind{reflect.Int8, reflect.Uint}] = cvtInt8Uint
	cvtOps[convertKind{reflect.Int8, reflect.Int8}] = cvtInt8Int8
	cvtOps[convertKind{reflect.Int8, reflect.Uint8}] = cvtInt8Uint8
	cvtOps[convertKind{reflect.Int8, reflect.Int16}] = cvtInt8Int16
	cvtOps[convertKind{reflect.Int8, reflect.Uint16}] = cvtInt8Uint16
	cvtOps[convertKind{reflect.Int8, reflect.Int32}] = cvtInt8Int32
	cvtOps[convertKind{reflect.Int8, reflect.Uint32}] = cvtInt8Uint32
	cvtOps[convertKind{reflect.Int8, reflect.Int64}] = cvtInt8Int64
	cvtOps[convertKind{reflect.Int8, reflect.Uint64}] = cvtInt8Uint64
	cvtOps[convertKind{reflect.Int8, reflect.Float32}] = cvtInt8Float32
	cvtOps[convertKind{reflect.Int8, reflect.Float64}] = cvtInt8Float64
	cvtOps[convertKind{reflect.Int8, reflect.Uintptr}] = cvtInt8Uintptr
	cvtOps[convertKind{reflect.Int8, reflect.String}] = cvtInt8String
	cvtOps[convertKind{reflect.Uint8, reflect.Int}] = cvtUint8Int
	cvtOps[convertKind{reflect.Uint8, reflect.Uint}] = cvtUint8Uint
	cvtOps[convertKind{reflect.Uint8, reflect.Int8}] = cvtUint8Int8
	cvtOps[convertKind{reflect.Uint8, reflect.Uint8}] = cvtUint8Uint8
	cvtOps[convertKind{reflect.Uint8, reflect.Int16}] = cvtUint8Int16
	cvtOps[convertKind{reflect.Uint8, reflect.Uint16}] = cvtUint8Uint16
	cvtOps[convertKind{reflect.Uint8, reflect.Int32}] = cvtUint8Int32
	cvtOps[convertKind{reflect.Uint8, reflect.Uint32}] = cvtUint8Uint32
	cvtOps[convertKind{reflect.Uint8, reflect.Int64}] = cvtUint8Int64
	cvtOps[convertKind{reflect.Uint8, reflect.Uint64}] = cvtUint8Uint64
	cvtOps[convertKind{reflect.Uint8, reflect.Float32}] = cvtUint8Float32
	cvtOps[convertKind{reflect.Uint8, reflect.Float64}] = cvtUint8Float64
	cvtOps[convertKind{reflect.Uint8, reflect.Uintptr}] = cvtUint8Uintptr
	cvtOps[convertKind{reflect.Uint8, reflect.String}] = cvtUint8String
	cvtOps[convertKind{reflect.Int16, reflect.Int}] = cvtInt16Int
	cvtOps[convertKind{reflect.Int16, reflect.Uint}] = cvtInt16Uint
	cvtOps[convertKind{reflect.Int16, reflect.Int8}] = cvtInt16Int8
	cvtOps[convertKind{reflect.Int16, reflect.Uint8}] = cvtInt16Uint8
	cvtOps[convertKind{reflect.Int16, reflect.Int16}] = cvtInt16Int16
	cvtOps[convertKind{reflect.Int16, reflect.Uint16}] = cvtInt16Uint16
	cvtOps[convertKind{reflect.Int16, reflect.Int32}] = cvtInt16Int32
	cvtOps[convertKind{reflect.Int16, reflect.Uint32}] = cvtInt16Uint32
	cvtOps[convertKind{reflect.Int16, reflect.Int64}] = cvtInt16Int64
	cvtOps[convertKind{reflect.Int16, reflect.Uint64}] = cvtInt16Uint64
	cvtOps[convertKind{reflect.Int16, reflect.Float32}] = cvtInt16Float32
	cvtOps[convertKind{reflect.Int16, reflect.Float64}] = cvtInt16Float64
	cvtOps[convertKind{reflect.Int16, reflect.Uintptr}] = cvtInt16Uintptr
	cvtOps[convertKind{reflect.Int16, reflect.String}] = cvtInt16String
	cvtOps[convertKind{reflect.Uint16, reflect.Int}] = cvtUint16Int
	cvtOps[convertKind{reflect.Uint16, reflect.Uint}] = cvtUint16Uint
	cvtOps[convertKind{reflect.Uint16, reflect.Int8}] = cvtUint16Int8
	cvtOps[convertKind{reflect.Uint16, reflect.Uint8}] = cvtUint16Uint8
	cvtOps[convertKind{reflect.Uint16, reflect.Int16}] = cvtUint16Int16
	cvtOps[convertKind{reflect.Uint16, reflect.Uint16}] = cvtUint16Uint16
	cvtOps[convertKind{reflect.Uint16, reflect.Int32}] = cvtUint16Int32
	cvtOps[convertKind{reflect.Uint16, reflect.Uint32}] = cvtUint16Uint32
	cvtOps[convertKind{reflect.Uint16, reflect.Int64}] = cvtUint16Int64
	cvtOps[convertKind{reflect.Uint16, reflect.Uint64}] = cvtUint16Uint64
	cvtOps[convertKind{reflect.Uint16, reflect.Float32}] = cvtUint16Float32
	cvtOps[convertKind{reflect.Uint16, reflect.Float64}] = cvtUint16Float64
	cvtOps[convertKind{reflect.Uint16, reflect.Uintptr}] = cvtUint16Uintptr
	cvtOps[convertKind{reflect.Uint16, reflect.String}] = cvtUint16String
	cvtOps[convertKind{reflect.Int32, reflect.Int}] = cvtInt32Int
	cvtOps[convertKind{reflect.Int32, reflect.Uint}] = cvtInt32Uint
	cvtOps[convertKind{reflect.Int32, reflect.Int8}] = cvtInt32Int8
	cvtOps[convertKind{reflect.Int32, reflect.Uint8}] = cvtInt32Uint8
	cvtOps[convertKind{reflect.Int32, reflect.Int16}] = cvtInt32Int16
	cvtOps[convertKind{reflect.Int32, reflect.Uint16}] = cvtInt32Uint16
	cvtOps[convertKind{reflect.Int32, reflect.Int32}] = cvtInt32Int32
	cvtOps[convertKind{reflect.Int32, reflect.Uint32}] = cvtInt32Uint32
	cvtOps[convertKind{reflect.Int32, reflect.Int64}] = cvtInt32Int64
	cvtOps[convertKind{reflect.Int32, reflect.Uint64}] = cvtInt32Uint64
	cvtOps[convertKind{reflect.Int32, reflect.Float32}] = cvtInt32Float32
	cvtOps[convertKind{reflect.Int32, reflect.Float64}] = cvtInt32Float64
	cvtOps[convertKind{reflect.Int32, reflect.Uintptr}] = cvtInt32Uintptr
	cvtOps[convertKind{reflect.Int32, reflect.String}] = cvtInt32String
	cvtOps[convertKind{reflect.Uint32, reflect.Int}] = cvtUint32Int
	cvtOps[convertKind{reflect.Uint32, reflect.Uint}] = cvtUint32Uint
	cvtOps[convertKind{reflect.Uint32, reflect.Int8}] = cvtUint32Int8
	cvtOps[convertKind{reflect.Uint32, reflect.Uint8}] = cvtUint32Uint8
	cvtOps[convertKind{reflect.Uint32, reflect.Int16}] = cvtUint32Int16
	cvtOps[convertKind{reflect.Uint32, reflect.Uint16}] = cvtUint32Uint16
	cvtOps[convertKind{reflect.Uint32, reflect.Int32}] = cvtUint32Int32
	cvtOps[convertKind{reflect.Uint32, reflect.Uint32}] = cvtUint32Uint32
	cvtOps[convertKind{reflect.Uint32, reflect.Int64}] = cvtUint32Int64
	cvtOps[convertKind{reflect.Uint32, reflect.Uint64}] = cvtUint32Uint64
	cvtOps[convertKind{reflect.Uint32, reflect.Float32}] = cvtUint32Float32
	cvtOps[convertKind{reflect.Uint32, reflect.Float64}] = cvtUint32Float64
	cvtOps[convertKind{reflect.Uint32, reflect.Uintptr}] = cvtUint32Uintptr
	cvtOps[convertKind{reflect.Uint32, reflect.String}] = cvtUint32String
	cvtOps[convertKind{reflect.Int64, reflect.Int}] = cvtInt64Int
	cvtOps[convertKind{reflect.Int64, reflect.Uint}] = cvtInt64Uint
	cvtOps[convertKind{reflect.Int64, reflect.Int8}] = cvtInt64Int8
	cvtOps[convertKind{reflect.Int64, reflect.Uint8}] = cvtInt64Uint8
	cvtOps[convertKind{reflect.Int64, reflect.Int16}] = cvtInt64Int16
	cvtOps[convertKind{reflect.Int64, reflect.Uint16}] = cvtInt64Uint16
	cvtOps[convertKind{reflect.Int64, reflect.Int32}] = cvtInt64Int32
	cvtOps[convertKind{reflect.Int64, reflect.Uint32}] = cvtInt64Uint32
	cvtOps[convertKind{reflect.Int64, reflect.Int64}] = cvtInt64Int64
	cvtOps[convertKind{reflect.Int64, reflect.Uint64}] = cvtInt64Uint64
	cvtOps[convertKind{reflect.Int64, reflect.Float32}] = cvtInt64Float32
	cvtOps[convertKind{reflect.Int64, reflect.Float64}] = cvtInt64Float64
	cvtOps[convertKind{reflect.Int64, reflect.Uintptr}] = cvtInt64Uintptr
	cvtOps[convertKind{reflect.Int64, reflect.String}] = cvtInt64String
	cvtOps[convertKind{reflect.Uint64, reflect.Int}] = cvtUint64Int
	cvtOps[convertKind{reflect.Uint64, reflect.Uint}] = cvtUint64Uint
	cvtOps[convertKind{reflect.Uint64, reflect.Int8}] = cvtUint64Int8
	cvtOps[convertKind{reflect.Uint64, reflect.Uint8}] = cvtUint64Uint8
	cvtOps[convertKind{reflect.Uint64, reflect.Int16}] = cvtUint64Int16
	cvtOps[convertKind{reflect.Uint64, reflect.Uint16}] = cvtUint64Uint16
	cvtOps[convertKind{reflect.Uint64, reflect.Int32}] = cvtUint64Int32
	cvtOps[convertKind{reflect.Uint64, reflect.Uint32}] = cvtUint64Uint32
	cvtOps[convertKind{reflect.Uint64, reflect.Int64}] = cvtUint64Int64
	cvtOps[convertKind{reflect.Uint64, reflect.Uint64}] = cvtUint64Uint64
	cvtOps[convertKind{reflect.Uint64, reflect.Float32}] = cvtUint64Float32
	cvtOps[convertKind{reflect.Uint64, reflect.Float64}] = cvtUint64Float64
	cvtOps[convertKind{reflect.Uint64, reflect.Uintptr}] = cvtUint64Uintptr
	cvtOps[convertKind{reflect.Uint64, reflect.String}] = cvtUint64String
	cvtOps[convertKind{reflect.Float32, reflect.Int}] = cvtFloat32Int
	cvtOps[convertKind{reflect.Float32, reflect.Uint}] = cvtFloat32Uint
	cvtOps[convertKind{reflect.Float32, reflect.Int8}] = cvtFloat32Int8
	cvtOps[convertKind{reflect.Float32, reflect.Uint8}] = cvtFloat32Uint8
	cvtOps[convertKind{reflect.Float32, reflect.Int16}] = cvtFloat32Int16
	cvtOps[convertKind{reflect.Float32, reflect.Uint16}] = cvtFloat32Uint16
	cvtOps[convertKind{reflect.Float32, reflect.Int32}] = cvtFloat32Int32
	cvtOps[convertKind{reflect.Float32, reflect.Uint32}] = cvtFloat32Uint32
	cvtOps[convertKind{reflect.Float32, reflect.Int64}] = cvtFloat32Int64
	cvtOps[convertKind{reflect.Float32, reflect.Uint64}] = cvtFloat32Uint64
	cvtOps[convertKind{reflect.Float32, reflect.Float32}] = cvtFloat32Float32
	cvtOps[convertKind{reflect.Float32, reflect.Float64}] = cvtFloat32Float64
	cvtOps[convertKind{reflect.Float32, reflect.Uintptr}] = cvtFloat32Uintptr
	cvtOps[convertKind{reflect.Float64, reflect.Int}] = cvtFloat64Int
	cvtOps[convertKind{reflect.Float64, reflect.Uint}] = cvtFloat64Uint
	cvtOps[convertKind{reflect.Float64, reflect.Int8}] = cvtFloat64Int8
	cvtOps[convertKind{reflect.Float64, reflect.Uint8}] = cvtFloat64Uint8
	cvtOps[convertKind{reflect.Float64, reflect.Int16}] = cvtFloat64Int16
	cvtOps[convertKind{reflect.Float64, reflect.Uint16}] = cvtFloat64Uint16
	cvtOps[convertKind{reflect.Float64, reflect.Int32}] = cvtFloat64Int32
	cvtOps[convertKind{reflect.Float64, reflect.Uint32}] = cvtFloat64Uint32
	cvtOps[convertKind{reflect.Float64, reflect.Int64}] = cvtFloat64Int64
	cvtOps[convertKind{reflect.Float64, reflect.Uint64}] = cvtFloat64Uint64
	cvtOps[convertKind{reflect.Float64, reflect.Float32}] = cvtFloat64Float32
	cvtOps[convertKind{reflect.Float64, reflect.Float64}] = cvtFloat64Float64
	cvtOps[convertKind{reflect.Float64, reflect.Uintptr}] = cvtFloat64Uintptr
	cvtOps[convertKind{reflect.Complex64, reflect.Complex64}] = cvtComplex64Complex64
	cvtOps[convertKind{reflect.Complex64, reflect.Complex128}] = cvtComplex64Complex128
	cvtOps[convertKind{reflect.Complex128, reflect.Complex64}] = cvtComplex128Complex64
	cvtOps[convertKind{reflect.Complex128, reflect.Complex128}] = cvtComplex128Complex128
	cvtOps[convertKind{reflect.Uintptr, reflect.Int}] = cvtUintptrInt
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint}] = cvtUintptrUint
	cvtOps[convertKind{reflect.Uintptr, reflect.Int8}] = cvtUintptrInt8
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint8}] = cvtUintptrUint8
	cvtOps[convertKind{reflect.Uintptr, reflect.Int16}] = cvtUintptrInt16
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint16}] = cvtUintptrUint16
	cvtOps[convertKind{reflect.Uintptr, reflect.Int32}] = cvtUintptrInt32
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint32}] = cvtUintptrUint32
	cvtOps[convertKind{reflect.Uintptr, reflect.Int64}] = cvtUintptrInt64
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint64}] = cvtUintptrUint64
	cvtOps[convertKind{reflect.Uintptr, reflect.Float32}] = cvtUintptrFloat32
	cvtOps[convertKind{reflect.Uintptr, reflect.Float64}] = cvtUintptrFloat64
	cvtOps[convertKind{reflect.Uintptr, reflect.Uintptr}] = cvtUintptrUintptr
	cvtOps[convertKind{reflect.Uintptr, reflect.String}] = cvtUintptrString
	cvtOps[convertKind{reflect.String, reflect.String}] = cvtStringString
}

func cvtBoolBool(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = (bool)(*(*bool)(sPtr))
}

func cvtIntInt(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int)(sPtr))
}

func cvtIntUint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int)(sPtr))
}

func cvtIntInt8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int)(sPtr))
}

func cvtIntUint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int)(sPtr))
}

func cvtIntInt16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int)(sPtr))
}

func cvtIntUint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int)(sPtr))
}

func cvtIntInt32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int)(sPtr))
}

func cvtIntUint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int)(sPtr))
}

func cvtIntInt64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int)(sPtr))
}

func cvtIntUint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int)(sPtr))
}

func cvtIntFloat32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int)(sPtr))
}

func cvtIntFloat64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int)(sPtr))
}

func cvtIntUintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int)(sPtr))
}

func cvtIntString(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*int)(sPtr))
}

func cvtUintInt(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint)(sPtr))
}

func cvtUintUint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint)(sPtr))
}

func cvtUintInt8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint)(sPtr))
}

func cvtUintUint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint)(sPtr))
}

func cvtUintInt16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint)(sPtr))
}

func cvtUintUint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint)(sPtr))
}

func cvtUintInt32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint)(sPtr))
}

func cvtUintUint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint)(sPtr))
}

func cvtUintInt64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint)(sPtr))
}

func cvtUintUint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint)(sPtr))
}

func cvtUintFloat32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint)(sPtr))
}

func cvtUintFloat64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint)(sPtr))
}

func cvtUintUintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint)(sPtr))
}

func cvtUintString(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*uint)(sPtr))
}

func cvtInt8Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int8)(sPtr))
}

func cvtInt8Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int8)(sPtr))
}

func cvtInt8Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int8)(sPtr))
}

func cvtInt8Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int8)(sPtr))
}

func cvtInt8Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int8)(sPtr))
}

func cvtInt8Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int8)(sPtr))
}

func cvtInt8Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int8)(sPtr))
}

func cvtInt8Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int8)(sPtr))
}

func cvtInt8Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int8)(sPtr))
}

func cvtInt8Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int8)(sPtr))
}

func cvtInt8Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int8)(sPtr))
}

func cvtInt8Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int8)(sPtr))
}

func cvtInt8Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int8)(sPtr))
}

func cvtInt8String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*int8)(sPtr))
}

func cvtUint8Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint8)(sPtr))
}

func cvtUint8Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint8)(sPtr))
}

func cvtUint8Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint8)(sPtr))
}

func cvtUint8Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint8)(sPtr))
}

func cvtUint8Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint8)(sPtr))
}

func cvtUint8Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint8)(sPtr))
}

func cvtUint8Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint8)(sPtr))
}

func cvtUint8Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint8)(sPtr))
}

func cvtUint8Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint8)(sPtr))
}

func cvtUint8Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint8)(sPtr))
}

func cvtUint8Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint8)(sPtr))
}

func cvtUint8Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint8)(sPtr))
}

func cvtUint8Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint8)(sPtr))
}

func cvtUint8String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*uint8)(sPtr))
}

func cvtInt16Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int16)(sPtr))
}

func cvtInt16Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int16)(sPtr))
}

func cvtInt16Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int16)(sPtr))
}

func cvtInt16Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int16)(sPtr))
}

func cvtInt16Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int16)(sPtr))
}

func cvtInt16Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int16)(sPtr))
}

func cvtInt16Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int16)(sPtr))
}

func cvtInt16Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int16)(sPtr))
}

func cvtInt16Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int16)(sPtr))
}

func cvtInt16Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int16)(sPtr))
}

func cvtInt16Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int16)(sPtr))
}

func cvtInt16Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int16)(sPtr))
}

func cvtInt16Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int16)(sPtr))
}

func cvtInt16String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*int16)(sPtr))
}

func cvtUint16Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint16)(sPtr))
}

func cvtUint16Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint16)(sPtr))
}

func cvtUint16Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint16)(sPtr))
}

func cvtUint16Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint16)(sPtr))
}

func cvtUint16Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint16)(sPtr))
}

func cvtUint16Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint16)(sPtr))
}

func cvtUint16Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint16)(sPtr))
}

func cvtUint16Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint16)(sPtr))
}

func cvtUint16Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint16)(sPtr))
}

func cvtUint16Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint16)(sPtr))
}

func cvtUint16Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint16)(sPtr))
}

func cvtUint16Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint16)(sPtr))
}

func cvtUint16Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint16)(sPtr))
}

func cvtUint16String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*uint16)(sPtr))
}

func cvtInt32Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int32)(sPtr))
}

func cvtInt32Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int32)(sPtr))
}

func cvtInt32Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int32)(sPtr))
}

func cvtInt32Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int32)(sPtr))
}

func cvtInt32Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int32)(sPtr))
}

func cvtInt32Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int32)(sPtr))
}

func cvtInt32Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int32)(sPtr))
}

func cvtInt32Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int32)(sPtr))
}

func cvtInt32Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int32)(sPtr))
}

func cvtInt32Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int32)(sPtr))
}

func cvtInt32Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int32)(sPtr))
}

func cvtInt32Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int32)(sPtr))
}

func cvtInt32Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int32)(sPtr))
}

func cvtInt32String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*int32)(sPtr))
}

func cvtUint32Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint32)(sPtr))
}

func cvtUint32Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint32)(sPtr))
}

func cvtUint32Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint32)(sPtr))
}

func cvtUint32Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint32)(sPtr))
}

func cvtUint32Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint32)(sPtr))
}

func cvtUint32Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint32)(sPtr))
}

func cvtUint32Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint32)(sPtr))
}

func cvtUint32Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint32)(sPtr))
}

func cvtUint32Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint32)(sPtr))
}

func cvtUint32Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint32)(sPtr))
}

func cvtUint32Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint32)(sPtr))
}

func cvtUint32Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint32)(sPtr))
}

func cvtUint32Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint32)(sPtr))
}

func cvtUint32String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*uint32)(sPtr))
}

func cvtInt64Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int64)(sPtr))
}

func cvtInt64Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int64)(sPtr))
}

func cvtInt64Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int64)(sPtr))
}

func cvtInt64Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int64)(sPtr))
}

func cvtInt64Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int64)(sPtr))
}

func cvtInt64Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int64)(sPtr))
}

func cvtInt64Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int64)(sPtr))
}

func cvtInt64Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int64)(sPtr))
}

func cvtInt64Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int64)(sPtr))
}

func cvtInt64Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int64)(sPtr))
}

func cvtInt64Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int64)(sPtr))
}

func cvtInt64Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int64)(sPtr))
}

func cvtInt64Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int64)(sPtr))
}

func cvtInt64String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*int64)(sPtr))
}

func cvtUint64Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint64)(sPtr))
}

func cvtUint64Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint64)(sPtr))
}

func cvtUint64Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint64)(sPtr))
}

func cvtUint64Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint64)(sPtr))
}

func cvtUint64Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint64)(sPtr))
}

func cvtUint64Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint64)(sPtr))
}

func cvtUint64Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint64)(sPtr))
}

func cvtUint64Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint64)(sPtr))
}

func cvtUint64Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint64)(sPtr))
}

func cvtUint64Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint64)(sPtr))
}

func cvtUint64Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint64)(sPtr))
}

func cvtUint64Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint64)(sPtr))
}

func cvtUint64Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint64)(sPtr))
}

func cvtUint64String(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*uint64)(sPtr))
}

func cvtFloat32Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*float32)(sPtr))
}

func cvtFloat32Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*float32)(sPtr))
}

func cvtFloat32Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*float32)(sPtr))
}

func cvtFloat32Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*float32)(sPtr))
}

func cvtFloat32Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*float32)(sPtr))
}

func cvtFloat32Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*float32)(sPtr))
}

func cvtFloat32Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*float32)(sPtr))
}

func cvtFloat32Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*float32)(sPtr))
}

func cvtFloat32Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*float32)(sPtr))
}

func cvtFloat32Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*float32)(sPtr))
}

func cvtFloat32Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*float32)(sPtr))
}

func cvtFloat32Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*float32)(sPtr))
}

func cvtFloat32Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*float32)(sPtr))
}

func cvtFloat64Int(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*float64)(sPtr))
}

func cvtFloat64Uint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*float64)(sPtr))
}

func cvtFloat64Int8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*float64)(sPtr))
}

func cvtFloat64Uint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*float64)(sPtr))
}

func cvtFloat64Int16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*float64)(sPtr))
}

func cvtFloat64Uint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*float64)(sPtr))
}

func cvtFloat64Int32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*float64)(sPtr))
}

func cvtFloat64Uint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*float64)(sPtr))
}

func cvtFloat64Int64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*float64)(sPtr))
}

func cvtFloat64Uint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*float64)(sPtr))
}

func cvtFloat64Float32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*float64)(sPtr))
}

func cvtFloat64Float64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*float64)(sPtr))
}

func cvtFloat64Uintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*float64)(sPtr))
}

func cvtComplex64Complex64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*complex64)(dPtr) = (complex64)(*(*complex64)(sPtr))
}

func cvtComplex64Complex128(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*complex128)(dPtr) = (complex128)(*(*complex64)(sPtr))
}

func cvtComplex128Complex64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*complex64)(dPtr) = (complex64)(*(*complex128)(sPtr))
}

func cvtComplex128Complex128(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*complex128)(dPtr) = (complex128)(*(*complex128)(sPtr))
}

func cvtUintptrInt(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uintptr)(sPtr))
}

func cvtUintptrUint(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uintptr)(sPtr))
}

func cvtUintptrInt8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uintptr)(sPtr))
}

func cvtUintptrUint8(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uintptr)(sPtr))
}

func cvtUintptrInt16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uintptr)(sPtr))
}

func cvtUintptrUint16(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uintptr)(sPtr))
}

func cvtUintptrInt32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uintptr)(sPtr))
}

func cvtUintptrUint32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uintptr)(sPtr))
}

func cvtUintptrInt64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uintptr)(sPtr))
}

func cvtUintptrUint64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uintptr)(sPtr))
}

func cvtUintptrFloat32(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uintptr)(sPtr))
}

func cvtUintptrFloat64(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uintptr)(sPtr))
}

func cvtUintptrUintptr(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uintptr)(sPtr))
}

func cvtUintptrString(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*uintptr)(sPtr))
}

func cvtStringString(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*string)(sPtr))
}

func newValue(k reflect.Kind) unsafe.Pointer {
	switch k {
	case reflect.Bool:
		return newBool()
	case reflect.Int:
		return newInt()
	case reflect.Uint:
		return newUint()
	case reflect.Int8:
		return newInt8()
	case reflect.Uint8:
		return newUint8()
	case reflect.Int16:
		return newInt16()
	case reflect.Uint16:
		return newUint16()
	case reflect.Int32:
		return newInt32()
	case reflect.Uint32:
		return newUint32()
	case reflect.Int64:
		return newInt64()
	case reflect.Uint64:
		return newUint64()
	case reflect.Float32:
		return newFloat32()
	case reflect.Float64:
		return newFloat64()
	case reflect.Complex64:
		return newComplex64()
	case reflect.Complex128:
		return newComplex128()
	case reflect.Uintptr:
		return newUintptr()
	case reflect.String:
		return newString()
	default:
		return nil
	}
}

func newBool() unsafe.Pointer {
	v := false
	return unsafe.Pointer(&v)
}

func newInt() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newUint() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newInt8() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newUint8() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newInt16() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newUint16() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newInt32() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newUint32() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newInt64() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newUint64() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newFloat32() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newFloat64() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newComplex64() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newComplex128() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newUintptr() unsafe.Pointer {
	v := 0
	return unsafe.Pointer(&v)
}

func newString() unsafe.Pointer {
	v := ""
	return unsafe.Pointer(&v)
}
