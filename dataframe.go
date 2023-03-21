package gopolars

/*
#cgo LDFLAGS: -lcpolars
#include "headers/libcpolars.h"
#include "headers/utils.h"
*/
import "C"
import (
	"unsafe"
)

type DataFrame struct {
	data unsafe.Pointer
}

func ReadCSV(path string) (*DataFrame, error) {
	readResult := C.lcp_read_csv(C.CString(path))

	if readResult.result_code != 0 {
		return nil, &PolarsError{
			message: C.GoString(readResult.message),
		}
	}

	return &DataFrame{
		data: readResult.data,
	}, nil
}

func (df *DataFrame) Columns(names ...string) (*DataFrame, error) {
	length := C.size_t(len(names))
	cArgs := C.alloc_str_array(length)
	for i, name := range names {
		C.set_str_array(cArgs, C.int(i), C.CString(name))
	}
	defer C.free_str_array(cArgs, length)

	result := C.lcp_columns(df.data, cArgs, length)
	if result.result_code != 0 {
		return nil, &PolarsError{
			message: C.GoString(result.message),
		}
	}

	return &DataFrame{
		data: result.data,
	}, nil
}

func (df *DataFrame) FrameEqual(other *DataFrame, missing bool) bool {
	imissing := 0
	if missing {
		imissing = 1
	}
	return C.lcp_frame_equal(df.data, other.data, C.int(imissing)) == 1
}

func (df *DataFrame) Head(n int) *DataFrame {
	return &DataFrame{
		C.lcp_head(df.data, C.int(n)),
	}
}

func (df *DataFrame) Lazy() *LazyFrame {
	lf := C.lcp_lazy(df.data)
	return &LazyFrame{
		data: lf,
	}
}

func (df *DataFrame) String() string {
	cstr := C.lcp_dataframe_to_str(df.data)
	return C.GoString(cstr)
}
