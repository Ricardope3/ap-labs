// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Point struct
type Point struct{ x, y float64 }

// Distance traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// Distance same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		} else {
		}
	}
	return sum
}

func getRandomPoint() Point {

	min := -100
	max := 100
	time.Sleep(10)
	x := rand.Intn(max-min) + min
	y := rand.Intn(max-min) + min
	return Point{float64(x), float64(y)}
}

func printPoint(p Point) {

	fmt.Printf("  - (  %f,   %f)\n", p.x, p.y)
}

func printPath(p Path) {
	fmt.Println("- Figure's vertices")
	for i := 0; i < len(p); i++ {
		printPoint(p[i])
	}
}

//!-path

func onSegment(p, q, r Point) bool {
	if q.x < math.Max(p.x, r.x) && q.x > math.Min(p.x, r.x) && q.y < math.Max(p.y, r.y) && q.y > math.Min(p.y, r.y) {
		return true
	}
	return false
}
func orientation(p, q, r Point) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)

	if val == 0 {
		return 0
	}
	if val > 0 {
		return 1
	}
	return 2
}

func intersect(p1, q1, p2, q2 Point) bool {

	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)
	if o1 != o2 && o3 != o4 {
		return true

	}
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}
	return false
}

func candidatePointDoesIntersect(path *Path, candidate Point, i int) bool {
	p1 := (*path)[i-2]
	p2 := (*path)[i-1]
	if intersect(p1, p2, p2, candidate) {
		return true
	}
	return false
}
func addPointToPath(path *Path, size int) error {
	candidate := getRandomPoint()
	if candidatePointDoesIntersect(path, candidate, size) {
		return errors.New("")
	}
	(*path)[size] = candidate
	return nil
}
func geometry(numSides int) {
	fmt.Printf("- Generating a [%d] sides figure\n", numSides)
	path := make(Path, numSides)
	path[0] = getRandomPoint()
	path[1] = getRandomPoint()
	size := 2
	for size < len(path) {
		err := addPointToPath(&path, size)
		if err == nil {
			size++
		}
	}
	printPath(path)
	fmt.Printf("Perimetro: %f\n", path.Distance())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numSides, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Argument bust me an int")
		os.Exit(1)
	}
	if numSides < 3 {
		fmt.Println("At least three points are needed to create a polygon")
		os.Exit(1)
	}
	geometry(numSides)

}
