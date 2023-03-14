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

type NodeMap struct {
	ptr *C.BGAPI2_NodeMap
}

func (n *NodeMap) GetNode(name string) (*Node, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var value *C.BGAPI2_Node
	result := C.BGAPI2_NodeMap_GetNode(n.ptr, cName, &value)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{value}, nil
}

func (n *NodeMap) GetNodeCount() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_NodeMap_GetNodeCount(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (n *NodeMap) GetNodeByIndex(index uint64) (*Node, error) {
	var value *C.BGAPI2_Node
	result := C.BGAPI2_NodeMap_GetNodeByIndex(n.ptr, C.bo_uint64(index), &value)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &Node{value}, nil
}

func (n *NodeMap) GetNodePresent(name string) (bool, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var value C.bo_bool
	result := C.BGAPI2_NodeMap_GetNodePresent(n.ptr, cName, &value)
	return value == 1, ResultToError(result)
}
