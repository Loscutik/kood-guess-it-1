/*
package imlements mathematition statistics functions
*/
package statistics

import (
	"math"
	"sort"
)
/*
	calculates  the arithmetic average. 
*/
func Average(nmbs []float64) float64 {
	if len(nmbs) == 0 {
		return 0
	}
	s := 0.0
	for _, n := range nmbs {
		s += n
	}
	return s / float64(len(nmbs))
}

/*
	calculates the median of a sets of numbers. 
*/
func Median(nmbs []float64) float64 {
	sort.Float64s(nmbs)
	l := len(nmbs)
	if l%2 == 0 {
		return (nmbs[l/2-1]+nmbs[l/2]) / 2.0
	} else {
		return nmbs[l/2]
	}
}

/*
	calculates the variance of a sets of numbers. 
*/
func Variance(nmbs []float64) float64 {
	if len(nmbs) == 0 {
		return 0
	}
	s := 0.0
	av := Average(nmbs)
	for _, n := range nmbs {
		s += (n - av) * (n - av)
	}
	return s / float64(len(nmbs))
}

/*
	calculates the standard deviation of a sets of numbers. 
*/
func StandardDeviation(nmbs []float64) float64 {
	return float64(math.Round(math.Sqrt(float64(Variance(nmbs)))))
}

/*
	calculates the average and standard deviation of a sets of numbers. 
*/
func AverageAndStandardDeviation(nmbs []float64)( average float64, deviation float64) {
	l:=len(nmbs)
	if l == 0 {
		return 0,0
	}
	average=Average(nmbs)
	s := 0.0
	for _, n := range nmbs {
		s += (n - average) * (n - average)
	}
	return average, math.Sqrt(s/float64(l))
}

/*
	calculates  the arithmetic average of set of integer.
*/
func AverageInt(nmbs []int) float64 {
	if len(nmbs) == 0 {
		return 0
	}
	s := 0
	for _, n := range nmbs {
		s += n
	}
	return float64(s) / float64(len(nmbs))
}

/*
	calculates the median of a sets of integer numbers. 
*/
func MedianInt(nmbs []int) float64 {
	sort.Ints(nmbs)
	l := len(nmbs)
	if l%2 == 0 {
		return float64(nmbs[l/2-1]+nmbs[l/2]) / 2.0
	} else {
		return float64(nmbs[l/2]) 
	}
}

/*
calculates the variance of a sets of integer numbers. 
*/
func VarianceInt(nmbs []int) float64 {
	if len(nmbs) == 0 {
		return 0
	}
	s := 0.0
	av := AverageInt(nmbs)
	for _, n := range nmbs {
		s += (float64(n) - av) * (float64(n) - av)
	}
	return float64(s) / float64(len(nmbs))
}

/*
calculates the standard deviation of a sets of integer numbers. 
*/
func StandardDeviationInt(nmbs []int) float64 {
	return math.Sqrt(VarianceInt(nmbs))
}


/*
	calculates the average and standard deviation of a sets of numbers. 
*/

func AverageAndStandardDeviationInt(nmbs []int)( average float64, deviation float64) {
	l:=len(nmbs)
	if l == 0 {
		return 0,0
	}
	average=AverageInt(nmbs)
	s := 0.0
	for _, n := range nmbs {
		s += (float64(n) - average) * (float64(n) - average)
	}
	return average, math.Sqrt(s/float64(l))
}