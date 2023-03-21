package gopolars

/*
#cgo LDFLAGS: -lcpolars
#include "headers/libcpolars.h"
#include "headers/utils.h"
*/
import "C"
import "unsafe"

type Series struct {
	data unsafe.Pointer
}

func NewIntSeries(name string, data []int) *Series {
	length := C.size_t(len(data))
	cName := C.CString(name)
	cData := C.alloc_int_array(length)
	defer C.free_int_array(cData)

	for i, val := range data {
		C.set_int_array(cData, C.int(i), C.int(val))
	}

	series := C.lcp_new_int_series(cName, cData, length)
	return &Series{
		data: series,
	}
}

func (s *Series) String() string {
	cstr := C.lcp_series_to_str(s.data)
	return C.GoString(cstr)
}
