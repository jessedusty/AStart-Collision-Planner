package grid

// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package grid implements D* Lite grid-based pathfinding.

import (
	"ee631_midterm/dstarlite"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Coord represents an single grid coordinate pair [x, y]

type Coord []int

func neighborPerms(c Coord) (res []Coord) {
	for i := 0; i < len(c); i++ {
		if c[i] == 0 {
			curr1 := c.Clone()
			curr1[i] += 1
			curr2 := c.Clone()
			curr2[i] -= 1
			res = append(res, curr1, curr2)
		}
	}
	return
}

func dedupCoord(l []Coord) []Coord {
	m := make(map[string]Coord)
	for _, n := range l {
		m[n.Hash()] = n
	}

	var retVal []Coord
	for _, v := range m {
		retVal = append(retVal, v)
	}
	return retVal
}

func _neighborFinder(l []Coord, n int) []Coord {
	if n <= 0 {
		return []Coord{}
	}

	var retVal []Coord
	for _, c := range l {
		current := neighborPerms(c)
		retVal = append(retVal, current...)
		retVal = append(retVal, _neighborFinder(current, n-1)...)
	}

	return dedupCoord(retVal)
}

func neighborFinder(dimensions int) []Coord {
	return _neighborFinder([]Coord{make([]int, dimensions)}, dimensions)
}

func (a Coord) Clone() Coord {
	res := make(Coord, len(a))
	copy(res, a)
	return res
}

// Implements dstarlite.State interface.
func (a Coord) Equals(other dstarlite.State) bool {
	b := other.(Coord)
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Dist returns manhattan distance between two coordinates
// returns -1 if the Coordinates have different dimensions
func (a Coord) ManhattanDist(b Coord) float64 {
	if len(a) != len(b) {
		return -1
	}

	sum := float64(0)
	for i := 0; i < len(a); i++ {
		sum += math.Abs(float64(b[i] - a[i]))
	}

	return sum
}

// Dist returns manhattan distance between two coordinates
// returns -1 if the Coordinates have different dimensions
func (a Coord) Dist(b Coord) float64 {
	if len(a) != len(b) {
		return -1
	}

	sum := float64(0)
	for i := 0; i < len(a); i++ {
		sum += math.Pow(float64(b[i]-a[i]), 2)
	}

	return math.Sqrt(sum)
}

func (a Coord) Hash() string {
	valuesText := make([]string, len(a))
	for i, v := range a {
		valuesText[i] = strconv.Itoa(v)
	}
	return strings.Join(valuesText, ",")
}

func (a Coord) add(b Coord) Coord {
	if len(a) != len(b) {
		panic("Coordinates are different lengths, can't add!")
	}
	retVal := make(Coord, len(a))
	for i, _ := range b {
		retVal[i] = a[i] + b[i]
	}
	return retVal
}

func (a Coord) insideExtents(extents []int) bool {
	for i, _ := range extents {
		if a[i] > extents[i] {
			return false
		}
	}
	return true
}

// Data represents an actual grid's data.
type Data struct {
	dsl          *dstarlite.Planner
	coords       map[string]float64
	extents      []int
	start, goal  Coord
	neighborMask []Coord
	inverted     bool
}

func (d *Data) GetMapData() *map[string]float64 {
	return &d.coords
}

// Start returns the start coordinate, as it is currently.
func (d *Data) Start() Coord {
	return d.dsl.Start().(Coord)
}

// Goal returns the goal coordinate, as it is currently.
func (d *Data) Goal() Coord {
	return d.dsl.Goal().(Coord)
}

// Size returns the width and height of the grid.
func (d *Data) Size() []int {
	return d.extents
}

// Implements dstarlite.Data interface.
func (d *Data) Cost(aa, bb dstarlite.State) float64 {
	a := aa.(Coord)
	b := bb.(Coord)
	// Greater than 2 to handle diagonal
	if a.ManhattanDist(b) > 2 {
		return math.Inf(1)
	} else {
		costA, ok := d.coords[a.Hash()]
		if !ok && d.inverted {
			costA = 1
		}
		if costA < 0 {
			return math.Inf(1)
		}

		costB := d.coords[b.Hash()]
		if !ok && d.inverted {
			costB = 1
		}
		if costB < 0 {
			return math.Inf(1)
		}

		cost := ((costA + costB) / 2.0) + 1.0
		if cost <= 0 {
			panic("cost <= 0; should not happen!")
		}
		return cost
	}
}

// Implements dstarlite.Data interface.
func (d *Data) Dist(aa, bb dstarlite.State) float64 {
	a := aa.(Coord)
	b := bb.(Coord)
	return a.Dist(b)
}

// Returns an slice of neighbors to the current grid data cell excluding
// neighbor cells whom have an negative cost value.
func (d *Data) neighbors(ss dstarlite.State) (sl []dstarlite.State) {
	s := ss.(Coord)

	next := func(c1 Coord) {
		_, ok := d.Get(c1)
		if ok || d.inverted {
			sl = append(sl, c1)
		}
	}

	for _, c := range d.neighborMask {
		next(s.add(c))
	}

	return sl
}

// Implements dstarlite.Data interface.
func (d *Data) Succ(s dstarlite.State) []dstarlite.State {
	return d.neighbors(s)
}

// Implements dstarlite.Data interface.
func (d *Data) Pred(s dstarlite.State) []dstarlite.State {
	return d.neighbors(s)
}

// Set changes the cost of traversal to the given coordinate on the grid to the
// specified value.
//
// If the coordinate is outside of the grid's dimensions, an panic will occur.
func (d *Data) Set(pos Coord, value float64) {
	if !pos.insideExtents(d.extents) {
		panic(fmt.Sprintf("Set(): Coordinate %v outside of grid's dimensions %v", pos, d.extents))
	}

	if value == d.coords[pos.Hash()] {
		// Cost of traversal has not changed; simply do nothing.
		return
	}

	// Find all affected neighbor cells
	preds := d.Pred(pos)
	succs := d.Succ(pos)

	// Keep track of old costs
	predVals := make(map[string]float64)
	for _, sPrime := range preds {
		predVals[sPrime.(Coord).Hash()] = d.Cost(sPrime, pos)
	}

	succVals := make(map[string]float64)
	for _, sPrime := range succs {
		succVals[sPrime.(Coord).Hash()] = d.Cost(pos, sPrime)
	}

	// Change stored cost
	d.coords[pos.Hash()] = value

	// Flag changed to new costs
	for _, sPrime := range preds {
		d.dsl.FlagChanged(sPrime, pos, predVals[sPrime.(Coord).Hash()], d.Cost(sPrime, pos))
	}

	for _, sPrime := range succs {
		d.dsl.FlagChanged(pos, sPrime, succVals[sPrime.(Coord).Hash()], d.Cost(pos, sPrime))
	}
}

// Get returns the attached interface value from the given coordinate on the
// grid.
//
// If the coordinate is outside of the grid's dimensions, ok will equal false.
func (d *Data) Get(pos Coord) (value float64, ok bool) {
	if !pos.insideExtents(d.extents) {
		ok = false
		return
	}

	value, ok = d.coords[pos.Hash()]
	return
}

// UpdateStart updates the starting position. This can be used to move a long
// the path efficiently.
func (d *Data) UpdateStart(start Coord) {
	d.dsl.UpdateStart(start)
}

// Plan recomputes the lowest cost path through the map, taking into account
// changes in start location and edge costs.
//
// If no path is found, nil is returned.
func (d *Data) Plan() (path []Coord) {
	sPath := d.dsl.Plan()

	path = make([]Coord, len(sPath))
	for i, s := range sPath {
		path[i] = s.(Coord)
	}

	return
}

func allTrue(s []bool) bool {
	for _, v := range s {
		if v == false {
			return false
		}
	}
	return true
}

// Fill fills grid with value from pos, increments any dimension set to 1 in the mask
func (d *Data) Fill(value float64, pos Coord, mask Coord) {
	log.Printf("Fill(%v, %v, %v)", value, pos, mask)
	if !pos.insideExtents(d.extents) {
		panic(fmt.Sprintf("Set(): Coordinate %v outside of grid's dimensions %v", pos, d.extents))
	}

	finished := make([]bool, len(mask))
	for i, shouldFill := range mask {
		if shouldFill == 0 {
			finished[i] = true
		}
	}

	for !allTrue(finished) {
		d.Set(pos, value)
		for dim, shouldFill := range mask {
			if shouldFill != 0 && !finished[dim] {
				pos[dim] += 1
				if pos[dim] >= d.extents[dim] {
					finished[dim] = true
				}
			}
		}
	}
	log.Printf("Fill Finish")
}

// New returns an new grid data structure given an width and height where each
// cell indicates the cost of traversal.
//
// If the width or height are <= 0; an panic will occur.
//
// If either start or goal coordinates are outside the coordinates of the grid,
// an panic will occur.
func NewGrid(extents []int, start, goal Coord, undefinedGood bool) *Data {
	for _, v := range extents {
		if v < 0 {
			panic("NewGrid(): Cannot create grid with negative extents size")
		}
	}

	if !start.insideExtents(extents) {
		panic("NewGrid(): Start coordinate outside of grid's dimensions!")
	}
	if !goal.insideExtents(extents) {
		panic("NewGrid(): Goal coordinate outside of grid's dimensions!")
	}

	d := new(Data)
	d.inverted = undefinedGood
	d.extents = extents
	d.start = start
	d.goal = goal
	d.coords = make(map[string]float64)
	d.dsl = dstarlite.New(d, start, goal)
	d.neighborMask = neighborFinder(len(start))
	return d
}
