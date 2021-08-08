package main

import (
	"fmt"
)

func main()  {
	var input1 = [][]int{{1,0,0,0},{0,0,0,0},{0,0,2,-1}}

	result1 := uniquePathsIII(input1)
	if result1 != 2 {
		fmt.Printf("Wrong result for input1. expected 2 but got %d \n", result1)
	} else {
		fmt.Printf("Result1: %d\n", result1)
	}

	var input2 = [][]int{{1,0,0,0},{0,0,0,0},{0,0,0,2}}

	result2 := uniquePathsIII(input2)
	if result2 != 4 {
		fmt.Printf("Wrong result for input2. expected 4 but got %d \n", result2)
	} else {
		fmt.Printf("Result2: %d\n", result2)
	}

	var input3 = [][]int{{0,1},{2,0}}
	result3 := uniquePathsIII(input3)
	if result3 != 0 {
		fmt.Printf("Wrong result for input3. expected 0 but got %d \n", result3)
	} else {
		fmt.Printf("Result3: %d\n", result3)
	}
}

func uniquePathsIII(grid [][]int) int {
	ys := len(grid)
	xs := len(grid[0])

	var start [2]int
	var available int

	for i:=0;i < ys; i++ {
		for j:=0;j < xs; j++ {
			if grid[i][j] == 0 {
				available++
			} else if grid[i][j] == 1 {
				start = [2]int{i, j}
			}
		}
	}

	var result int64

	next(start, grid, ys, xs, make([][2]int, 0, available), available, &result)

	return int(result)
}

func next(current [2]int, grid [][]int, ys int, xs int, used [][2]int, counter int, result *int64) {
	//go to the right!
	if current[1] < xs-1 && grid[current[0]][current[1]+1] != -1 && !isInList([2]int{current[0], current[1]+1}, used) {
		if grid[current[0]][current[1]+1] == 2 && counter == 0 {
			//finish
			*result++
			//atomic.AddInt64(result, 1)
		} else if grid[current[0]][current[1]+1] == 0 {
			crused := used[:]
			crused = append(crused, current)
			next([2]int{current[0], current[1]+1}, grid, ys, xs, crused, counter-1, result)
		}
	}

	//go to the left
	if current[1] > 0 && grid[current[0]][current[1]-1] != -1 && !isInList([2]int{current[0], current[1]-1}, used) {
		if grid[current[0]][current[1]-1] == 2 && counter == 0 {
			//finish
			*result++
			//atomic.AddInt64(result, 1)
		} else if grid[current[0]][current[1]-1] == 0 {
			clused := used[:]
			clused = append(clused, current)
			next([2]int{current[0], current[1]-1}, grid, ys, xs, clused, counter-1, result)
		}
	}

	//go to the top
	if current[0] < ys-1 && grid[current[0]+1][current[1]] != -1 && !isInList([2]int{current[0]+1, current[1]}, used) {
		if grid[current[0]+1][current[1]] == 2 && counter == 0 {
			//finish
			*result++
			//atomic.AddInt64(result, 1)
		} else if grid[current[0]+1][current[1]] == 0 {
			ctused := used[:]
			ctused = append(ctused, current)
			next([2]int{current[0]+1, current[1]}, grid, ys, xs, ctused, counter-1, result)
		}
	}

	//go to the bottom
	if current[0] > 0 && grid[current[0]-1][current[1]] != -1 && !isInList([2]int{current[0]-1, current[1]}, used) {
		if grid[current[0]-1][current[1]] == 2 && counter == 0 {
			//finish
			*result++
			//atomic.AddInt64(result, 1)
		} else if grid[current[0]-1][current[1]] == 0 {
			cbused := used[:]
			cbused = append(cbused, current)
			next([2]int{current[0]-1, current[1]}, grid, ys, xs, cbused, counter-1, result)
		}
	}
}

func isInList(point [2]int, used [][2]int) bool {
 	for _, v := range used {
 		if point == v {
 			return true
		}
	}

	return false
}
