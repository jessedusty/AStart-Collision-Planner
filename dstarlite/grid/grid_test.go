package grid

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"testing"
)

func contains(s []Coord, e Coord) bool {
	for _, n := range s {
		if n.Equals(e) {
			return true
		}
	}
	return false
}

func convertToCoord(a []int) Coord {
	b := make(Coord, len(a))
	for i, n := range a {
		b[i] = n
	}
	return b
}

func convertToCoords(a [][]int) []Coord {
	retVal := make([]Coord, len(a))
	for i, n := range a {
		retVal[i] = convertToCoord(n)
	}
	return retVal
}

func assertContainsSame(t *testing.T, a []Coord, rhs [][]int) {
	b := convertToCoords(rhs)
	for _, n := range a {
		if !contains(b, n) {
			t.Fatalf("B missing %v that is in A", n)
			t.Fail()
		}
	}
	if len(a) != len(b) {
		for _, n := range b {
			if !contains(a, n) {
				t.Fatalf("A missing %v that is in B\n === A === \n%v\n\n=== B === \n %v \n", n, a, b)
				t.Fail()
			}
		}
	}
}

//
//func assertEqual(t *testing.T, a interface{}, b interface{}) {
//	if !reflect.DeepEqual(a, b) {
//		t.Fatalf("%s != %s", a, b)
//	}
//}

func TestConvertToCoords(t *testing.T) {
	a := [][]int{{1, 0}}
	_ = convertToCoords(a)
}

func TestDedupCoord(t *testing.T) {
	assertContainsSame(t, dedupCoord(convertToCoords([][]int{{1, 1}, {1, 1}, {0, 1}})), [][]int{{1, 1}, {0, 1}})
}

func TestNeighborPerms(t *testing.T) {
	assertContainsSame(t, neighborPerms([]int{0}), [][]int{{1}, {-1}})
	assertContainsSame(t, neighborPerms([]int{0, 0}), [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}})
	//assertContainsSame(t, neighborPerms([]int{0, 0}), [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
}

func TestPNeighborFinder(t *testing.T) {
	assertContainsSame(t, _neighborFinder(convertToCoords([][]int{{0}}), 1), [][]int{{1}, {-1}})
	assertContainsSame(t, _neighborFinder(convertToCoords([][]int{{0, 0}}), 2), [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}})
}

func TestNeighborFinder(t *testing.T) {
	assertContainsSame(t, neighborFinder(1), [][]int{{1}, {-1}})
	assertContainsSame(t, neighborFinder(2), [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}})
}

type Changeable interface {
	image.Image
	Set(x, y int, c color.Color)
}

func TestMazeThing(t *testing.T) {
	src, err := imaging.Open("30.png")

	if err != nil {
		t.Fatal(err, "Could not open image")
	}

	bounds := src.Bounds().Max

	// old starting coordinates
	// grid := NewGrid(bounds.X, bounds.Y, Coord{12, 14}, Coord{500, 500}, true)
	grid := NewGrid(bounds.X, bounds.Y, Coord{2, 2}, Coord{27, 27})
	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			r, _, _, _ := src.At(x, y).RGBA()
			if r > 128 {
				grid.Set(Coord{x, y}, 1)
			}
		}
	}

	coords := grid.Plan()
	if len(coords) < 10 {
		t.Fatal("No plan found")
	}

	assertContainsSame(t, coords, [][]int{{2, 2}, {3, 2}, {4, 2}, {5, 2}, {5, 3}, {6, 3}, {7, 3}, {7, 4}, {8, 4}, {8, 5}, {8, 6}, {9, 6}, {9, 7}, {10, 7}, {11, 7}, {11, 6}, {12, 6}, {12, 5}, {12, 4}, {12, 3}, {13, 3}, {14, 3}, {15, 3}, {15, 2}, {16, 2}, {17, 2}, {18, 2}, {18, 3}, {19, 3}, {19, 4}, {20, 4}, {20, 5}, {21, 5}, {21, 6}, {21, 7}, {21, 8}, {21, 9}, {21, 10}, {21, 11}, {22, 11}, {23, 11}, {24, 11}, {24, 12}, {24, 13}, {24, 14}, {23, 14}, {23, 15}, {23, 16}, {22, 16}, {22, 17}, {22, 18}, {21, 18}, {21, 19}, {20, 19}, {19, 19}, {19, 18}, {18, 18}, {18, 17}, {18, 16}, {17, 16}, {17, 15}, {17, 14}, {17, 13}, {18, 13}, {18, 12}, {18, 11}, {17, 11}, {17, 10}, {17, 9}, {16, 9}, {15, 9}, {14, 9}, {14, 10}, {13, 10}, {13, 11}, {12, 11}, {12, 12}, {11, 12}, {11, 13}, {11, 14}, {10, 14}, {9, 14}, {9, 13}, {9, 12}, {9, 11}, {9, 10}, {8, 10}, {7, 10}, {6, 10}, {6, 9}, {6, 8}, {5, 8}, {4, 8}, {3, 8}, {2, 8}, {2, 9}, {2, 10}, {3, 10}, {3, 11}, {4, 11}, {4, 12}, {4, 13}, {4, 14}, {3, 14}, {3, 15}, {2, 15}, {2, 16}, {2, 17}, {2, 18}, {2, 19}, {2, 20}, {2, 21}, {3, 21}, {3, 22}, {3, 23}, {4, 23}, {5, 23}, {5, 24}, {6, 24}, {7, 24}, {8, 24}, {8, 23}, {9, 23}, {10, 23}, {11, 23}, {12, 23}, {12, 22}, {12, 21}, {11, 21}, {11, 20}, {11, 19}, {12, 19}, {13, 19}, {14, 19}, {14, 20}, {14, 21}, {15, 21}, {16, 21}, {16, 22}, {16, 23}, {15, 23}, {15, 24}, {15, 25}, {16, 25}, {16, 26}, {17, 26}, {17, 27}, {18, 27}, {19, 27}, {19, 26}, {19, 25}, {19, 24}, {19, 23}, {20, 23}, {21, 23}, {21, 22}, {22, 22}, {23, 22}, {24, 22}, {24, 23}, {25, 23}, {25, 24}, {25, 25}, {26, 25}, {27, 25}, {27, 26}, {27, 27}})

	outImg := src.(Changeable)

	for _, c := range coords {
		outImg.Set(c[0], c[1], color.RGBA{0, 255, 0, 255})
	}

	fmt.Println(coords)

	err = imaging.Save(outImg, "out.png")
	if err != nil {
		log.Fatalln(err)
	}
}

func TestMazeThingInv(t *testing.T) {
	src, err := imaging.Open("30.png")

	if err != nil {
		t.Fatal(err, "Could not open image")
	}

	bounds := src.Bounds().Max

	grid := NewGrid(bounds.X, bounds.Y, Coord{2, 2}, Coord{27, 27})

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			r, _, _, _ := src.At(x, y).RGBA()
			if r < 128 {
				grid.Set(Coord{x, y}, -1)
			}
		}
	}

	coords := grid.Plan()

	if len(coords) < 10 {
		t.Fatal("No plan found")
	}

	assertContainsSame(t, coords, [][]int{{2, 2}, {3, 2}, {4, 2}, {5, 2}, {5, 3}, {6, 3}, {7, 3}, {7, 4}, {8, 4}, {8, 5}, {8, 6}, {9, 6}, {9, 7}, {10, 7}, {11, 7}, {11, 6}, {12, 6}, {12, 5}, {12, 4}, {12, 3}, {13, 3}, {14, 3}, {15, 3}, {15, 2}, {16, 2}, {17, 2}, {18, 2}, {18, 3}, {19, 3}, {19, 4}, {20, 4}, {20, 5}, {21, 5}, {21, 6}, {21, 7}, {21, 8}, {21, 9}, {21, 10}, {21, 11}, {22, 11}, {23, 11}, {24, 11}, {24, 12}, {24, 13}, {24, 14}, {23, 14}, {23, 15}, {23, 16}, {22, 16}, {22, 17}, {22, 18}, {21, 18}, {21, 19}, {20, 19}, {19, 19}, {19, 18}, {18, 18}, {18, 17}, {18, 16}, {17, 16}, {17, 15}, {17, 14}, {17, 13}, {18, 13}, {18, 12}, {18, 11}, {17, 11}, {17, 10}, {17, 9}, {16, 9}, {15, 9}, {14, 9}, {14, 10}, {13, 10}, {13, 11}, {12, 11}, {12, 12}, {11, 12}, {11, 13}, {11, 14}, {10, 14}, {9, 14}, {9, 13}, {9, 12}, {9, 11}, {9, 10}, {8, 10}, {7, 10}, {6, 10}, {6, 9}, {6, 8}, {5, 8}, {4, 8}, {3, 8}, {2, 8}, {2, 9}, {2, 10}, {3, 10}, {3, 11}, {4, 11}, {4, 12}, {4, 13}, {4, 14}, {3, 14}, {3, 15}, {2, 15}, {2, 16}, {2, 17}, {2, 18}, {2, 19}, {2, 20}, {2, 21}, {3, 21}, {3, 22}, {3, 23}, {4, 23}, {5, 23}, {5, 24}, {6, 24}, {7, 24}, {8, 24}, {8, 23}, {9, 23}, {10, 23}, {11, 23}, {12, 23}, {12, 22}, {12, 21}, {11, 21}, {11, 20}, {11, 19}, {12, 19}, {13, 19}, {14, 19}, {14, 20}, {14, 21}, {15, 21}, {16, 21}, {16, 22}, {16, 23}, {15, 23}, {15, 24}, {15, 25}, {16, 25}, {16, 26}, {17, 26}, {17, 27}, {18, 27}, {19, 27}, {19, 26}, {19, 25}, {19, 24}, {19, 23}, {20, 23}, {21, 23}, {21, 22}, {22, 22}, {23, 22}, {24, 22}, {24, 23}, {25, 23}, {25, 24}, {25, 25}, {26, 25}, {27, 25}, {27, 26}, {27, 27}})

	outImg := src.(Changeable)

	for _, c := range coords {
		outImg.Set(c[0], c[1], color.RGBA{0, 255, 0, 255})
	}

	fmt.Println(coords)

	err = imaging.Save(outImg, "out1.png")
	if err != nil {
		log.Fatalln(err)
	}
}
