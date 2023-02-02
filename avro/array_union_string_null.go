// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     stream_data_record_message_schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUnionStringNull(r []*UnionStringNull, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionStringNull(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionStringNullWrapper struct {
	Target *[]*UnionStringNull
}

func (_ ArrayUnionStringNullWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayUnionStringNullWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionStringNullWrapper) Finalize()        {}
func (_ ArrayUnionStringNullWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayUnionStringNullWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionStringNull, 0, s)
	}
}
func (r ArrayUnionStringNullWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionStringNullWrapper) AppendArray() types.Field {
	var v *UnionStringNull
	v = NewUnionStringNull()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
