package bgapi2

/*
#cgo CFLAGS: -I/opt/baumer-gapi-sdk-c/include
#cgo LDFLAGS: -L/opt/baumer-gapi-sdk-c/lib -lbgapi2_genicam
#include "bgapi2_genicam/bgapi2_types.h"
#include "bgapi2_genicam/bgapi2_genicam.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Interface struct {
	ptr *C.BGAPI2_Interface
}

func (i *Interface) Open() error {
	result := C.BGAPI2_Interface_Open(i.ptr)
	return ResultToError(result)
}

func (i *Interface) IsOpen() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Interface_IsOpen(i.ptr, &value)
	return value == 1, ResultToError(result)
}

func (i *Interface) UpdateDeviceList(timeout uint64) (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Interface_UpdateDeviceList(i.ptr, &value, C.bo_uint64(timeout))
	return value == 1, ResultToError(result)
}

func (i *Interface) GetDevice(index uint) (*Device, error) {
	var device *C.BGAPI2_Device

	result := C.BGAPI2_Interface_GetDevice(i.ptr, C.bo_uint(index), &device)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Device{device}, nil
}

func (i *Interface) GetNumDevices() (uint, error) {
	var value C.bo_uint
	result := C.BGAPI2_Interface_GetNumDevices(i.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

func (i *Interface) GetDevices() ([]*Device, error) {
	_, err := i.UpdateDeviceList(timeout)
	if err != nil {
		return nil, err
	}

	numDevices, err := i.GetNumDevices()
	if err != nil {
		return nil, err
	}

	devices := make([]*Device, numDevices)
	for index, _ := range devices {
		devices[index], err = i.GetDevice(uint(index))
		if err != nil {
			return nil, err
		}
	}

	return devices, nil
}

func (i *Interface) GetParent() (*System, error) {
	var parent *C.BGAPI2_System

	result := C.BGAPI2_Interface_GetParent(i.ptr, &parent)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &System{parent}, nil
}

func (i *Interface) Close() error {
	result := C.BGAPI2_Interface_Close(i.ptr)
	return ResultToError(result)
}

func (i *Interface) GetNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var node *C.BGAPI2_Node

	result := C.BGAPI2_Interface_GetNode(i.ptr, cName, &node)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{node}, nil
}

func (i *Interface) GetNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_Interface_GetNodeTree(i.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (i *Interface) GetNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Interface_GetNodeList(i.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (i *Interface) GetId() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Interface_GetID(i.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (i *Interface) GetDisplayName() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Interface_GetDisplayName(i.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (i *Interface) GetTLType() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Interface_GetTLType(i.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}
