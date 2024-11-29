package root

import (
	"fmt"
	"math"
)

// Constants
var (
	// Precision of rott-finding
	Precision float64 = 1e-6

	// MaxIteration is max allowable amount of iteration.
	// Typically for precition=1e-6 need 20 iterations.
	//
	// Example:
	//
	//	go test -v -run="Test/Case_26"
	//
	//	It.  X value         Y value         Xerror
	//  0    5.000000e-01    5.714286e-03    1.000000e+00
	//  1    7.500000e-01    2.142857e-03    5.000000e-01
	//  2    8.750000e-01    3.571429e-04    2.500000e-01
	//  3    9.375000e-01   -5.357143e-04    1.250000e-01
	//  4    9.062500e-01   -8.928571e-05    6.666667e-02
	//  5    8.906250e-01    1.339286e-04    3.448276e-02
	//  6    8.984375e-01    2.232143e-05    1.724138e-02
	//  7    9.023438e-01   -3.348214e-05    8.620690e-03
	//  8    9.003906e-01   -5.580357e-06    4.329004e-03
	//  9    8.994141e-01    8.370536e-06    2.169197e-03
	// 10    8.999023e-01    1.395089e-06    1.084599e-03
	// 11    9.001465e-01   -2.092634e-06    5.422993e-04
	// 12    9.000244e-01   -3.487723e-07    2.712232e-04
	// 13    8.999634e-01    5.231585e-07    1.356300e-04
	// 14    8.999939e-01    8.719308e-08    6.781500e-05
	// 15    9.000092e-01   -1.307896e-07    3.390750e-05
	// 16    9.000015e-01   -2.179827e-08    1.695404e-05
	// 17    8.999977e-01    3.269741e-08    8.477091e-06
	// 18    8.999996e-01    5.449568e-09    4.238545e-06
	// 19    9.000006e-01   -8.174351e-09    2.119273e-06
	// 20    9.000001e-01   -1.362392e-09    1.059637e-06
	MaxIteration int = 500
)

type ErrorFind struct {
	Type ErrType
	Err  error
}

func (e ErrorFind) Error() string {
	return fmt.Sprintf("%s:%s", e.Type, e.Err)
}

type ErrType int8

const (
	MaximalIteration ErrType = iota
	InternalErr
	NotValidValue
	Recovery
)

func (et ErrType) String() string {
	switch et {
	case MaximalIteration:
		return "max iteration"
	case InternalErr:
		return "internal error"
	case NotValidValue:
		return "not valid value"
	case Recovery:
		return "recovery"
	}
	return "undefined"
}

// Find
// In mathematics, the bisection method is a root-finding method that applies
// to any continuous functions for which one knows two values with opposite
// signs.
// The method consists of repeatedly bisecting the interval defined
// by these values and then selecting the subinterval in which the function
// changes sign, and therefore must contain a root.
//
// Documentation: https://en.wikipedia.org/wiki/Bisection_method
//
//	Input data:
//		f    - function of variable X for root-finding
//		minX - minimal X
//		maxX - maximal X
//	Output data:
//		root - root of function
//		err  - error if some is not ok
//
// Notes:
//   - Concurrency acceptable
//   - Panic-free function
//
// Last operation of finding is run function.
func Find[F64 ~float64, F64R ~float64](f func(F64) (F64R, error), minX, maxX F64) (root F64, err error) {
	// recovering
	defer func() {
		if r := recover(); r != nil {
			err = ErrorFind{
				Type: Recovery,
				Err:  fmt.Errorf("%#v", r),
			}
		}
	}()
	// replace borders
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	// preparing variables
	var (
		xLeft, xRigth = minX, maxX
		middle        = func() F64 {
			return xLeft + (xRigth-xLeft)/2.0
		}
		xRoot            = middle()
		yLeft, errLeft   = f(xLeft)
		yRoot, errRoot   = f(xRoot)
		yRigth, errRigth = f(xRigth)

		prec    = Precision
		maxIter = MaxIteration
	)
	// another algo
	// just for information
	//
	// middle = func() float64 {
	// 	// xLeft , yLeft  //
	// 	// xRigth, yRigth //
	// 	var (
	// 		dY = yRigth - yLeft
	// 		dX = xRigth - xLeft
	// 	)
	// 	if 1e-6 < math.Abs(dX) {
	// 		a := dY / dX
	// 		b := yRigth - a*xRigth
	// 		x := -b / a
	// 		if xLeft < x && x < xRigth {
	// 			return x
	// 		}
	// 	}
	// 	return xLeft + (xRigth-xLeft)/2.0
	// }
	// check errors
	for _, errLocal := range []error{errLeft, errRigth, errRoot} {
		if errLocal != nil {
			err = errLocal
			return
		}
	}

	if math.Abs(float64(yLeft)) < prec {
		// find the solution
		root = xLeft
		_, err = f(F64(root))
		return
	}
	if math.Abs(float64(yRigth)) < prec {
		// find the solution
		root = xRigth
		_, err = f(F64(root))
		return
	}

	// iterations
	for iter := 0; ; iter++ {
		// check max iteration
		if iter >= maxIter {
			err = ErrorFind{
				Type: MaximalIteration,
				Err:  fmt.Errorf("Too many iterations: %d", iter),
			}
			return
		}
		if xLeft == 0 {
			if math.Abs(float64(yRoot)) < prec && math.Abs(float64(xRigth-xLeft)) < prec {
				break // find the solution
			}
		} else {
			if math.Abs(float64(yRoot)) < prec && math.Abs(float64((xRigth-xLeft)/xLeft)) < prec {
				break // find the solution
			}
		}
		if math.Signbit(float64(yLeft)) != math.Signbit(float64(yRoot)) {
			xRigth, yRigth = xRoot, yRigth
		} else if math.Signbit(float64(yRoot)) != math.Signbit(float64(yRigth)) {
			xLeft, yLeft = xRoot, yRoot
		} else {
			err = ErrorFind{
				Type: InternalErr,
				Err: fmt.Errorf("No root: [%.3e, %.3e, %.3e]",
					yLeft, yRoot, yRigth),
			}
			return
		}
		// preparing next middle point
		xRoot = middle()
		if yRoot, errRoot = f(xRoot); errRoot != nil {
			err = ErrorFind{
				Type: InternalErr,
				Err:  errRoot,
			}
			return
		}
		if math.IsNaN(float64(xRoot)) {
			err = ErrorFind{
				Type: NotValidValue,
				Err:  fmt.Errorf("xRoot is NaN"),
			}
			return
		}
		if math.IsNaN(float64(yRoot)) {
			err = ErrorFind{
				Type: NotValidValue,
				Err:  fmt.Errorf("yRoot is NaN"),
			}
			return
		}
		if math.IsInf(float64(xRoot), 0) {
			err = ErrorFind{
				Type: NotValidValue,
				Err:  fmt.Errorf("xRoot is Inf"),
			}
			return
		}
		if math.IsInf(float64(yRoot), 0) {
			err = ErrorFind{
				Type: NotValidValue,
				Err:  fmt.Errorf("yRoot is Inf"),
			}
			return
		}
	}
	root = xRoot
	_, err = f(F64(root))
	return
}
