package math

import (
	"fmt"
	"testing"
)

/*
	'Go test' recompiles each package along with any files with names matching
	the file pattern "*_test.go".
	
	These additional files can contain test functions, benchmark functions, fuzz
	tests and example functions. See 'go help testfunc' for more.
	Each listed package causes the execution of a separate test binary.

	more information:

		`go help test `

	to run:

		`go test`

	sample output:

	PASS
	ok      packagesapp/math        0.401s
*/

type testpair struct {
	values []float64
	average float64
}

var testsAvg = []testpair{
	{ []float64{1,2}, 1.5 },
	// { []float64{1,2}, 2 }, //	<- to make it fail
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
	{ []float64{}, 0},	// <- test empty list
}

var testsMin = []testpair{
	{ []float64{1,2}, 1 },
	// { []float64{1,2}, 2 }, //	<- to make it fail
	{ []float64{1,2,3,4,5,6}, 1 },
	{ []float64{-1,1}, -1 },
	{ []float64{}, 0},	// <- test empty list
}

var testsMax = []testpair{
	{ []float64{1,2}, 2 },
	// { []float64{1,2}, 1 }, //	<- to make it fail
	{ []float64{1,2,3,4,5,6}, 6 },
	{ []float64{-1,1}, 1 },
	{ []float64{}, 0},	// <- test empty list
}

func TestAverage(t *testing.T) {
	fmt.Println("TestAverage")
	for _, pair := range testsAvg {
		v := Average(pair.values)
		
		if v != pair.average {
			t.Error(
      	"For", pair.values,
      	"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	fmt.Println("TestMin")
	for _, pair := range testsMin {
		v := Min(pair.values)
		
		if v != pair.average {
			t.Error(
      	"For", pair.values,
      	"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	fmt.Println("TestMax")
	for _, pair := range testsMax {
		v := Max(pair.values)
		
		if v != pair.average {
			t.Error(
      	"For", pair.values,
      	"expected", pair.average,
				"got", v,
			)
		}
	}
}