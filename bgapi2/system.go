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

const timeout = 100

type System struct {
	ptr *C.BGAPI2_System
}

func UpdateSystemList() error {
	result := C.BGAPI2_UpdateSystemList()
	return ResultToError(result)
}

func UpdateSystemListFromPath(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	result := C.BGAPI2_UpdateSystemListFromPath(cPath)

	return ResultToError(result)
}

func LoadSystemFromPath(path string) (*System, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	var system *C.BGAPI2_System

	result := C.BGAPI2_LoadSystemFromPath(cPath, &system)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &System{system}, nil
}

func GetNumSystems() (uint, error) {
	var value C.bo_uint
	result := C.BGAPI2_GetNumSystems(&value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

func GetSystem(index uint) (*System, error) {
	var system *C.BGAPI2_System

	result := C.BGAPI2_GetSystem(C.bo_uint(index), &system)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &System{system}, nil
}

func GetSystems() ([]*System, error) {
	err := UpdateSystemList()
	if err != nil {
		return nil, err
	}

	numSystems, err := GetNumSystems()
	if err != nil {
		return nil, err
	}

	systems := make([]*System, numSystems)
	for i, _ := range systems {
		systems[i], err = GetSystem(uint(i))
		if err != nil {
			return nil, err
		}
	}

	return systems, nil
}

func (s *System) Open() error {
	result := C.BGAPI2_System_Open(s.ptr)
	return ResultToError(result)
}

func (s *System) IsOpen() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_System_IsOpen(s.ptr, &value)
	return value == 1, ResultToError(result)
}

func (s *System) UpdateInterfaceList(timeout uint64) (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_System_UpdateInterfaceList(s.ptr, &value, C.bo_uint64(timeout))
	return value == 1, ResultToError(result)
}

func (s *System) GetInterface(index uint) (*Interface, error) {
	var iface *C.BGAPI2_Interface

	result := C.BGAPI2_System_GetInterface(s.ptr, C.bo_uint(index), &iface)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Interface{iface}, nil
}

func (s *System) GetNumInterfaces() (uint, error) {
	var value C.bo_uint
	result := C.BGAPI2_System_GetNumInterfaces(s.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

func (s *System) GetInterfaces() ([]*Interface, error) {
	_, err := s.UpdateInterfaceList(timeout)
	if err != nil {
		return nil, err
	}

	numInterfaces, err := s.GetNumInterfaces()
	if err != nil {
		return nil, err
	}

	interfaces := make([]*Interface, numInterfaces)
	for i, _ := range interfaces {
		interfaces[i], err = s.GetInterface(uint(i))
		if err != nil {
			return nil, err
		}
	}

	return interfaces, nil
}

func (s *System) Close() error {
	result := C.BGAPI2_System_Close(s.ptr)
	return ResultToError(result)
}

func (s *System) ReleaseSystem() error {
	result := C.BGAPI2_ReleaseSystem(s.ptr)
	s.ptr = nil
	return ResultToError(result)
}

func (s *System) GetNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var node *C.BGAPI2_Node

	result := C.BGAPI2_System_GetNode(s.ptr, cName, &node)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{node}, nil
}

func (s *System) GetNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_System_GetNodeTree(s.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (s *System) GetNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_System_GetNodeList(s.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (s *System) GetId() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetID(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (s *System) GetVendor() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetVendor(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (s *System) GetModel() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetModel(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (s *System) GetVersion() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetVersion(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (s *System) GetTLType() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetTLType(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (s *System) GetFileName() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetFileName(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (s *System) GetPathName() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetPathName(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (s *System) GetDisplayName() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_System_GetDisplayName(s.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}
