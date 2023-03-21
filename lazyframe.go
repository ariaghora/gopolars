package gopolars

/*
#cgo LDFLAGS: -lcpolars
#include "headers/libcpolars.h"
#include "headers/utils.h"
*/
import "C"
import "unsafe"

type LazyFrame struct {
	data unsafe.Pointer
}

func (lf *LazyFrame) Collect() (*DataFrame, error) {
	collectResult := C.lcp_collect(lf.data)

	if collectResult.result_code != 0 {
		return nil, &PolarsError{
			message: C.GoString(collectResult.message),
		}
	}

	return &DataFrame{
		data: collectResult.data,
	}, nil
}

func (lf *LazyFrame) Filter(predicate *Expr) *LazyFrame {
	return &LazyFrame{C.lcp_filter(lf.data, predicate.data)}
}

func (lf *LazyFrame) Select(exprs ...*Expr) *LazyFrame {
	len := C.int(len(exprs))
	pExprs := C.lcp_alloc_expr_arr(len)
	for i, e := range exprs {
		C.lcp_set_expr_arr(pExprs, e.data, C.int(i))
	}
	s := C.lcp_select(lf.data, pExprs, len)
	return &LazyFrame{s}
}
