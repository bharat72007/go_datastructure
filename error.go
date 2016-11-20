package go_datastructure

import "errors"

var (
	ErrorOverFlow  = errors.New("DataStructure is full")
	ErrorUnderFlow = errors.New("DatatStructure is Empty")
)

type InvalidOperationError struct {
	ds_name   string
	oper_name string
}

func (e *InvalidOperationError) Error() string {
	return e.ds_name + ":" + e.oper_name
}
