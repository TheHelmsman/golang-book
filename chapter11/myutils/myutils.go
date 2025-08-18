//	myutils/myutils.go
package myutils

import "math"

/*
	calculates average
	To generate `go doc` run in parent folder:		
		go doc packagesapp/myutils Average
*/
func Average(xs []float64) float64 {
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
	max := math.Inf(-1)

	for _, x := range xs {
		if max < float64(x) {
			max = float64(x)
		}
	}
	
	return max
}