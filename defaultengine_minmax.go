// Code generated by genlib2. DO NOT EDIT.

package tensor

import (
	"github.com/pkg/errors"
	"github.com/pdevine/tensor/internal/storage"
)

var (
	_ MinBetweener = StdEng{}
	_ MaxBetweener = StdEng{}
)

func (e StdEng) MinBetween(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "MinBetween failed")
	}

	var reuse DenseTensor
	var safe bool
	if reuse, safe, _, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.MinBetween")
	}
	// check to see if anything needs to be created
	if reuse == nil {
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && reuse == nil:
			err = e.E.MinBetweenIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.MinBetweenIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			panic("Unreachable")
		}
		return
	}

	// standard
	switch {
	case !safe && reuse == nil:
		err = e.E.MinBetween(typ, dataA, dataB)
		retVal = a
	case safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.MinBetween(typ, dataReuse, dataB)
		retVal = reuse
	default:
		panic("Unreachable")
	}
	return
}

func (e StdEng) MaxBetween(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "MaxBetween failed")
	}

	var reuse DenseTensor
	var safe bool
	if reuse, safe, _, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.MaxBetween")
	}
	// check to see if anything needs to be created
	if reuse == nil {
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && reuse == nil:
			err = e.E.MaxBetweenIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.MaxBetweenIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			panic("Unreachable")
		}
		return
	}

	// standard
	switch {
	case !safe && reuse == nil:
		err = e.E.MaxBetween(typ, dataA, dataB)
		retVal = a
	case safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.MaxBetween(typ, dataReuse, dataB)
		retVal = reuse
	default:
		panic("Unreachable")
	}
	return
}

func (e StdEng) MinBetweenScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "MinBetween failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "MinBetween failed")
	}

	var reuse DenseTensor
	var safe bool
	if reuse, safe, _, _, _, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter, newAlloc bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, newAlloc, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.MinBetween")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, newAlloc, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.MinBetween")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	if reuse == nil {
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && reuse == nil:
			err = e.E.MinBetweenIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.MinBetweenIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.MinBetweenIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			panic("Unreachable")
		}
		if newAlloc {
			freeScalar(scalarHeader.Raw)
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.MinBetween(typ, dataReuse, dataB)
			retVal = reuse
			return
		case safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.MinBetween(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && reuse == nil:
		err = e.E.MinBetween(typ, dataA, dataB)
		retVal = a
	case safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.MinBetween(typ, dataReuse, dataB)
		retVal = reuse
	case safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.MinBetween(typ, dataA, dataReuse)
		retVal = reuse
	default:
		panic("Unreachable")
	}
	if newAlloc {
		freeScalar(scalarHeader.Raw)
	}
	returnHeader(scalarHeader)
	return
}

func (e StdEng) MaxBetweenScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "MaxBetween failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "MaxBetween failed")
	}

	var reuse DenseTensor
	var safe bool
	if reuse, safe, _, _, _, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter, newAlloc bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, newAlloc, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.MaxBetween")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, newAlloc, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.MaxBetween")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	if reuse == nil {
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && reuse == nil:
			err = e.E.MaxBetweenIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.MaxBetweenIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.MaxBetweenIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			panic("Unreachable")
		}
		if newAlloc {
			freeScalar(scalarHeader.Raw)
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.MaxBetween(typ, dataReuse, dataB)
			retVal = reuse
			return
		case safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.MaxBetween(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && reuse == nil:
		err = e.E.MaxBetween(typ, dataA, dataB)
		retVal = a
	case safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.MaxBetween(typ, dataReuse, dataB)
		retVal = reuse
	case safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.MaxBetween(typ, dataA, dataReuse)
		retVal = reuse
	default:
		panic("Unreachable")
	}
	if newAlloc {
		freeScalar(scalarHeader.Raw)
	}
	returnHeader(scalarHeader)
	return
}
