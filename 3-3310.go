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
	const TAM int = len(adj) 
	visited := make([]bool, TAM)
	res := []int{}

	for i := 0; i < TAM; i++ {
		if visited[i] == false {
			dfsRec(adj, visited, i, &res)
		}
	}

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
	const EX_1 int[] = {
		{ 1, 2}, 
		{0, 3 }, 
		{2, 0 }, 
		{  5, 4} 
	}
	
	res := dfs()
}
