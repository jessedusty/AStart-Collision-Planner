package astar_planner

import "fmt"

func neighborPerms(c []int) (res [][]int) {
	for i := 0; i < len(c); i++ {
		if c[i] == 0 {
			curr1 := CloneCoord(c)
			curr1[i] += 1
			//curr2 := CloneCoord(c)
			//curr2[i] -= 1
			res = append(res, curr1)
		}
	}
	return
}

func dedupCoord(l [][]int) [][]int {
	m := make(map[string][]int)
	for _, n := range l {
		m[fmt.Sprint(n)] = n
	}

	var retVal [][]int
	for _, v := range m {
		retVal = append(retVal, v)
	}
	return retVal
}

func _neighborFinder(l [][]int, n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	var retVal [][]int
	for _, c := range l {
		current := neighborPerms(c)
		retVal = append(retVal, current...)
		retVal = append(retVal, _neighborFinder(current, n-1)...)
	}

	return dedupCoord(retVal)
}

func neighborFinder(dimensions int) [][]int {
	return _neighborFinder([][]int{make([]int, dimensions)}, dimensions)
}
