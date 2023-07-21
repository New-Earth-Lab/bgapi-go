package bgapi2

/*
#cgo CFLAGS: -I/opt/baumer-gapi-sdk-c/include
#cgo LDFLAGS: -L/opt/baumer-gapi-sdk-c/lib -lbgapi2_genicam
#include "bgapi2_genicam/bgapi2_types.h"
#include "bgapi2_genicam/bgapi2_genicam.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/lirm/aeron-go/aeron/atomic"
)

var buffers = sync.Map{}

type Buffer struct {
	ptr        *C.BGAPI2_Buffer
	userObject any
}

func CreateBuffer() (*Buffer, error) {
	buffer := new(Buffer)

	result := C.BGAPI2_CreateBuffer(&buffer.ptr)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	buffers.Store(buffer.ptr, buffer)

	return buffer, nil
}

func getBufferFromCBuffer(ptr *C.BGAPI2_Buffer) (*Buffer, error) {
	v, ok := buffers.Load(ptr)
	if !ok {
		return nil, fmt.Errorf("bgapi2: Unable to get buffer")
	}

	buffer := v.(*Buffer)

	return buffer, nil
}

func CreateBufferWithUserPtr(userObject any) (*Buffer, error) {
	buffer := new(Buffer)

	result := C.BGAPI2_CreateBufferWithUserPtr(&buffer.ptr, unsafe.Pointer(nil))
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	buffers.Store(buffer.ptr, buffer)

	return buffer, nil
}

func CreateBufferWithOptimizedSize(userObject any) (*Buffer, error) {
	buffer := new(Buffer)

	result := C.BGAPI2_CreateBufferWithOptimizedSize(&buffer.ptr, unsafe.Pointer(nil))
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	buffers.Store(buffer.ptr, buffer)

	return buffer, nil
}

func (b *Buffer) Delete() error {
	var userObj unsafe.Pointer

	buffers.Delete(b.ptr)
	result := C.BGAPI2_DeleteBuffer(b.ptr, &userObj)
	return ResultToError(result)
}

func (b *Buffer) GetNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var node *C.BGAPI2_Node

	result := C.BGAPI2_Buffer_GetNode(b.ptr, cName, &node)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{node}, nil
}

func (b *Buffer) GetNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_Buffer_GetNodeTree(b.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (b *Buffer) GetNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Buffer_GetNodeList(b.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (b *Buffer) GetChunkNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Buffer_GetChunkNodeList(b.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (b *Buffer) GetId() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Buffer_GetID(b.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (b *Buffer) GetMemPtr() (unsafe.Pointer, error) {
	var value unsafe.Pointer
	result := C.BGAPI2_Buffer_GetMemPtr(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (b *Buffer) GetMemBuffer() (*atomic.Buffer, error) {
	ptr, err := b.GetMemPtr()
	if err != nil {
		return nil, err
	}

	size, err := b.GetMemSize()
	if err != nil {
		return nil, err
	}

	return atomic.MakeBuffer(ptr, size), nil
}

func (b *Buffer) GetMemImageBuffer() (*atomic.Buffer, error) {
	ptr, err := b.GetMemPtr()
	if err != nil {
		return nil, err
	}

	size, err := b.GetImageLength()
	if err != nil {
		return nil, err
	}

	offset, err := b.GetImageOffset()
	if err != nil {
		return nil, err
	}

	return atomic.MakeBuffer(unsafe.Add(ptr, offset), size), nil
}

func (b *Buffer) GetMemSize() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetMemSize(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetUserObject() (any, error) {
	return b.userObject, nil
}

func (b *Buffer) GetTimestamp() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetTimestamp(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetHostTimestamp() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetHostTimestamp(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetNewData() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Buffer_GetNewData(b.ptr, &value)
	return value == 1, ResultToError(result)
}

func (b *Buffer) GetIsQueued() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Buffer_GetIsQueued(b.ptr, &value)
	return value == 1, ResultToError(result)
}

func (b *Buffer) GetIsAcquiring() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Buffer_GetIsAcquiring(b.ptr, &value)
	return value == 1, ResultToError(result)
}

func (b *Buffer) GetIsIncomplete() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Buffer_GetIsIncomplete(b.ptr, &value)
	return value == 1, ResultToError(result)
}

func (b *Buffer) GetTLType() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Buffer_GetTLType(b.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (b *Buffer) GetSizeFilled() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetSizeFilled(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetWidth() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetWidth(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetHeight() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetHeight(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetXOffset() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetXOffset(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetYOffset() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetYOffset(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetXPadding() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetXPadding(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetYPadding() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetYPadding(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetFrameId() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetFrameID(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetImagePresent() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Buffer_GetImagePresent(b.ptr, &value)
	return value == 1, ResultToError(result)
}

func (b *Buffer) GetImageOffset() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetImageOffset(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetImageLength() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetImageLength(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetPayloadType() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Buffer_GetPayloadType(b.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (b *Buffer) GetPixelFormat() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Buffer_GetPixelFormat(b.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (b *Buffer) GetDeliveredImageHeight() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetDeliveredImageHeight(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetDeliveredChunkPayloadSize() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetDeliveredChunkPayloadSize(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) ContainsChunk() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Buffer_GetContainsChunk(b.ptr, &value)
	return value == 1, ResultToError(result)
}

func (b *Buffer) GetChunkLayoutId() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Buffer_GetChunkLayoutID(b.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (b *Buffer) GetFileName() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Buffer_GetFileName(b.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (b *Buffer) GetParent() (*DataStream, error) {
	var parent *C.BGAPI2_DataStream

	result := C.BGAPI2_Buffer_GetParent(b.ptr, &parent)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return getDataStreamFromCDataStream(parent), nil
}
