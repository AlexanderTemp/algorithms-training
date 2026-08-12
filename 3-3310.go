// 3310. Remove Methods From Project

package main

import "fmt"
import "math"

func dfsRec(adj [][]int, visited []bool, s int, res *[]int) {
	visited[s] = true
	*res = append(*res, s)

	for _, i := range adj[s] {
		if !visited[i] {
			dfsRec(adj, visited, i, res)
		}
	}
}

func dfs(adj [][]int) []int {
	visited := make([]bool, len(adj))
	res := []int{}
	dfsRec(adj, visited, 0, &res)
	return res
}

// PEOR handler (primer intento = fuerza bruta)
//

func minimumDistance(nums []int) int {

}

// MEJOR handler Uso Leet sugiere:
//

func minimumDistanceBest(nums []int) int {
}

func main3() {
}
