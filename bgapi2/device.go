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
	"unsafe"
)

type Device struct {
	ptr *C.BGAPI2_Device
}

type DeviceEvent struct {
	ptr *C.BGAPI2_DeviceEvent
}

func (d *Device) Open() error {
	result := C.BGAPI2_Device_Open(d.ptr)
	return ResultToError(result)
}

func (d *Device) OpenExclusive() error {
	result := C.BGAPI2_Device_OpenExclusive(d.ptr)
	return ResultToError(result)
}

func (d *Device) OpenReadOnly() error {
	result := C.BGAPI2_Device_OpenReadOnly(d.ptr)
	return ResultToError(result)
}

func (d *Device) IsOpen() (bool, error) {
	var isOpen C.bo_bool
	result := C.BGAPI2_Device_IsOpen(d.ptr, &isOpen)
	return isOpen == 1, ResultToError(result)
}

func (d *Device) GetDataStream(index uint) (*DataStream, error) {
	var dataStream *C.BGAPI2_DataStream

	result := C.BGAPI2_Device_GetDataStream(d.ptr, C.bo_uint(index), &dataStream)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return getDataStreamFromCDataStream(dataStream), err
}

func (d *Device) GetNumDataStreams() (uint, error) {
	var count C.bo_uint
	result := C.BGAPI2_Device_GetNumDataStreams(d.ptr, &count)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint(count), err
}

func (d *Device) Close() error {
	result := C.BGAPI2_Device_Close(d.ptr)
	return ResultToError(result)
}

func (d *Device) GetNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	node := &Node{}

	result := C.BGAPI2_Device_GetNode(d.ptr, cName, &node.ptr)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return node, err
}

func (d *Device) GetNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_Device_GetNodeTree(d.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (d *Device) GetNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Device_GetNodeList(d.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

//BGAPI2_Device_SetDeviceEventMode
//BGAPI2_Device_GetDeviceEventMode
//BGAPI2_CreateDeviceEvent
//BGAPI2_ReleaseDeviceEvent
//BGAPI2_Device_GetDeviceEvent
//BGAPI2_Device_CancelGetDeviceEvent
//BGAPI2_Device_RegisterDeviceEventHandler

func (d *Device) GetPayloadSize() (uint64, error) {
	var payloadSize C.bo_uint64
	result := C.BGAPI2_Device_GetPayloadSize(d.ptr, &payloadSize)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint64(payloadSize), err
}

func (d *Device) GetRemoteNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var node *C.BGAPI2_Node

	result := C.BGAPI2_Device_GetRemoteNode(d.ptr, cName, &node)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{node}, err
}

func (d *Device) GetRemoteNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_Device_GetRemoteNodeTree(d.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (d *Device) GetRemoteNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Device_GetRemoteNodeList(d.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (d *Device) GetId() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetID(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetVendor() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetVendor(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetModel() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetModel(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetSerialNumber() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetSerialNumber(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetTLType() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetTLType(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetDisplayName() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetDisplayName(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetAccessStatus() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetAccessStatus(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetRemoteConfigurationFile() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetRemoteConfigurationFile(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) SetRemoteConfigurationFile(configFile string) error {
	cConfigFile := C.CString(configFile)
	defer C.free(unsafe.Pointer(cConfigFile))

	result := C.BGAPI2_Device_SetRemoteConfigurationFile(d.ptr, cConfigFile)
	return ResultToError(result)
}

func (d *Device) StartStacking(replaceMode bool) error {
	var value C.bo_bool
	if replaceMode {
		value = 1
	} else {
		value = 0
	}
	result := C.BGAPI2_Device_StartStacking(d.ptr, value)
	return ResultToError(result)
}

func (d *Device) WriteStack() error {
	result := C.BGAPI2_Device_WriteStack(d.ptr)
	return ResultToError(result)
}

func (d *Device) CancelStack() error {
	result := C.BGAPI2_Device_CancelStack(d.ptr)
	return ResultToError(result)
}

func (d *Device) IsUpdateModeAvailable() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Device_IsUpdateModeAvailable(d.ptr, &value)
	return value == 1, ResultToError(result)
}

func (d *Device) IsUpdateModeActive() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Device_IsUpdateModeActive(d.ptr, &value)
	return value == 1, ResultToError(result)
}

func (d *Device) SetUpdateMode(updateMode bool, customKey string) error {
	var value C.bo_bool
	if updateMode {
		value = 1
	} else {
		value = 0
	}
	cCustomKey := C.CString(customKey)
	defer C.free(unsafe.Pointer(cCustomKey))
	result := C.BGAPI2_Device_SetUpdateMode(d.ptr, value, cCustomKey)
	return ResultToError(result)
}

func (d *Device) GetUpdateNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var node *C.BGAPI2_Node

	result := C.BGAPI2_Device_GetUpdateNode(d.ptr, cName, &node)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{node}, err
}

func (d *Device) GetUpdateNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_Device_GetUpdateNodeTree(d.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (d *Device) GetUpdateNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Device_GetUpdateNodeList(d.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (d *Device) GetUpdateConfigurationFile() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Device_GetUpdateConfigurationFile(d.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (d *Device) GetParent() (*Interface, error) {
	var parent *C.BGAPI2_Interface

	result := C.BGAPI2_Device_GetParent(d.ptr, &parent)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Interface{parent}, nil
}
