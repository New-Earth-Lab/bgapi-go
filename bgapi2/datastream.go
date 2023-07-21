package bgapi2

/*
#cgo CFLAGS: -I/opt/baumer-gapi-sdk-c/include
#cgo LDFLAGS: -L/opt/baumer-gapi-sdk-c/lib -lbgapi2_genicam
#include "bgapi2_genicam/bgapi2_types.h"
#include "bgapi2_genicam/bgapi2_genicam.h"
#include <stdlib.h>

extern void bufferEventHandlerWrapper(void* callbackOwner, BGAPI2_Buffer* pBuffer);

*/
import "C"
import (
	"sync"
	"unsafe"
)

type BufferEventHandler func(dataStream *DataStream, buffer *Buffer, userObject any)

var dataStreams = sync.Map{}

type DataStream struct {
	ptr          *C.BGAPI2_DataStream
	userData     any
	eventHandler BufferEventHandler
}

func getDataStreamFromCDataStream(ptr *C.BGAPI2_DataStream) *DataStream {
	d := &DataStream{
		ptr: ptr,
	}

	val, _ := dataStreams.LoadOrStore(ptr, d)

	return val.(*DataStream)
}

func (d *DataStream) Open() error {
	result := C.BGAPI2_DataStream_Open(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) Close() error {
	result := C.BGAPI2_DataStream_Close(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) IsOpen() (bool, error) {
	var isOpen C.bo_bool
	result := C.BGAPI2_DataStream_IsOpen(d.ptr, &isOpen)
	return isOpen == 1, ResultToError(result)
}

func (d *DataStream) GetNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var node *C.BGAPI2_Node

	result := C.BGAPI2_DataStream_GetNode(d.ptr, cName, &node)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{node}, nil
}

func (d *DataStream) GetNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_DataStream_GetNodeTree(d.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (d *DataStream) GetNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_DataStream_GetNodeList(d.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (d *DataStream) SetNewBufferEventMode(eventMode int) error {
	var cEventMode C.BGAPI2_EventMode
	switch eventMode {
	case EventMode_Unregistered:
		cEventMode = C.EVENTMODE_UNREGISTERED
	case EventMode_Polling:
		cEventMode = C.EVENTMODE_POLLING
	case EventMode_EventHandler:
		cEventMode = C.EVENTMODE_EVENT_HANDLER
	}
	result := C.BGAPI2_DataStream_SetNewBufferEventMode(d.ptr, cEventMode)
	return ResultToError(result)
}

func (d *DataStream) GetNewBufferEventMode() (int, error) {
	var value C.BGAPI2_EventMode
	result := C.BGAPI2_DataStream_GetNewBufferEventMode(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	var retval int
	switch value {
	case C.EVENTMODE_UNREGISTERED:
		retval = EventMode_Unregistered
	case C.EVENTMODE_POLLING:
		retval = EventMode_Polling
	case C.EVENTMODE_EVENT_HANDLER:
		retval = EventMode_EventHandler
	}
	return retval, nil
}

func (d *DataStream) GetId() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_DataStream_GetID(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *DataStream) GetNumDelivered() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_DataStream_GetNumDelivered(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}

func (d *DataStream) GetNumUnderrun() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_DataStream_GetNumUnderrun(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}

func (d *DataStream) GetNumAnnounced() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_DataStream_GetNumAnnounced(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}

func (d *DataStream) GetNumQueued() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_DataStream_GetNumQueued(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}

func (d *DataStream) GetNumAwaitDelivery() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_DataStream_GetNumAwaitDelivery(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}

func (d *DataStream) GetNumStarted() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_DataStream_GetNumStarted(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}

func (d *DataStream) GetPayloadSize() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_DataStream_GetPayloadSize(d.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}

func (d *DataStream) IsGrabbing() (bool, error) {
	var isOpen C.bo_bool
	result := C.BGAPI2_DataStream_GetIsGrabbing(d.ptr, &isOpen)
	return isOpen == 1, ResultToError(result)
}

func (d *DataStream) DefinesPayloadSize() (bool, error) {
	var isOpen C.bo_bool
	result := C.BGAPI2_DataStream_GetDefinesPayloadSize(d.ptr, &isOpen)
	return isOpen == 1, ResultToError(result)
}

func (d *DataStream) GetTLType() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_DataStream_GetTLType(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *DataStream) StartAcquisition(numToAcquire uint64) error {
	result := C.BGAPI2_DataStream_StartAcquisition(d.ptr, C.bo_uint64(numToAcquire))
	return ResultToError(result)
}

func (d *DataStream) StartAcquisitionContinuous() error {
	result := C.BGAPI2_DataStream_StartAcquisitionContinuous(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) StopAcquisition() error {
	result := C.BGAPI2_DataStream_StopAcquisition(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) AbortAcquisition() error {
	result := C.BGAPI2_DataStream_AbortAcquisition(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) FlushToOutputQueue() error {
	result := C.BGAPI2_DataStream_FlushInputToOutputQueue(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) FlushAllToInputQueue() error {
	result := C.BGAPI2_DataStream_FlushAllToInputQueue(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) FlushUnqueuedToInputQueue() error {
	result := C.BGAPI2_DataStream_FlushUnqueuedToInputQueue(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) DiscardOutputBuffers() error {
	result := C.BGAPI2_DataStream_DiscardOutputBuffers(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) DiscardAllBuffers() error {
	result := C.BGAPI2_DataStream_DiscardAllBuffers(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) AnnounceBuffer(buffer *Buffer) error {
	result := C.BGAPI2_DataStream_AnnounceBuffer(d.ptr, buffer.ptr)
	return ResultToError(result)
}

func (d *DataStream) RevokeBuffer(buffer *Buffer) error {
	var userObj unsafe.Pointer
	result := C.BGAPI2_DataStream_RevokeBuffer(d.ptr, buffer.ptr, &userObj)
	return ResultToError(result)
}

func (d *DataStream) QueueBuffer(buffer *Buffer) error {
	result := C.BGAPI2_DataStream_QueueBuffer(d.ptr, buffer.ptr)
	return ResultToError(result)
}

func (d *DataStream) GetFilledBuffer(timeout uint64) (*Buffer, error) {
	var buffer *C.BGAPI2_Buffer
	result := C.BGAPI2_DataStream_GetFilledBuffer(d.ptr, &buffer, C.bo_uint64(timeout))
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return getBufferFromCBuffer(buffer)
}

func (d *DataStream) CancelGetFilledBuffer() error {
	result := C.BGAPI2_DataStream_CancelGetFilledBuffer(d.ptr)
	return ResultToError(result)
}

func (d *DataStream) GetBufferId(index uint) (*Buffer, error) {
	var buffer *C.BGAPI2_Buffer
	result := C.BGAPI2_DataStream_GetBufferID(d.ptr, C.bo_uint(index), &buffer)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return getBufferFromCBuffer(buffer)
}

func (d *DataStream) RegisterNewBufferEventHandler(handler BufferEventHandler, userData any) error {
	d.userData = userData
	d.eventHandler = handler
	result := C.BGAPI2_DataStream_RegisterNewBufferEventHandler(d.ptr, unsafe.Pointer(d.ptr),
		(C.BGAPI2_NewBufferEventHandler)(C.bufferEventHandlerWrapper))

	return ResultToError(result)
}

//export bufferEventHandler
//go:nocheckptr go:nosplit
func bufferEventHandler(callBackOwner unsafe.Pointer, pBuffer unsafe.Pointer) {
	d := getDataStreamFromCDataStream((*C.BGAPI2_DataStream)(callBackOwner))
	buffer, _ := getBufferFromCBuffer((*C.BGAPI2_Buffer)(pBuffer))

	if d.eventHandler != nil {
		d.eventHandler(d, buffer, d.userData)
	}

}

func (d *DataStream) DeleteBufferEventHandler() error {
	result := C.BGAPI2_DataStream_RegisterNewBufferEventHandler(d.ptr, unsafe.Pointer(nil),
		(C.BGAPI2_NewBufferEventHandler)(nil))
	d.eventHandler = nil
	d.userData = nil

	return ResultToError(result)
}

func (d *DataStream) GetParent() (*Device, error) {
	var parent *C.BGAPI2_Device

	result := C.BGAPI2_DataStream_GetParent(d.ptr, &parent)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Device{parent}, nil
}
