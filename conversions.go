package main

import (
	"container/list"
	"fmt"
)

func adjListToAdjMatrix(adjListArr *[]list.List) [][]int {
	arrSize := len(*adjListArr)
	adjMatrix := make([][]int, arrSize)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, arrSize)
	}
	for _, v := range *adjListArr {
		if tmp := v.Front(); tmp != nil {
			if row, ok := v.Front().Value.(int); ok {
				next := tmp.Next()
				for next != nil {
					if col, ok := next.Value.(int); ok {
						adjMatrix[row][col] = 1
					}
					next = next.Next()
				}
			}
		}
	}
	return adjMatrix
}

func adjMatrixToAdjList(matrix *[][]int) []list.List {
	size := len(*matrix)
	adjListArr := make([]list.List, size)
	for i, li := range *matrix {
		adjListArr[i].PushBack(i)
		for j, v := range li {
			if v == 1 {
				adjListArr[i].PushBack(j)
			}
		}
	}
	return adjListArr
}

func printAdjList(adjListArr *[]list.List) {
	for _, li := range *adjListArr {
		printList(li.Front())
		fmt.Println()
	}
}

func printList(e *list.Element) {
	if e != nil {
		fmt.Printf("%v ", e.Value)
		printList(e.Next())
	}
}

func printAdjMatrix(matrix *[][]int) {
	for _, v := range *matrix {
		for _, tmp := range v {
			fmt.Printf(" %d ", tmp)
		}
		fmt.Println()
	}
}
func numberOfVertices(adjListArray *[]list.List) int {
	counter := 0
	for _, v := range *adjListArray {
		for t := v.Front(); t != nil; t = t.Next() {
			counter++
		}
	}
	return counter
}
