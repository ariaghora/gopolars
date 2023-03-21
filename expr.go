package gopolars

/*
#cgo LDFLAGS: -lcpolars
#include "headers/libcpolars.h"
#include "headers/utils.h"
*/
import "C"
import "unsafe"

type Expr struct {
	data unsafe.Pointer
}

func Col(name string) *Expr {
	return &Expr{C.lcp_expr_column(C.CString(name))}
}

func Int(val int) *Expr {
	return &Expr{C.lcp_expr_i32(C.int(val))}
}

func Str(val string) *Expr {
	return &Expr{C.lcp_expr_str(C.CString(val))}
}

func (e *Expr) Alias(name string) *Expr {
	return &Expr{C.lcp_expr_alias(e.data, C.CString(name))}
}

func (e *Expr) Eq(other *Expr) *Expr {
	return &Expr{C.lcp_expr_eq(e.data, other.data)}
}
