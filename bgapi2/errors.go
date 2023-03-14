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
	"unsafe"
)

func ResultToError(result C.BGAPI2_RESULT) error {
	if result == C.BGAPI2_RESULT_SUCCESS {
		return nil
	}

	var errorCode C.BGAPI2_RESULT
	var varStringLen C.bo_uint64 = strLen

	varString := (*C.char)(C.malloc(strLen))
	defer C.free(unsafe.Pointer(varString))
	C.BGAPI2_GetLastError(&errorCode, varString, &varStringLen)

	return fmt.Errorf("bgapi2: %s", C.GoString(varString))
}

func ResultToErrorOld(result C.BGAPI2_RESULT) error {
	switch result {
	case C.BGAPI2_RESULT_SUCCESS:
		return nil
	case C.BGAPI2_RESULT_ERROR:
		return fmt.Errorf("bgapi: Error")
	case C.BGAPI2_RESULT_NOT_INITIALIZED:
		return fmt.Errorf("bgapi2: Not initialized")
	case C.BGAPI2_RESULT_NOT_IMPLEMENTED:
		return fmt.Errorf("bgapi2: Not implemented")
	case C.BGAPI2_RESULT_RESOURCE_IN_USE:
		return fmt.Errorf("bgapi2: Resource in use")
	case C.BGAPI2_RESULT_ACCESS_DENIED:
		return fmt.Errorf("bgapi2: Access denied")
	case C.BGAPI2_RESULT_INVALID_HANDLE:
		return fmt.Errorf("bgapi2: Invalid handle")
	case C.BGAPI2_RESULT_NO_DATA:
		return fmt.Errorf("bgapi2: No data")
	case C.BGAPI2_RESULT_INVALID_PARAMETER:
		return fmt.Errorf("bgapi2: Invalid parameter")
	case C.BGAPI2_RESULT_TIMEOUT:
		return fmt.Errorf("bgapi2: Timeout")
	case C.BGAPI2_RESULT_ABORT:
		return fmt.Errorf("bgapi2: Abort")
	case C.BGAPI2_RESULT_INVALID_BUFFER:
		return fmt.Errorf("bgapi2: Invalid buffer")
	case C.BGAPI2_RESULT_NOT_AVAILABLE:
		return fmt.Errorf("bgapi2: Not available")
	case C.BGAPI2_RESULT_OBJECT_INVALID:
		return fmt.Errorf("bgapi2: Object invalid")
	case C.BGAPI2_RESULT_LOWLEVEL_ERROR:
		return fmt.Errorf("bgapi2: Low-level error")
	default:
		return fmt.Errorf("bgapi2: Undefined error")
	}
}
