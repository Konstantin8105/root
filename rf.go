package rf

import (
	"fmt"
	"math"
)

// Constants
const (
	// Precision of rott-finding
	Precision float64 = 1e-6

	// MaxIteration is max allowable amount of iteration.
	// Typically for precition=1e-6 need 20 iterations.
	MaxIteration int = 500
)

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
//	* Concurrency acceptable
//	* Panic-free function
//
// Last operation of finding is run function.
//
func Find(f func(float64) (float64, error), minX, maxX float64) (root float64, err error) {
	// error handling
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cannot find [%.5e,%.5e]: %v", minX, maxX, err)
		}
	}()
	// recovering
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%#v", r)
		}
	}()
	// replace borders
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	// preparing variables
	var (
		left, rigth    = minX, maxX
		fl, fr, frigth float64

		prec    = Precision
		maxIter = MaxIteration
	)
	// iterations
	for iter := 0; ; iter++ {
		if iter >= maxIter {
			return -1, fmt.Errorf("Too many iterations: %d", iter)
		}
		root = left + (rigth-left)/2.0 // middle coordinate X
		if fl, err = f(left); err != nil {
			return
		}
		if frigth, err = f(rigth); err != nil {
			return
		}
		if fr, err = f(root); err != nil {
			return
		}

		if math.Abs(fr) < prec && math.Abs(rigth-left) < prec {
			break
		}
		if math.Signbit(fl) != math.Signbit(fr) {
			rigth = root
		} else if math.Signbit(fr) != math.Signbit(frigth) {
			left = root
		} else {
			err = fmt.Errorf("No root in x = [%.2f,%.2f,%.2f]. y=[%.2f,%.2f,%.2f]",
				left, root, rigth, fl, fr, frigth)
			return
		}
	}
	_, err = f(root)
	return
}
