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

type Node struct {
	ptr *C.BGAPI2_Node
}

func (n *Node) GetInterface() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetInterface(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetExtension() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetExtension(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetToolTip() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetToolTip(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetDescription() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetDescription(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetDisplayName() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetDisplayname(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetVisibility() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetVisibility(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetEventId() (int64, error) {
	var value C.bo_int64
	result := C.BGAPI2_Node_GetEventID(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return int64(value), nil
}

func (n *Node) GetImplmented() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_GetImplemented(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetAvailable() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_GetAvailable(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetLocked() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_GetLocked(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetImposedAccessMode() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetImposedAccessMode(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetCurrentAccessMode() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetCurrentAccessMode(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) IsReadable() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_IsReadable(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) IsWriteable() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_IsWriteable(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetAlias() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetAlias(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetValue() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetValue(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) SetValue(value string) error {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	result := C.BGAPI2_Node_SetValue(n.ptr, cValue)
	return ResultToError(result)
}

func (n *Node) GetRepresentation() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetRepresentation(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetIntMin() (int64, error) {
	var value C.bo_int64
	result := C.BGAPI2_Node_GetIntMin(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return int64(value), nil
}

func (n *Node) GetIntMax() (int64, error) {
	var value C.bo_int64
	result := C.BGAPI2_Node_GetIntMax(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return int64(value), nil
}

func (n *Node) GetIntInc() (int64, error) {
	var value C.bo_int64
	result := C.BGAPI2_Node_GetIntInc(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return int64(value), nil
}

func (n *Node) GetInt() (int64, error) {
	var value C.bo_int64
	result := C.BGAPI2_Node_GetInt(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return int64(value), nil
}

func (n *Node) SetInt(value int64) error {
	result := C.BGAPI2_Node_SetInt(n.ptr, C.bo_int64(value))
	return ResultToError(result)
}

func (n *Node) HasUnit() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_HasUnit(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetUnit() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetUnit(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) GetDoubleMin() (float64, error) {
	var value C.bo_double
	result := C.BGAPI2_Node_GetDoubleMin(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return float64(value), nil
}

func (n *Node) GetDoubleMax() (float64, error) {
	var value C.bo_double
	result := C.BGAPI2_Node_GetDoubleMax(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return float64(value), nil
}

func (n *Node) GetDoubleInc() (float64, error) {
	var value C.bo_double
	result := C.BGAPI2_Node_GetDoubleInc(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return float64(value), nil
}

func (n *Node) HasInc() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_HasInc(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetDoublePrecision() (uint64, error) {
	var value C.bo_uint64
	result := C.BGAPI2_Node_GetDoublePrecision(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return uint64(value), nil
}

func (n *Node) GetDouble() (float64, error) {
	var value C.bo_double
	result := C.BGAPI2_Node_GetDouble(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return float64(value), nil
}

func (n *Node) SetDouble(value float64) error {
	result := C.BGAPI2_Node_SetDouble(n.ptr, C.bo_double(value))
	return ResultToError(result)
}

func (n *Node) GetMaxStringLength() (int64, error) {
	var value C.bo_int64
	result := C.BGAPI2_Node_GetMaxStringLength(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return int64(value), nil
}

func (n *Node) GetString() (string, error) {
	var varStringLen C.ulong = strLen
	varString := (*C.char)(C.malloc(varStringLen))
	defer C.free(unsafe.Pointer(varString))

	result := C.BGAPI2_Node_GetString(n.ptr, varString, &varStringLen)
	err := ResultToError(result)
	if err != nil {
		return "", err
	}

	return C.GoString(varString), nil
}

func (n *Node) SetString(value string) error {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	result := C.BGAPI2_Node_SetString(n.ptr, cValue)
	return ResultToError(result)
}

func (n *Node) GetEnumNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Node_GetEnumNodeList(n.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (n *Node) Execute() error {
	result := C.BGAPI2_Node_Execute(n.ptr)
	return ResultToError(result)
}

func (n *Node) IsDone() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_IsDone(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetBool() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_GetBool(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) SetBool(val bool) error {
	var value C.bo_bool
	if val {
		value = 1
	} else {
		value = 0
	}
	result := C.BGAPI2_Node_SetBool(n.ptr, value)
	return ResultToError(result)
}

func (n *Node) GetNodeTree() (*NodeMap, error) {
	var tree *C.BGAPI2_NodeMap
	result := C.BGAPI2_Node_GetNodeTree(n.ptr, &tree)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{tree}, nil
}

func (n *Node) GetNodeList() (*NodeMap, error) {
	var list *C.BGAPI2_NodeMap
	result := C.BGAPI2_Node_GetNodeList(n.ptr, &list)
	err := ResultToError(result)
	if err != nil {
		return nil, err
	}

	return &NodeMap{list}, nil
}

func (n *Node) IsSelector() (bool, error) {
	var value C.bo_bool
	result := C.BGAPI2_Node_IsSelector(n.ptr, &value)
	return value == 1, ResultToError(result)
}

func (n *Node) GetLength() (int64, error) {
	var value C.bo_int64
	result := C.BGAPI2_Node_GetLength(n.ptr, &value)
	err := ResultToError(result)
	if err != nil {
		return 0, err
	}

	return int64(value), nil
}
