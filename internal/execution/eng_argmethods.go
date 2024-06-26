// Code generated by genlib2. DO NOT EDIT.

package execution

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pdevine/tensor/internal/storage"
)

func (e E) ArgmaxIter(t reflect.Type, a *storage.Header, it Iterator, lastSize int) (indices []int, err error) {
	var next int
	switch t {
	case Int:
		data := a.Ints()
		tmp := make([]int, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxI(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int8:
		data := a.Int8s()
		tmp := make([]int8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxI8(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int16:
		data := a.Int16s()
		tmp := make([]int16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxI16(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int32:
		data := a.Int32s()
		tmp := make([]int32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxI32(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int64:
		data := a.Int64s()
		tmp := make([]int64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxI64(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint:
		data := a.Uints()
		tmp := make([]uint, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxU(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint8:
		data := a.Uint8s()
		tmp := make([]uint8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxU8(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint16:
		data := a.Uint16s()
		tmp := make([]uint16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxU16(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint32:
		data := a.Uint32s()
		tmp := make([]uint32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxU32(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint64:
		data := a.Uint64s()
		tmp := make([]uint64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxU64(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float32:
		data := a.Float32s()
		tmp := make([]float32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxF32(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float64:
		data := a.Float64s()
		tmp := make([]float64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxF64(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case String:
		data := a.Strings()
		tmp := make([]string, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgmaxStr(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	default:
		return nil, errors.Errorf("Unsupported type %v for Argmax", t)
	}
}
func (e E) ArgminIter(t reflect.Type, a *storage.Header, it Iterator, lastSize int) (indices []int, err error) {
	var next int
	switch t {
	case Int:
		data := a.Ints()
		tmp := make([]int, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminI(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int8:
		data := a.Int8s()
		tmp := make([]int8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminI8(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int16:
		data := a.Int16s()
		tmp := make([]int16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminI16(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int32:
		data := a.Int32s()
		tmp := make([]int32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminI32(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int64:
		data := a.Int64s()
		tmp := make([]int64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminI64(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint:
		data := a.Uints()
		tmp := make([]uint, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminU(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint8:
		data := a.Uint8s()
		tmp := make([]uint8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminU8(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint16:
		data := a.Uint16s()
		tmp := make([]uint16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminU16(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint32:
		data := a.Uint32s()
		tmp := make([]uint32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminU32(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint64:
		data := a.Uint64s()
		tmp := make([]uint64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminU64(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float32:
		data := a.Float32s()
		tmp := make([]float32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminF32(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float64:
		data := a.Float64s()
		tmp := make([]float64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminF64(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case String:
		data := a.Strings()
		tmp := make([]string, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			if len(tmp) == lastSize {
				am := ArgminStr(tmp)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	default:
		return nil, errors.Errorf("Unsupported type %v for Argmin", t)
	}
}
func (e E) ArgmaxIterMasked(t reflect.Type, a *storage.Header, mask []bool, it Iterator, lastSize int) (indices []int, err error) {
	newMask := make([]bool, 0, lastSize)
	var next int
	switch t {
	case Int:
		data := a.Ints()
		tmp := make([]int, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedI(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int8:
		data := a.Int8s()
		tmp := make([]int8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedI8(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int16:
		data := a.Int16s()
		tmp := make([]int16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedI16(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int32:
		data := a.Int32s()
		tmp := make([]int32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedI32(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int64:
		data := a.Int64s()
		tmp := make([]int64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedI64(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint:
		data := a.Uints()
		tmp := make([]uint, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedU(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint8:
		data := a.Uint8s()
		tmp := make([]uint8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedU8(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint16:
		data := a.Uint16s()
		tmp := make([]uint16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedU16(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint32:
		data := a.Uint32s()
		tmp := make([]uint32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedU32(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint64:
		data := a.Uint64s()
		tmp := make([]uint64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedU64(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float32:
		data := a.Float32s()
		tmp := make([]float32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedF32(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float64:
		data := a.Float64s()
		tmp := make([]float64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedF64(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case String:
		data := a.Strings()
		tmp := make([]string, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgmaxMaskedStr(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	default:
		return nil, errors.Errorf("Unsupported type %v for Argmax", t)
	}
}
func (e E) ArgminIterMasked(t reflect.Type, a *storage.Header, mask []bool, it Iterator, lastSize int) (indices []int, err error) {
	newMask := make([]bool, 0, lastSize)
	var next int
	switch t {
	case Int:
		data := a.Ints()
		tmp := make([]int, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedI(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int8:
		data := a.Int8s()
		tmp := make([]int8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedI8(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int16:
		data := a.Int16s()
		tmp := make([]int16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedI16(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int32:
		data := a.Int32s()
		tmp := make([]int32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedI32(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Int64:
		data := a.Int64s()
		tmp := make([]int64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedI64(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint:
		data := a.Uints()
		tmp := make([]uint, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedU(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint8:
		data := a.Uint8s()
		tmp := make([]uint8, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedU8(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint16:
		data := a.Uint16s()
		tmp := make([]uint16, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedU16(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint32:
		data := a.Uint32s()
		tmp := make([]uint32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedU32(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Uint64:
		data := a.Uint64s()
		tmp := make([]uint64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedU64(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float32:
		data := a.Float32s()
		tmp := make([]float32, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedF32(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case Float64:
		data := a.Float64s()
		tmp := make([]float64, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedF64(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	case String:
		data := a.Strings()
		tmp := make([]string, 0, lastSize)
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			tmp = append(tmp, data[next])
			newMask = append(newMask, mask[next])
			if len(tmp) == lastSize {
				am := ArgminMaskedStr(tmp, mask)
				indices = append(indices, am)

				// reset
				tmp = tmp[:0]
				newMask = newMask[:0]
			}
		}
		if _, ok := err.(NoOpError); ok {
			err = nil
		}
		return
	default:
		return nil, errors.Errorf("Unsupported type %v for Argmin", t)
	}
}
func (e E) ArgmaxFlatMasked(t reflect.Type, a *storage.Header, mask []bool) (retVal int) {
	switch t {
	case Int:
		return ArgmaxMaskedI(a.Ints(), mask)
	case Int8:
		return ArgmaxMaskedI8(a.Int8s(), mask)
	case Int16:
		return ArgmaxMaskedI16(a.Int16s(), mask)
	case Int32:
		return ArgmaxMaskedI32(a.Int32s(), mask)
	case Int64:
		return ArgmaxMaskedI64(a.Int64s(), mask)
	case Uint:
		return ArgmaxMaskedU(a.Uints(), mask)
	case Uint8:
		return ArgmaxMaskedU8(a.Uint8s(), mask)
	case Uint16:
		return ArgmaxMaskedU16(a.Uint16s(), mask)
	case Uint32:
		return ArgmaxMaskedU32(a.Uint32s(), mask)
	case Uint64:
		return ArgmaxMaskedU64(a.Uint64s(), mask)
	case Float32:
		return ArgmaxMaskedF32(a.Float32s(), mask)
	case Float64:
		return ArgmaxMaskedF64(a.Float64s(), mask)
	case String:
		return ArgmaxMaskedStr(a.Strings(), mask)
	default:
		return -1
	}
}
func (e E) ArgminFlatMasked(t reflect.Type, a *storage.Header, mask []bool) (retVal int) {
	switch t {
	case Int:
		return ArgminMaskedI(a.Ints(), mask)
	case Int8:
		return ArgminMaskedI8(a.Int8s(), mask)
	case Int16:
		return ArgminMaskedI16(a.Int16s(), mask)
	case Int32:
		return ArgminMaskedI32(a.Int32s(), mask)
	case Int64:
		return ArgminMaskedI64(a.Int64s(), mask)
	case Uint:
		return ArgminMaskedU(a.Uints(), mask)
	case Uint8:
		return ArgminMaskedU8(a.Uint8s(), mask)
	case Uint16:
		return ArgminMaskedU16(a.Uint16s(), mask)
	case Uint32:
		return ArgminMaskedU32(a.Uint32s(), mask)
	case Uint64:
		return ArgminMaskedU64(a.Uint64s(), mask)
	case Float32:
		return ArgminMaskedF32(a.Float32s(), mask)
	case Float64:
		return ArgminMaskedF64(a.Float64s(), mask)
	case String:
		return ArgminMaskedStr(a.Strings(), mask)
	default:
		return -1
	}
}
func (e E) ArgmaxFlat(t reflect.Type, a *storage.Header) (retVal int) {
	switch t {
	case Int:
		return ArgmaxI(a.Ints())
	case Int8:
		return ArgmaxI8(a.Int8s())
	case Int16:
		return ArgmaxI16(a.Int16s())
	case Int32:
		return ArgmaxI32(a.Int32s())
	case Int64:
		return ArgmaxI64(a.Int64s())
	case Uint:
		return ArgmaxU(a.Uints())
	case Uint8:
		return ArgmaxU8(a.Uint8s())
	case Uint16:
		return ArgmaxU16(a.Uint16s())
	case Uint32:
		return ArgmaxU32(a.Uint32s())
	case Uint64:
		return ArgmaxU64(a.Uint64s())
	case Float32:
		return ArgmaxF32(a.Float32s())
	case Float64:
		return ArgmaxF64(a.Float64s())
	case String:
		return ArgmaxStr(a.Strings())
	default:
		return -1
	}
}
func (e E) ArgminFlat(t reflect.Type, a *storage.Header) (retVal int) {
	switch t {
	case Int:
		return ArgminI(a.Ints())
	case Int8:
		return ArgminI8(a.Int8s())
	case Int16:
		return ArgminI16(a.Int16s())
	case Int32:
		return ArgminI32(a.Int32s())
	case Int64:
		return ArgminI64(a.Int64s())
	case Uint:
		return ArgminU(a.Uints())
	case Uint8:
		return ArgminU8(a.Uint8s())
	case Uint16:
		return ArgminU16(a.Uint16s())
	case Uint32:
		return ArgminU32(a.Uint32s())
	case Uint64:
		return ArgminU64(a.Uint64s())
	case Float32:
		return ArgminF32(a.Float32s())
	case Float64:
		return ArgminF64(a.Float64s())
	case String:
		return ArgminStr(a.Strings())
	default:
		return -1
	}
}
