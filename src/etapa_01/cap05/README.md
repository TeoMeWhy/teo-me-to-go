# Capítulo 05

## Exemplos

### Criação de `array`.

Array com tamanho 5 do tipo inteiro, alterando o valor do índice 4 (quinto elemento)
```go
{{#include exemplo01/main.go}}
```

### Soma de uma `array` do tipo `float64`.

Inicialização do `array` com 5 elementos (todos com valor 0), onde cada um é alterado, e depois utilizamos um laço `for` para percorrer os índices realizando a soma de cada elemento.

```go
{{#include exemplo02/main.go}}
```

### Função `len()` para tamanho de `array`.
```go
{{#include exemplo03/main.go}}
```

### Percorrendo elementos do `array` sem utilizar índices.
```go
{{#include exemplo04/main.go}}
```

### Iniciando um array já com valores
```go
{{#include exemplo05/main.go}}
```

### Exemplos de formas diferentes para criar um `array`.
```go
{{#include exemplo06/main.go}}
```

### Criação de fatias (`slices`).
```go
{{#include exemplo07/main.go}}
```

### Copiando conteúdo de uma fatia para uma fatia menor.
```go
{{#include exemplo08/main.go}}
```

### Definição de mapas (chave/valor).
```go
{{#include exemplo09/main.go}}
```

### Acessando elementos de uma mapa.

Em Go, caso a chave que tentamos acessa não exista, nos é retornado um valor vazio correspondente ao tipo do valor. Isto é, `""` para strings e `0` para int.

Além disso, e possível checar se o elemento realmente não existe com base no segundo valor que nos é retornado, sendo este um booleando (`true` ou `false`).
```go
{{#include exemplo10/main.go}}
```

### Definindo chaves/valor no momento de criação dos mapas.
```go
{{#include exemplo11/main.go}}
```

### Mapas que tem mapas como valores associados à suas chaves (mapas de mapas).
```go
{{#include exemplo12/main.go}}
```

---

## Exercícios

### Exercício 01
```go
{{#include ex01/main.go}}
```

### Exercício 02
```go
{{#include ex02/main.go}}
```

### Exercício 03
```go
{{#include ex03/main.go}}
```

### Exercício 04
```go
{{#include ex04/main.go}}
```
