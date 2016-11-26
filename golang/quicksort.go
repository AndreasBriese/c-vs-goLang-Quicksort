package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func swap(x *int32, y *int32) {
	x, y = y, x
}

func partition(array []int32, p uint, q uint, pivotLocation uint) uint {

	pivot := array[pivotLocation]
	swap(&array[pivotLocation], &array[q])
	i, j := p, p
	d := (q - p)
	if d&15 == 0 { // %16
		for j = p; j < q; j += 16 {
			if array[j] <= pivot {
				swap(&array[i], &array[j])
				i++
			}
			if array[j+1] <= pivot {
				swap(&array[i], &array[j+1])
				i++
			}
			if array[j+2] <= pivot {
				swap(&array[i], &array[j+2])
				i++
			}
			if array[j+3] <= pivot {
				swap(&array[i], &array[j+3])
				i++
			}
			if array[j+4] <= pivot {
				swap(&array[i], &array[j+4])
				i++
			}
			if array[j+5] <= pivot {
				swap(&array[i], &array[j+1])
				i++
			}
			if array[j+6] <= pivot {
				swap(&array[i], &array[j+2])
				i++
			}
			if array[j+7] <= pivot {
				swap(&array[i], &array[j+3])
				i++
			}
			if array[j+8] <= pivot {
				swap(&array[i], &array[j])
				i++
			}
			if array[j+9] <= pivot {
				swap(&array[i], &array[j+1])
				i++
			}
			if array[j+10] <= pivot {
				swap(&array[i], &array[j+2])
				i++
			}
			if array[j+11] <= pivot {
				swap(&array[i], &array[j+3])
				i++
			}
			if array[j+12] <= pivot {
				swap(&array[i], &array[j+4])
				i++
			}
			if array[j+13] <= pivot {
				swap(&array[i], &array[j+1])
				i++
			}
			if array[j+14] <= pivot {
				swap(&array[i], &array[j+2])
				i++
			}
			if array[j+15] <= pivot {
				swap(&array[i], &array[j+3])
				i++
			}
		}
		swap(&array[q], &array[i])
		return i
	}
	if d&7 == 0 { // %8
		for j = p; j < q; j += 8 {
			if array[j] <= pivot {
				swap(&array[i], &array[j])
				i++
			}
			if array[j+1] <= pivot {
				swap(&array[i], &array[j+1])
				i++
			}
			if array[j+2] <= pivot {
				swap(&array[i], &array[j+2])
				i++
			}
			if array[j+3] <= pivot {
				swap(&array[i], &array[j+3])
				i++
			}
			if array[j+4] <= pivot {
				swap(&array[i], &array[j+4])
				i++
			}
			if array[j+5] <= pivot {
				swap(&array[i], &array[j+1])
				i++
			}
			if array[j+6] <= pivot {
				swap(&array[i], &array[j+2])
				i++
			}
			if array[j+7] <= pivot {
				swap(&array[i], &array[j+3])
				i++
			}
		}
		swap(&array[q], &array[i])
		return i
	}
	if d&3 == 0 { // %4
		for j = p; j < q; j += 4 {
			if array[j] <= pivot {
				swap(&array[i], &array[j])
				i++
			}
			if array[j+1] <= pivot {
				swap(&array[i], &array[j+1])
				i++
			}
			if array[j+2] <= pivot {
				swap(&array[i], &array[j+2])
				i++
			}
			if array[j+3] <= pivot {
				swap(&array[i], &array[j+3])
				i++
			}
		}
		swap(&array[q], &array[i])
		return i
	}
	
	for j = p; j < q; j += 4 {
		if array[j] <= pivot {
			swap(&array[i], &array[j])
			i++
		}
	}

	swap(&array[q], &array[i])
	return i
}

func quicksort(array []int32, start uint, end uint) {
	if start < end {
		pivot := (end + start) / 2
		r := partition(array, start, end, pivot)
		if r > start {
			quicksort(array, start, r-1)
		}
		quicksort(array, r+1, end)
	}
}

func main() {

	runtime.GOMAXPROCS(2)
	runtime.LockOSThread()

	a1 := make([]int32, 100000000)

	for i := 0; i < 100000000; i++ {
		a1[i] = int32(rand.Int())
	}

	fmt.Printf("a1[0]: %d, a1[1]: %d, a1[2]: %d\n", a1[0], a1[1], a1[2])

	// now run the test
	t0 := time.Now()
	quicksort(a1, 0, 100000000-1)
	t1 := time.Since(t0)
	fmt.Printf("n of %v took %v to run.\n", len(a1), t1.Seconds())
	fmt.Printf("a1[0]: %d, a1[1]: %d, a1[2]: %d\n", a1[0], a1[1], a1[2])
	
}
