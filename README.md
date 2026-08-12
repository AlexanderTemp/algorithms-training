# LeetCode Solutions (Go) 💻

This repository contains my LeetCode practice solutions using Go.

Each file is named after the corresponding problem and includes:

- First approach to understand the problem
- Optimized solution based on research and LeetCode hints
- Comments explaining the reasoning and improvements

`Goal: improve problem solving skills and recognize algorithm patterns 🚀`

--- 

## Go cheat-sheet
```go

// Slice 1D
dp := make([]int, n+1)

// Slice 2D
dp := make([][]int, n)

for i := range dp {
	dp[i] = make([]int, m)
}

// Inicializar
for i := range dp {
	dp[i] = -1
}

// Loop
for i := 0; i < n; i++ { ... }

// Range
for _, x := range nums { ... }

// Min
dp[i] = min(a, b)

// Max
dp[i] = max(a, b)

// Memoization
if dp[state] != -1 {
	return dp[state]
}

// Recursión
solve(nextState)

// Append
nums = append(nums, x)

// Map key -> value
memo := make(map[int]int)

if value, ok := memo[key]; ok {
	return value
}

```

`Patrón DP - Top Down + Memoization`

```go
func solve(state int, dp []int) int {

	// 1. Caso base
	if ... {
		return ...
	}

	// 2. Ya calculado
	if dp[state] != -1 {
		return dp[state]
	}

	// 3. Transiciones
	option1 := solve(..., dp)
	option2 := solve(..., dp)

	// 4. Guardar resultado
	dp[state] = max(option1, option2)

	// 5. Retornar
	return dp[state]
}
```

`Patrón DP 2D`

```go
func solve(i, j int, dp [][]int) int {

	// Caso base
	if ... {
		return ...
	}

	// Memo
	if dp[i][j] != -1 {
		return dp[i][j]
	}

	// Transiciones
	option1 := solve(i+1, j, dp)
	option2 := solve(i, j+1, dp)

	// Guardar
	dp[i][j] = max(option1, option2)

	return dp[i][j]
}
```

`DP Bottom-Up`

```go
dp := make([]int, n+1)

dp[0] = ...

for i := 1; i <= n; i++ {
	dp[i] = ...
}

return dp[n]
```

> DP = Estado + Caso Base + Transición + Memoria

- 1. Lo que representa dp[state]
- 3. De qué estados anteriores depende
- 4. Cómo combino esos estados
- 5. Se guarda el resultado o no


## Go Resumen

### Vars

```go
var x int = 10 // Estricto desde que nace
x :=10 // Deduce por valor almacenado

int int64 float64 string bool
```

### If 
```go
if x > 10 { ... }
else if x == 10 { ... }
else { ... }
```


### For

```go
for i := 0, i < n; i++ { ... }

for _, x :range nums { ... } // Recorrer slice 
for i, x :range nums { ... } // Indice + valor
for n > 0 { ... }            // While
```

### Arrays

```go
var a [5]int
a := [5]int{1, 2, 3, 4, 5}
```

### Slices -> Uso para Stack, Queu

```go
nums := []int{}
nums := []int{1, 2, 3, 4}

nums = append(nums, 5)
len(nums)
nums[1:4]   // subslice

stack := []int{}

stack = append(stack, x) // Push
x := stack[len(stack) - 1] // Pop
stack = stack[:len(stack)-1]

x := stack[len(stack)-1] // Top


queue := []int{}
queue = append(queue, x) // Agregar
x := queue[0]
queue  = queue[1:] // Sacar 

```

### Crear DP

`1D`

```go
n := 10 // 1D
dp := make([]int, n+1)

dp := make([]int, n+1)
for i := range dp {
  dp[i] = -1
}
```

`2D`

```go
n := 5 
m := 10

dp := make([][]int, n)

for i := range dp {
  dp[i] = make([]int, m)
}

dp := make([][]int, n)

for i := range dp {
  dp[i] = make([]int, m)
  
  for j := range dp[i] {
    dp[i][j] = -1
  }
}
```



### Strings

```go

s := "hello"
len(s)

s[0] // devlueve byte
r := []rune(s)

```

`string -> byte`

```go
s := "hello"

chars := []byte(s)
chars[0] = 'H'
s = string(chars)

```

### Funcs

```go
func add(a int, b int)  int { }

func add(a, b int) int  { } // más corto

func calc(a, b int) (int, int) { } // varios valores

sum, diff := calc(10,5)
```


### Recursivas

`DP Top-Down`

```go
func fib(n int, dp []int) int {
  if n <= 1 {
    return n
  }
    
  if dp[n] != -1 {
    return dp[n]
  }

  dp[n] = fib(n-1, dp) + fib(n-2, dp)

  return dp[n]
}
```

`DP Bottom-Up`

```go
func fib(n int) int {
  if n <= 1 {
    return n
  }

  dp := make([]int, n+1)
  
  dp[0] = 0 
  dp[1] = 1

  for i := 2, i <= n; i++ {
    dp[i] = dp[i - 1] + dp [i - 2]
  }

  return dp[n]
}
```

### Infinito y constantes

```go
const INF = int(1e9) // Para problemas de mínimos

import "math"
x := math.MaxInt

dp[i] = math.MaxInt + 1 // NO Hacer comparar antes si se excede el valor
```

```go

const MOD int64 = 1_000_000_007

dp[i] = (dp[i] + dp[i-1] % MOD)

```

### Matriz de entrada

```go
grid := [][]int {
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

n := len(grid)
m := len(grid[0])

```

### Map 

```go 
m := make(map[int]int)

m[5] = 100 // Insertar
x := m[5]  // Obtener

x, ok := m[5] // Comprobar existencia
if ok { ... }

delete(m, 5)
```

`Map para estados DP`

```go
memo := make(map[string]int)
key := fmt.Sprintf("%d,%d", i , j)

if value, ok := memo[key]; ok {
  return value
}
```


### Grafo

```go
adj := make([][]int, n)

adj[u] = append(adj[u], v) // Agregar arista

// Grafo no dirigido
adj[u] = append(adj[u], v)
adj[v] = append(adj[v], u)

for _, next := range adf[node] { ... }
```

### Tricks

```go
a, b, = b, a // Intercambio

import "sort" // Ordenar ASC
sort.Ints(nums)
sort.String(words)

sort.Slice(nums, func(i, j int) bool {
  return nums[i] > nums[j]
})

b := make([]int, len(a)) // Copiar slices
copy(b, a)


reader := bufio.NewReader(os.Stdin) // leer input rápido 

fmt.Fscan(reader, &n)
fmt.Fscan(reader, &n, &m)

fmt.Println(ans)
fmt.Printf("%d\n", ans)

```
