package main

type bubble []int

func sort(a bubble) bubble {
	for swapped := true; swapped; {
		swapped = false
		for index := 1; index < len(a); index++ {
			if a[index-1] > a[index] {
				a[index-1], a[index] = a[index], a[index-1]
				swapped = true
			}
		}
	}
	return a
}
