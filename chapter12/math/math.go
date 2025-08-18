//	math/math.go
package math

import "math"

/*
	calculates average
	To generate `go doc` run in parent folder:		
		go doc packagesapp/myutils Average
*/
func Average(xs []float64) float64 {
	
	if len(xs) == 0 { 
		return 0 
	}

	total := float64(0)
	
	for _, x := range xs {
		total += x 
	}
	
	return total / float64(len(xs))
}

/*
	calculates min
	To generate `go doc` run in parent folder:		
		go doc packagesapp/myutils Average
*/
func Min(xs []float64) float64 {

	if len(xs) == 0 { 
		return 0 
	}

	min := math.Inf(1)

	for _, x := range xs {
		if min > float64(x) {
			min = float64(x)
		}
	}
	
	return min
}

/*
	calculates max
	To generate `go doc` run in parent folder:		
		go doc packagesapp/myutils Average
*/
func Max(xs []float64) float64 {

	if len(xs) == 0 { 
		return 0 
	}

	max := math.Inf(-1)

	for _, x := range xs {
		if max < float64(x) {
			max = float64(x)
		}
	}
	
	return max
}