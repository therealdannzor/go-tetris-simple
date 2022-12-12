package main

import "fmt"

type Row struct {
	fields  [10]bool
	counter int
}

// position is 0..9
type position int

func (r Row) mark(i position) {
	if r.fields[i] == false {
		r.fields[i] = true
		r.counter++
		if r.counter == 10 {
			r.clean()
		}
	} else {
		fmt.Println("unexpected error: cannot mark same point twice")
	}
}

func (r Row) clean() {
	var l [10]bool
	r.fields = l
	r.counter = 0
}

func (r Row) isFree(i position) bool {
	return !r.fields[i]
}
