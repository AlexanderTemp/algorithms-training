# LeetCode Solutions (Go) 💻

This repository contains my LeetCode practice solutions using Go.

Each file includes:

- First approach to understand the problem
- Optimized solution based on research and LeetCode hints
- Comments explaining the reasoning and improvements

`Goal: improve problem solving skills and recognize algorithm patterns 🚀`

---

## Índice

- [Cheat Sheet](#cheat-sheet)
  - [Slices & init](#slices--init)
  - [Memoization](#memoization)
  - [Patrones DP](#patrones-dp)
- [Trucazos](#trucazos)
- [Resumen Go](#resumen-go)
  - [Tipos](#tipos)
  - [If / For](#if--for)
  - [Arrays & Slices](#arrays--slices)
  - [Strings](#strings)
  - [Funciones](#funciones)
  - [Map](#map)
  - [Grafo](#grafo)
  - [Extras](#extras)

## Cheat Sheet

### Slices & init

```go
dp := make([]int, n+1)     // 1D
dp := make([][]int, n)     // 2D
for i := range dp {
	dp[i] = make([]int, m)
}

for i := range dp {         // inicializar en -1
	dp[i] = -1
}
```

### Memoization

```go
if dp[state] != -1 {
	return dp[state]
}

memo := make(map[int]int)
if value, ok := memo[key]; ok {
	return value
}
```

### Patrones DP

`Top-Down 1D`

```go
func solve(state int, dp []int) int {
	if ... {                          // 1. caso base
		return ...
	}
	if dp[state] != -1 {              // 2. ya calculado
		return dp[state]
	}

	option1 := solve(..., dp)         // 3. transiciones
	option2 := solve(..., dp)

	dp[state] = max(option1, option2) // 4. guardar
	return dp[state]                  // 5. retornar
}
```

`Top-Down 2D`

```go
func solve(i, j int, dp [][]int) int {
	if ... {
		return ...
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}

	option1 := solve(i+1, j, dp)
	option2 := solve(i, j+1, dp)

	dp[i][j] = max(option1, option2)
	return dp[i][j]
}
```

`Bottom-Up`

```go
dp := make([]int, n+1)
dp[0] = ...

for i := 1; i <= n; i++ {
	dp[i] = ...
}

return dp[n]
```

> DP = Estado + Caso Base + Transición + Memoria

| Paso         | Pregunta                             |
| ------------ | ------------------------------------- |
| Estado       | Qué representa `dp[state]`            |
| Dependencia  | De qué estados anteriores depende     |
| Transición   | Cómo combino esos estados             |
| Memoria      | Se guarda el resultado o no           |

## Trucazos

| Truco | Código |
| --- | --- |
| If declara variable | `if value, ok := memo[key]; ok { return value }` |
| If con inicializador | `if x := len(nums); x > 10 { ... }` |
| Ignorar valor | `value, _ := strconv.Atoi("123")` |
| Asignación múltiple | `a, b = b, a` |
| Múltiples retornos | `func divide(a, b int) (int, error)` → `result, err := divide(10, 2)` |
| map con `ok` | `value, ok := memo[key]` |
| nil check | `if nums == nil { ... }` |
| switch sin `break` | cada `case` corta solo, sin fallthrough |
| Función anónima | `double := func(x int) int { return x * 2 }` |
| Puntero | `p := &x` dirección · `*p` valor · `*p = 20` modifica `x` |

## Resumen Go

### Tipos

```go
var x int = 10 // estricto desde que nace
x := 10        // inferencia y declaración

int int64 float64 string bool
```

### If / For

```go
if x > 10 { ... } else if x == 10 { ... } else { ... }

for i := 0; i < n; i++ { ... }
for _, x := range nums { ... } // recorrer slice
for i, x := range nums { ... } // índice + valor
for n > 0 { ... }              // while
```

### Arrays & Slices

`Uso para Stack, Queue`

```go
var a [5]int
a := [5]int{1, 2, 3, 4, 5}

nums := []int{}
nums = append(nums, 5)
nums[1:4] // subslice

stack := []int{}
stack = append(stack, x)     // push
x = stack[len(stack)-1]      // top
stack = stack[:len(stack)-1] // pop

queue := []int{}
queue = append(queue, x) // agregar
x = queue[0]
queue = queue[1:] // sacar
```

### Strings

```go
s := "hello"
len(s)
s[0]           // devuelve byte
r := []rune(s) // runas (unicode)

chars := []byte(s)
chars[0] = 'H'
s = string(chars)
```

### Funciones

```go
func add(a int, b int) int { ... }
func add(a, b int) int { ... } // forma corta
```

### Map

```go
m := make(map[int]int)
m[5] = 100    // insertar
x := m[5]     // obtener
x, ok := m[5] // comprobar existencia
delete(m, 5)
```

`Map con clave compuesta (estados DP)`

```go
memo := make(map[string]int)
key := fmt.Sprintf("%d,%d", i, j)

if value, ok := memo[key]; ok {
	return value
}
```

### Grafo

```go
adj := make([][]int, n)
adj[u] = append(adj[u], v) // arista dirigida

adj[u] = append(adj[u], v) // grafo no dirigido
adj[v] = append(adj[v], u)

for _, next := range adj[node] { ... }
```

### Extras

```go
const INF = int(1e9)            // problemas de mínimos
const MOD int64 = 1_000_000_007

import "math"
x := math.MaxInt
dp[i] = (dp[i] + dp[i-1]) % MOD

grid := [][]int{
	{1, 2, 3},
	{4, 5, 6},
}
n, m := len(grid), len(grid[0])

import "sort"
sort.Ints(nums)
sort.Strings(words)
sort.Slice(nums, func(i, j int) bool {
	return nums[i] > nums[j]
})

b := make([]int, len(a)) // copiar slice
copy(b, a)

reader := bufio.NewReader(os.Stdin) // input rápido
fmt.Fscan(reader, &n, &m)
fmt.Printf("%d\n", ans)
```
