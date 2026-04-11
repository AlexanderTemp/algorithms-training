// 3740. Minimum Distance Between Three Equal Elements I

package main

import "fmt"
import "math"

// PEOR handler (primer intento = fuerza bruta): Uso de matemática para simplificar operación dado  eq1 (p1 < p2 < p3)
func minimumDistance(nums []int) int {
	tam := len(nums)
	abs := -1

	for i := 0; i < tam-2; i++ {
		pivot := nums[i]
		positions := []int{i}

		for j := i + 1; j < tam; j++ {
			if nums[j] == pivot {
				positions = append(positions, j)
			}
		}

		if len(positions) < 3 {
			continue
		}

		for j := range len(positions) - 2 {
			for k := j + 1; k < len(positions)-1; k++ {
				for l := k + 1; l < len(positions); l++ {
					p1, p2, p3 := positions[j], positions[k], positions[l]
					calc := int(math.Abs(float64(p1-p2)) + math.Abs(float64(p2-p3)) + math.Abs(float64(p3-p1)))

					if abs == -1 {
						abs = calc
					}
					if calc < abs {
						abs = calc
					}
				}
			}
		}

	}
	return abs
}

// MEJOR handler (optimizado 1) Uso Leet sugiere: Hash map, two pointers, greedy
//  1. Se sugiere que | p1 - p2 | + | p2 - p3 | + | p3 - p1 | se puede igualar entre cada elemento
//     | p1 - p2 | = p2 - p1 -> recordando eq1 se hace lo mismo con todas y se procede a la suma para igualar con la eq original
//     eqOriginal = (p2 - p1) + (p3 - p2) + (p3 - p1)
//     eqO = 2p3 - 2p1 + (-p2 + p2) = 2(p3 - p1) -> Cómo hablamos de distancias multiplicarlo x2 no es necesario no cambiara la normal de la distancia
//     eqO = p3 - p1 -> eq2
//  2. Con eq2 también se conoce que comparar todas las combinaciones no es necesario solo comparar p1, p2, p3 cercanos
//     Un solo for loop que valide 3 posiciones continuas basta
//  3. Se usa hash map clave valor por ser O(1)
//     m := make(map[int]int) // make(map[TipoClave]TipoValor)
//     En este caso sería m := make(map[int][]int)
//     m[Clave] = nuevoValor (inserta o actualiza depende del valor)
//     v := m[Clave] // Asignación de valor
func minimumDistanceBest(nums []int) int {
	m := make(map[int][]int)
	best := -1

	for i, v := range nums {
		// aplicación de hash map
		m[v] = append(m[v], i)
		tamTemp := len(m[v])

		if tamTemp > 2 {
			// 2 pointers equivale a revisar ventanas de tamaño 3 siempre
			p3, p1 := m[v][tamTemp-1], m[v][tamTemp-3]
			calc := p3 - p1
			// fmt.Printf("Para combinación de %d, las posiciones %v se hallo %d\n", v, m[v], calc)

			// Greedy decisión local óptima
			if best == -1 || best > calc {
				best = calc
			}
		}

	}

	// Se multiplica * 2 por formula eq0
	if best != -1 {
		return best * 2
	}

	return best
}

func main() {
	caso1 := []int{1, 2, 1, 1, 3}
	caso2 := []int{1, 1, 2, 3, 2, 1, 2}
	caso3 := []int{1}

	fmt.Println(minimumDistanceBest(caso1))
	fmt.Println(minimumDistanceBest(caso2))
	fmt.Println(minimumDistanceBest(caso3))

}
