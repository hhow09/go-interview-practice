package main

import (
	"sync"
)

// ConcurrentBFSQueries concurrently processes BFS queries on the provided graph.
// - graph: adjacency list, e.g., graph[u] = []int{v1, v2, ...}
// - queries: a list of starting nodes for BFS.
// - numWorkers: how many goroutines can process BFS queries simultaneously.
//
// Return a map from the query (starting node) to the BFS order as a slice of nodes.
// YOU MUST use concurrency (goroutines + channels) to pass the performance tests.
func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	// TODO: Implement concurrency-based BFS for multiple queries.
	// Return an empty map so the code compiles but fails tests if unchanged.
	if numWorkers == 0 {
		return map[int][]int{}
	}
	res := map[int][]int{}
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	jobs := make(chan int)
	for j := 0; j < numWorkers; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for node := range jobs {
				bfsRes := bfs(graph, node)
				mu.Lock()
				res[node] = bfsRes
				mu.Unlock()
			}
		}()
	}
	for _, startingNode := range queries {
		jobs <- startingNode
	}
	close(jobs)

	wg.Wait()
	return res
}

func bfs(graph map[int][]int, node int) []int {
	visited := make(map[int]bool)
	res := []int{}
	queue := []int{node}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		visited[node] = true
		res = append(res, node)
		for _, nb := range graph[node] {
			if !visited[nb] {
				queue = append(queue, nb)
			}
		}
	}
	return res
}

func main() {
	// You can insert optional local tests here if desired.

}
