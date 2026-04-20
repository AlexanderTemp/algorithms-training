// 1. Two sum
package main

import "fmt"

// PEOR handler en tiempo (primer intento = hash map): Sin éxito se revisan todos los casos hasta que se detiene sin embargo cumplió en LeetCode
func twoSum(nums []int, target int) []int {
	tam := len(nums)
	indices := []int{}

	possibleTargets := make(map[int]int)

	for i := 0; i < tam && len(indices) == 0; i++ {
		value := nums[i]

		if value <= target { // Esta comparación es mala pueden existir negativos
			for key, pastValue := range possibleTargets {
				if value+pastValue == target {
					indices = append(indices, key, i)
					break
				}

			}

			possibleTargets[i] = value
		}

	}

	fmt.Printf("Para la combinacion %v con target %d se encontró los indices %v\n", nums, target, indices)

	return indices
}

// MEJOR handler (optimizado) Uso Leet sugiere: Hash Table, Two Pointers
//
//  1. Se está iterando todo el map c/loop, mata ventaja de hash map
//
//  2. Mejor buscar el complemento guardar en el hash complementos no suma
//     Antes: value + pastValue == target ---- Ahora = complement := target - value
//
//  3. Se busca en el hashmap por el nuevo valor el complemento solo (2-pointers), guardar alrevés antes {indice: valor} ahora {valor, indice} (MÁS RÁPIDO)
//
//  4. No importa posiciones diferentes por la naturaleza del problema (solapamiento de valores en índice sin importancia)
//
//     A DISPOSICIÓN EL VALOR MÁS IMPORTANTE EN EL INDICE DEL HASHMAP para no iterarlo
//     EVITAR USO DE VARIABLE RES return directo
//     FIJA EL TAMAÑO MÁXIMO DEL HASH SI SE CONOCE EVITA REHASHING INTERNO (VITAL OJO)

func twoSumOptimized(nums []int, target int) []int {
	possibleComplements := make(map[int]int, len(nums))

	for i, value := range nums {
		indiceAntiguo, ok := possibleComplements[target-value]

		if ok {
			fmt.Printf("Para la combinacion %v con target %d se encontró los indices %v\n", nums, target, []int{i, indiceAntiguo})
			return []int{i, indiceAntiguo}
		}

		possibleComplements[value] = i

	}

	return []int{}
}

func main() {
	caso1Nums, caso1Target := []int{2, 7, 11, 15}, 9
	caso2Nums, caso2Target := []int{3, 2, 4}, 6
	caso3Nums, caso3Target := []int{3, 3}, 6

	twoSumOptimized(caso1Nums, caso1Target)
	twoSumOptimized(caso2Nums, caso2Target)
	twoSumOptimized(caso3Nums, caso3Target)

}
