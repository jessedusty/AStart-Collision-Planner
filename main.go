package main

import "ee631_midterm/dstar_grid"

func main() {
	g := dstar_grid.New(100, 100, []int{10,10}, []int{10,10})
	g.Set()

}
