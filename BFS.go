package main

import (
	"container/list"
	"fmt"
)

// Prints Breadth First Search from given vertices
func bfs(adjList *[]list.List, s int) {
	Q := list.New()
	size := numberOfVertices(adjList)
	visited := make([]bool, size)

	visited[s] = true
	Q.PushBack(s)

	for Q.Len() != 0 {
		s = Q.Front().Value.(int)
		tmp := Q.Front()
		Q.Remove(tmp)
		fmt.Println(s)
		for v := (*adjList)[s].Front(); v != nil; v = v.Next() {
			if !visited[v.Value.(int)] {
				visited[v.Value.(int)] = true
				Q.PushBack(v.Value.(int))
			}
		}
	}
}
