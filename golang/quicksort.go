package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func partition(array []int32, p uint, q uint, pivotLocation uint) uint {

	pivot := array[pivotLocation]

	array[pivotLocation], array[q] = array[q], array[pivotLocation]

	i, j := p, p
	d := (q - p)
	l := (d >> 4) << 4

	for j = p; j < p+l; j += 16 {
		
		if array[j] <= pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
		if array[j+1] <= pivot {
			array[i], array[j+1] = array[j+1], array[i]
			i++
		}
		if array[j+2] <= pivot {
			array[i], array[j+2] = array[j+2], array[i]
			i++
		}
		if array[j+3] <= pivot {
			array[i], array[j+3] = array[j+3], array[i]
			i++
		}
		if array[j+4] <= pivot {
			array[i], array[j+4] = array[j+4], array[i]
			i++
		}
		if array[j+5] <= pivot {
			array[i], array[j+5] = array[j+5], array[i]
			i++
		}
		if array[j+6] <= pivot {
			array[i], array[j+6] = array[j+6], array[i]
			i++
		}
		if array[j+7] <= pivot {
			array[i], array[j+7] = array[j+7], array[i]
			i++
		}
		if array[j+8] <= pivot {
			array[i], array[j+8] = array[j+8], array[i]
			i++
		}
		if array[j+9] <= pivot {
			array[i], array[j+9] = array[j+9], array[i]
			i++
		}
		if array[j+10] <= pivot {
			array[i], array[j+10] = array[j+10], array[i]
			i++
		}
		if array[j+11] <= pivot {
			array[i], array[j+11] = array[j+11], array[i]
			i++
		}
		if array[j+12] <= pivot {
			array[i], array[j+12] = array[j+12], array[i]
			i++
		}
		if array[j+13] <= pivot {
			array[i], array[j+13] = array[j+13], array[i]
			i++
		}
		if array[j+14] <= pivot {
			array[i], array[j+14] = array[j+14], array[i]
			i++
		}
		if array[j+15] <= pivot {
			array[i], array[j+15] = array[j+15], array[i]
			i++
		}
		
	}

	for j = p + l; j < q; j++ {
		if array[j] <= pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	
	array[q], array[i] = array[i], array[q]
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
