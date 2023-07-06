# Capítulo 06

## Exemplos

### Calculando média sem definir uma função para a média.
```go
{{#include exemplo01/main.go}}
```

### Definindo função para média.

Em Go, a palavra reservada `func` define funções. A seguir, adicionamos o nome da função e seus parâmetros (zero ou mais), bem como o tipo associado de cada parâmetro. Além disso, podemos retornar zero ou mais valores, especificando seus tipos.
```go
{{#include exemplo02/main.go}}
```

### Nomeando retorno

Podemos dizer no memento da assinatura da função, qual variável será retornada.
```go
{{#include exemplo03/main.go}}
```

### Funções variádicas.

Ao escrever a assinatura de uma função, seu último parâmetro pode ser variádico. Isso significa que podemos receber uma quantidade indeterminada de valores para este parâmetro.
```go
{{#include exemplo04/main.go}}
```

### Funções dentro de funções
```go
{{#include exemplo05/main.go}}
```

### Funções acessando variáveis de outra função
```go
{{#include exemplo06/main.go}}
```

### Clojure
```go
{{#include exemplo07/main.go}}
```

### Recursão
```go
{{#include exemplo08/main.go}}
```

### Defer

Maneira de executar uam instrução (normalmente uma função), ao final do escopo onde está sendo lançada. Isso garante a sua execução, sempre.
```go
{{#include exemplo09/main.go}}
```

### Combinando de `defer` + `panic`.

Isso é um anti-padrão de Go e deve ser evitado.
```go
{{#include exemplo10/main.go}}
```

### Função que não altera valor de `x`.

Para alterarmos o valor de `x`, precisamos acessar seu ponteiro, uma vez que o que é passado para a função `zero()` é uma cópia do valor de `x`, e não `x` propriamente.
```go
{{#include exemplo11/main.go}}
```

### Função com ponteiro

Utilizamos o `&` para acessa o endereço de memória e `*` para acessa o valor correspondente à este endereço.

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

### Exercício 05
```go
{{#include ex05/main.go}}
```

### Exercício 06
```go
{{#include ex06/main.go}}
```

### Exercício 07
```go
{{#include ex07/main.go}}
```

### Exercício 08
```go
{{#include ex08/main.go}}
```

### Exercício 09
```go
{{#include ex09/main.go}}
```

### Exercício 10
```go
{{#include ex10/main.go}}
```

### Exercício 11
```go
{{#include ex11/main.go}}
```