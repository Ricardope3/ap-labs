package main

import (
	"math/rand"
	"time"
	"fmt"
)

func printMatrix(matrix [][]int){
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func insertExit(floor [][]int, sideExits []int, side, indexExit, lenF int) bool {
	switch side {
	case 0:
		if (floor[0][indexExit] == 3 || floor[1][indexExit] == 1){
			return false
		}
		floor[0][indexExit] = 3
	case 1:
		if (floor[indexExit][lenF-1]==3 || floor[indexExit][lenF-2] == 1){
			return false
		}
		floor[indexExit][lenF-1] = 3
	case 2:
		if (floor[lenF-1][indexExit]==3 || floor[lenF-2][indexExit] == 1){
			return false
		}
		floor[lenF-1][indexExit] = 3
	case 3:
		if (floor[indexExit][0]==3 || floor[indexExit][1] == 1){
			return false
		}
		floor[indexExit][0] = 3
	default:
		sideExits[side]++
	}
	return true
}

func generateExits(floor [][]int){
	sideExits := make([]int, 4)
	lenF := len(floor)
	rand.Seed(time.Now().UnixNano())
	nExits := rand.Intn(5) + 1
	for i:=0;i<nExits;{
		side := rand.Intn(4)
		indexExit := rand.Intn(lenF)
		if (sideExits[side]>1){
			continue
		}
		valid := insertExit(floor, sideExits, side, indexExit, lenF)
		if(valid){
			i++
		}
	}
}

func main(){
	a := make([][]int, 12)
	for i := range a {
		a[i] = make([]int, 12)
	}
	lenA := len(a)
	for i := range a {
		a[0][i] = 1
		a[i][0] = 1
		a[i][lenA-1] = 1
		a[lenA-1][i] = 1
	}
	generateExits(a)
	printMatrix(a)
}