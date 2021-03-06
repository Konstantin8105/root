# root

root-finding algorithm

```
package root // import "."

const (
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
    Constants


func Find(f func(float64) (float64, error), minX, maxX float64) (root float64, err error)
    Find In mathematics, the bisection method is a root-finding method that
    applies to any continuous functions for which one knows two values with
    opposite signs. The method consists of repeatedly bisecting the interval
    defined by these values and then selecting the subinterval in which the
    function changes sign, and therefore must contain a root.

    Documentation: https://en.wikipedia.org/wiki/Bisection_method

        Input data:
        	f    - function of variable X for root-finding
        	minX - minimal X
        	maxX - maximal X
        Output data:
        	root - root of function
        	err  - error if some is not ok

    Notes:

        * Concurrency acceptable
        * Panic-free function

    Last operation of finding is run function.
```
