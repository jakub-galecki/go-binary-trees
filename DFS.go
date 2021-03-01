package main

import (
	"container/list"
	"fmt"
)

func dfsVisit(adjListArray *[]list.List, u int, visited *[]bool) {
	(*visited)[u] = true
	fmt.Printf("%d\n", u)
	for v := (*adjListArray)[u].Front(); v != nil; v = v.Next() {
		if !(*visited)[v.Value.(int)] {
			dfsVisit(adjListArray, v.Value.(int), visited)
		}
	}
}

func DFS(adjListArray *[]list.List, v ...int) {
	size := numberOfVertices(adjListArray)
	visited := make([]bool, size)

	if len(v) != 0 {
		if len(v) > 1 {
			panic("You can pass only one verice")
		} else {
			dfsVisit(adjListArray, v[0], &visited)
		}
	} else {
		for i := 0; i < len(*adjListArray); i++ {
			tmp := (*adjListArray)[i].Front().Value.(int)
			if !visited[tmp] {
				dfsVisit(adjListArray, tmp, &visited)
			}
		}
	}

}
