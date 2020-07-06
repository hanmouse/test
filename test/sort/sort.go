package main

import (
	"fmt"
	"sort"
)

type gear struct {
	name  string
	price float32
}

var myGears []gear = []gear{
	{"heresy", 99.99},
	{"asgard3", 199.99},
	{"e30", 120},
}

// for "less" function
type by func(g1, g2 *gear) bool

func (by by) Sort(gears []gear) {
	sorter := &gearSorter{
		gears: gears,
		by:    by,
	}
	sort.Sort(sorter)
}

type gearSorter struct {
	gears []gear
	by    func(g1, g2 *gear) bool
}

func (r *gearSorter) Len() int {
	return len(r.gears)
}

func (r *gearSorter) Swap(i, j int) {
	r.gears[i], r.gears[j] = r.gears[j], r.gears[i]
}

func (r *gearSorter) Less(i, j int) bool {
	return r.by(&r.gears[i], &r.gears[j])
}

func main() {

	name := func(g1, g2 *gear) bool {
		return g1.name < g2.name
	}

	price := func(g1, g2 *gear) bool {
		return g1.price < g2.price
	}

	fmt.Println("Before sorting:", myGears)

	by(name).Sort(myGears)
	fmt.Println("After sorting by name:", myGears)

	by(price).Sort(myGears)
	fmt.Println("After sorting by price:", myGears)
}
