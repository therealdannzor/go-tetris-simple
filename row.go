package main

import "fmt"

type Row struct {
	fields  [10]bool
	counter int
}

func (r Row) mark(i int) {
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
