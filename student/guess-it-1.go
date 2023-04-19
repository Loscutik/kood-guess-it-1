package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"guessit/statistics"
)

const (
	INIT_SIZE = 10 // this number of numbers will add to the set for analize without any conditions
	// these constants need for adjusting a guess interval
	PART_FOR_1_GUESS   = 20
	DIFF_AVR_MED       = 20
	PART_DEV_FOR_GUESS = 1
	DEVS_FOR_ADD       = 10
)

func main() {
	ordinary := make([]int, 0, INIT_SIZE) // slice for storing numbers which don't lie too far from the Average. Don't want to analyse esceptions
	var (
		n int 
		median, sDev float64
		err             error
	)
	scanner := bufio.NewScanner(os.Stdin)

	// add the first numbers to the ordinary slice and do guesses without using  StandardDeviation
	for i := 0; i < INIT_SIZE; i++ {
		if !scanner.Scan() {
			return
		}

		// finish the programme if nothing enters
		txt := scanner.Text()
		if txt == "" {
			return
		}

		n, err = strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}

		ordinary = append(ordinary, n)

		average := statistics.AverageInt(ordinary)
		guess(average, average/PART_FOR_1_GUESS)

	}

	ordinary = check(ordinary) // delete too big and too small numbers
	sDev = statistics.StandardDeviationInt(ordinary)
	median = statistics.MedianInt(ordinary)

	// add all next numbers to the ordinary slice and do guesses using StandardDeviation
	for scanner.Scan() {
		// finish the programme if nothing enters
		txt := scanner.Text()
		if txt == "" {
			return
		}
		n, err = strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}

		if float64(n) > median-sDev*DEVS_FOR_ADD && float64(n) < median+sDev*DEVS_FOR_ADD {
			ordinary = append(ordinary, n)
		}

		sDev = statistics.StandardDeviationInt(ordinary)
		median = statistics.MedianInt(ordinary)
		guess(median, sDev/PART_DEV_FOR_GUESS)
	}
}

/*
print the range which the next number may be in
*/
func guess(median float64, deviation float64) {
	fmt.Printf("%v %v\n", math.Round(median - deviation), math.Round(median + deviation))
}

/*
checks if there are too big or too small numbers in the given slice and returns the slice without these ones
*/
func check(ordinary []int) []int {
	average := statistics.AverageInt(ordinary)
	median := statistics.MedianInt(ordinary)

	if math.Abs(average-median) > median/PART_FOR_1_GUESS {
		max := maxIndex(ordinary)
		min := minIndex(ordinary)

		// delete the furthest from the medium number from given slice
		i := maxIndex([]int{int(median) - ordinary[min], ordinary[max] - int(median)})
		if i == 0 { // min is the furthest item
			ordinary = remove(ordinary, min)
			// repeat checking after deleting
			ordinary = check(ordinary)
		} else { // max is the furthest item
			ordinary = remove(ordinary, max)
			// repeat checking after deleting
			ordinary = check(ordinary)
		}
	}
	return ordinary
}

/*
removes the i-th item from given slice, the order of items in the slice is not kept
*/
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

/*
returns an index of the max item of the given slice
*/
func maxIndex(nmbs []int) int {
	l := len(nmbs)
	if l == 0 {
		return -1
	}

	max := 0
	for i := 1; i < l; i++ {
		if nmbs[max] < nmbs[i] {
			max = i
		}
	}
	return max
}

/*
returns an index of the min item of the given slice
*/
func minIndex(nmbs []int) int {
	l := len(nmbs)
	if l == 0 {
		return -1
	}

	min := 0
	for i := 1; i < l; i++ {
		if nmbs[min] > nmbs[i] {
			min = i
		}
	}
	return min
}

