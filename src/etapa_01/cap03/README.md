# Capítulo 03

## Exemplos

### Como criamos variáveis em Go:
```go
{{#include exemplo01/main.go}}
```

### Alterando valor de da variável

Aqui concatenamos nossa string com outra.

```go
{{#include exemplo02/main.go}}
```

### Mode de inicialização + atribuição

Podemos criar uma variável com `:=`, ja atribuindo um valor à ela, onde seu tipo é inferido por Go

```go
{{#include exemplo03/main.go}}
```

### Escopo de variável e funções

A variável `x` é acessada por `f()` pois foi definida no escopo global de nosso programa. E não dentro de `main()`.

```go
{{#include exemplo04/main.go}}
```

### Definindo constantes
```go
{{#include exemplo05/main.go}}
```

### Lendo dados do usuário com `fmt.Scanf()`
```go
{{#include exemplo06/main.go}}
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