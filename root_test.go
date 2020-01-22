package root_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/Konstantin8105/root"
)

func Test(t *testing.T) {
	tcs := []struct {
		f          func(x float64) float64
		Xmin, Xmax float64
	}{
		{
			func(x float64) float64 {
				return (3.8-3.0*math.Sin(math.Sqrt(x)))/0.35 - x
			},
			2.0,
			3.0,
		},
		{
			func(x float64) float64 {
				return 1.0/(3.0+math.Sin(3.6*x)) - x
			},
			0.0,
			0.85,
		},
		{
			func(x float64) float64 {
				return math.Cos(math.Sqrt(1.0-0.3*x*x*x)) - x
			},
			0,
			1,
		},
		{
			func(x float64) float64 {
				return math.Sin(math.Sqrt(1-0.4*x*x)) - x
			},
			0,
			1,
		},
		{
			func(x float64) float64 {
				return 0.25*x*x*x - x - 1.2502
			},
			2,
			3,
		},
		{
			func(x float64) float64 {
				return 0.1*x*x - x*math.Log(x)
			},
			1,
			2,
		},
		{
			func(x float64) float64 {
				return 3*x - 4*math.Log(x) - 5
			},
			2,
			4,
		},
		{
			func(x float64) float64 {
				return math.Exp(x) - math.Exp(-x) - 2
			},
			0,
			1,
		},
		{
			func(x float64) float64 {
				return x + math.Sqrt(x) + math.Pow(x, 1.0/3.0) - 2.5
			},
			0.4,
			1,
		},
		{
			func(x float64) float64 {
				return math.Tan(x) - math.Pow(math.Tan(x), 3.0)/3 + math.Pow(math.Tan(x), 5)/5.0 - 1./3.
			},
			0,
			0.8,
		},
		{
			func(x float64) float64 {
				return math.Cos(2.0/x) - 2.0*math.Sin(1./x) + 1./x
			},
			1,
			2,
		},
		{
			func(x float64) float64 {
				return math.Sin(math.Log(x)) - math.Cos(math.Log(x)) + 2.0*math.Log(x)
			},
			1,
			3,
		},
		{
			func(x float64) float64 {
				return math.Log(x) - x + 1.8
			},
			2,
			3,
		},
		{
			func(x float64) float64 {
				return 0.4 + math.Atan(math.Sqrt(x)) - x
			},
			1,
			2,
		},
		{
			func(x float64) float64 {
				return x*math.Tan(x) - 1/3.0
			},
			0.2,
			1,
		},
		{
			func(x float64) float64 {
				return math.Tan(0.55*x+0.1) - x*x
			},
			0,
			1,
		},
		{
			func(x float64) float64 {
				return 2.0 - math.Sin(1./x) - x
			},
			1.2,
			2,
		},
		{
			func(x float64) float64 {
				return 1.0 + math.Sin(x) - math.Log(1+x) - x
			},
			0,
			1.5,
		},
		{
			func(x float64) float64 {
				return math.Cos(math.Pow(x, 0.52)+2) + x
			},
			0.4,
			1,
		},
		{
			func(x float64) float64 {
				return math.Sqrt(math.Log(1+x)+3) - x
			},
			2,
			3,
		},
		{
			func(x float64) float64 {
				return math.Exp(x) + math.Log(x) - 10*x
			},
			3,
			4,
		},
		{
			func(x float64) float64 {
				return 3*x - 14 + math.Exp(x) - math.Exp(-x)
			},
			1,
			3,
		},
		{
			func(x float64) float64 {
				return 2*math.Pow(math.Log(x), 2) + 6*math.Log(x) - 5
			},
			1,
			3,
		},
		{
			func(x float64) float64 {
				return 2*x*math.Sin(x) - math.Cos(x)
			},
			0.4,
			1,
		},
		{
			// some strange function
			func(x float64) float64 {
				return PartLine(x, []xys{
					{0, 4}, {0.3, 1}, {1.3, 0.5}, {1.35, -0.5}, {2.0, -1.5},
				})
			},
			0,
			2,
		},
		{
			// some strange function
			func(x float64) float64 {
				return PartLine(x, []xys{
					{0, 3}, {0.25, 4}, {0.5, 0.1}, {2.0, -0.1},
				})
			},
			0,
			2,
		},
		{
			// some strange function
			func(x float64) float64 {
				return PartLine(x, []xys{
					{0, 3}, {0.1, 4}, {0.2, 0.01}, {1.6, -0.01}, {1.9, -4.0}, {2.0, -3.0},
				})
			},
			0,
			2,
		},
		{
			// some strange function
			func(x float64) float64 {
				return PartLine(x, []xys{
					{0, 3}, {0.1, 0.001}, {1.8, -0.001}, {2.0, -0.1},
				})
			},
			0,
			2,
		},
	}

	var counter int64

	for i := range tcs {
		t.Run(fmt.Sprintf("Case%3d", i), func(t *testing.T) {
			tempFunc := func(x float64) (float64, error) {
				counter++
				return tcs[i].f(x), nil
			}
			rootX, err := root.Find(tempFunc, tcs[i].Xmin, tcs[i].Xmax)
			if err != nil {
				t.Error(err)
			}
			if rootX < tcs[i].Xmin || tcs[i].Xmax < rootX {
				t.Errorf("not valid root")
			}
			if math.Abs(tcs[i].f(rootX)) > root.Precision {
				t.Errorf("not valid precision")
			}
		})
	}

	averageCalls := float64(counter) / float64(len(tcs))
	t.Logf("Average amount of calls: %.2f", averageCalls)
}

type xys struct {
	x, y float64
}

func PartLine(x float64, xy []xys) float64 {
	last := len(xy) - 1
	if x < xy[0].x {
		return line(x, xy[0].x, xy[0].y, xy[1].x, xy[1].y)
	}
	if x > xy[last].x {
		return line(x, xy[last-1].x, xy[last-1].y, xy[last].x, xy[last].y)
	}
	for i := 1; i < len(xy); i++ {
		if xy[i-1].x <= x && x <= xy[i].x {
			return line(x, xy[i-1].x, xy[i-1].y, xy[i].x, xy[i].y)
		}
	}
	return -42.0
}

func line(x, x0, y0, x1, y1 float64) float64 {
	a := (y1 - y0) / (x1 - x0)
	b := y0 - a*x0
	return a*x + b
}

func TestPanic(t *testing.T) {
	p := func(float64) (float64, error) {
		panic("PANIC")
	}
	_, err := root.Find(p, 0, 1)
	t.Logf("%v", err)
	if err == nil {
		t.Fatalf("Cannot panic finding")
	}
}

func TestChangeMinMax(t *testing.T) {
	nr := func(x float64) (float64, error) {
		return 2*x + 1, nil
	}
	_, err := root.Find(nr, 10, -10)
	t.Logf("%v", err)
	if err != nil {
		t.Fatalf("Finding not valid root")
	}
}

func TestNoRoot(t *testing.T) {
	nr := func(x float64) (float64, error) {
		return 2*x + 5, nil
	}
	_, err := root.Find(nr, 0, 1)
	t.Logf("%v", err)
	if err == nil {
		t.Fatalf("Finding not valid root")
	}
}

func TestNoSomeRoot(t *testing.T) {
	{
		// left
		nr := func(x float64) (float64, error) {
			if x < 0.5 {
				return -1, fmt.Errorf("left checking")
			}
			return 2*x + 5, nil
		}
		_, err := root.Find(nr, 0, 1)
		t.Logf("%v", err)
		if err == nil {
			t.Fatalf("Finding not valid root: left")
		}
	}
	{
		// center
		nr := func(x float64) (float64, error) {
			if x == 0.5 {
				return -1, fmt.Errorf("center checking")
			}
			return 2*x + 5, nil
		}
		_, err := root.Find(nr, 0, 1)
		t.Logf("%v", err)
		if err == nil {
			t.Fatalf("Finding not valid root: center")
		}
	}
	{
		// rigth
		nr := func(x float64) (float64, error) {
			if x > 0.5 {
				return -1, fmt.Errorf("rigth checking")
			}
			return 2*x + 5, nil
		}
		_, err := root.Find(nr, 0, 1)
		t.Logf("%v", err)
		if err == nil {
			t.Fatalf("Finding not valid root: rigth")
		}
	}
}
